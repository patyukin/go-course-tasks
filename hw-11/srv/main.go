package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gopkg.in/yaml.v3"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// RateLimiter определяет структуру для хранения информации о лимитах запросов
type RateLimiter struct {
	rate      int        // Количество токенов, добавляемых в ведро каждую секунду
	bucket    int        // Максимальное количество токенов в ведре
	tokens    int        // Текущее количество токенов
	lastCheck time.Time  // Время последней проверки ведра
	mu        sync.Mutex // Мьютекс для синхронизации доступа
}

// NewRateLimiter создает новый RateLimiter с заданными параметрами
func NewRateLimiter(rate, bucket int) *RateLimiter {
	return &RateLimiter{
		rate:      rate,
		bucket:    bucket,
		tokens:    bucket,
		lastCheck: time.Now(),
	}
}

// Allow проверяет, можно ли выполнить запрос
func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(rl.lastCheck).Seconds()

	// Добавляем токены за прошедшее время
	rl.tokens += int(elapsed * float64(rl.rate))
	if rl.tokens > rl.bucket {
		rl.tokens = rl.bucket
	}

	rl.lastCheck = now

	if rl.tokens > 0 {
		rl.tokens--
		return true
	}

	return false
}

// Config структура для хранения параметров конфигурации
type Config struct {
	Rate         int      `yaml:"rate"`
	Bucket       int      `yaml:"bucket"`
	GlobalRate   int      `yaml:"global_rate"`
	GlobalBucket int      `yaml:"global_bucket"`
	Address      string   `yaml:"address"`
	Whitelist    []string `yaml:"whitelist"`
	Blacklist    []string `yaml:"blacklist"`
}

// LoadConfig загружает конфигурацию из файла
func LoadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// RateLimiters структура для хранения лимитеров запросов для каждого IP и общего лимитера
type RateLimiters struct {
	client *redis.Client
	config *Config
	mu     sync.Mutex
	global *RateLimiter
}

// NewRateLimiters создает новую структуру RateLimiters
func NewRateLimiters(config *Config, client *redis.Client) *RateLimiters {
	ctx := context.Background()
	for _, ip := range config.Whitelist {
		client.SAdd(ctx, "whitelist", ip)
	}
	for _, ip := range config.Blacklist {
		client.SAdd(ctx, "blacklist", ip)
	}
	return &RateLimiters{
		client: client,
		config: config,
		global: NewRateLimiter(config.GlobalRate, config.GlobalBucket),
	}
}

func (rls *RateLimiters) GetLimiterForIP(ctx context.Context, ip string) (*RateLimiter, error) {
	rls.mu.Lock()
	defer rls.mu.Unlock()

	val, err := rls.client.Get(ctx, ip).Result()
	if errors.Is(err, redis.Nil) {
		// Если нет данных для данного IP, создаем новый лимитер
		return NewRateLimiter(rls.config.Rate, rls.config.Bucket), nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to get rate limiter for IP %s: %w", ip, err)
	}

	// Преобразуем данные из Redis в токены и lastCheck
	data := strings.Split(val, ":")
	tokens, err := strconv.Atoi(data[0])
	if err != nil {
		return nil, fmt.Errorf("failed to parse tokens for IP %s: %w", ip, err)
	}

	lastCheck, err := strconv.ParseInt(data[1], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse lastCheck for IP %s: %w", ip, err)
	}

	return &RateLimiter{
		rate:      rls.config.Rate,
		bucket:    rls.config.Bucket,
		tokens:    tokens,
		lastCheck: time.Unix(0, lastCheck),
	}, nil
}

func (rls *RateLimiters) SaveLimiterForIP(ctx context.Context, ip string, limiter *RateLimiter) error {
	rls.mu.Lock()
	defer rls.mu.Unlock()

	val := fmt.Sprintf("%d:%d", limiter.tokens, limiter.lastCheck.UnixNano())
	err := rls.client.Set(ctx, ip, val, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func checkRateLimit(ctx context.Context, conn net.Conn, rls *RateLimiters) error {
	ip, _, err := net.SplitHostPort(conn.RemoteAddr().String())
	fmt.Println("ip =", ip)
	if err != nil {
		return fmt.Errorf("error getting IP address: %w", err)
	}

	// Проверка черного списка
	blacklisted, err := rls.client.SIsMember(ctx, "blacklist", ip).Result()
	if err != nil {
		return fmt.Errorf("error checking blacklist: %w", err)
	}

	if blacklisted {
		if _, err = conn.Write([]byte("HTTP/1.1 403 Forbidden\r\n\r\nForbidden")); err != nil {
			return fmt.Errorf("error writing response: %w", err)
		}

		return fmt.Errorf("IP is blacklisted")
	}

	// Проверка белого списка
	whitelisted, err := rls.client.SIsMember(ctx, "whitelist", ip).Result()
	if err != nil {
		return fmt.Errorf("error checking whitelist: %w", err)
	}

	globalLimiter := rls.global

	if !globalLimiter.Allow() {
		if _, err = conn.Write([]byte("HTTP/1.1 429 Too Many Requests\r\n\r\nToo Many Requests (Global)")); err != nil {
			return fmt.Errorf("error writing response: %w", err)
		}

		return fmt.Errorf("global rate limit exceeded")
	}

	if !whitelisted {
		ipLimiter, iplErr := rls.GetLimiterForIP(ctx, ip)
		if iplErr != nil {
			return fmt.Errorf("error getting rate limiter: %w", err)
		}

		if !ipLimiter.Allow() {
			if _, err = conn.Write([]byte("HTTP/1.1 429 Too Many Requests\r\n\r\nToo Many Requests (Per IP)")); err != nil {
				return fmt.Errorf("error writing response: %w", err)
			}

			return fmt.Errorf("per IP rate limit exceeded")
		}

		// Сохраняем новое состояние лимитера в Redis
		if err = rls.SaveLimiterForIP(ctx, ip, ipLimiter); err != nil {
			return fmt.Errorf("error saving rate limiter: %w", err)
		}
	}

	return nil
}

func handleConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		if err := conn.Close(); err != nil {
			fmt.Println("Error closing connection from handleConnection:", err)
		}
	}(conn)

	if _, err := conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello, World!")); err != nil {
		return
	}
}

func main() {
	config, err := LoadConfig(os.Getenv("CONFIG_PATH"))
	if err != nil {
		log.Fatalln("Error loading config:", err)
	}

	redisAddr := os.Getenv("REDIS_ADDR")

	client := redis.NewClient(&redis.Options{Addr: redisAddr})

	ctx := context.Background()

	rls := NewRateLimiters(config, client)
	srv, err := net.Listen("tcp", config.Address)
	if err != nil {
		log.Fatalln("Error starting server:", err)
	}

	defer func(srv net.Listener) {
		if err = srv.Close(); err != nil {
			fmt.Println("Error closing server:", err)
		}
	}(srv)

	fmt.Println("Server started on", config.Address)

	for {
		fmt.Println("srv addr := ", srv.Addr())
		conn, connErr := srv.Accept()
		if connErr != nil {
			fmt.Println("Error accepting connection:", connErr)
			continue
		}

		if crlErr := checkRateLimit(ctx, conn, rls); crlErr != nil {
			fmt.Println("checkRateLimit error:", crlErr)
			fmt.Println("Closing connection...")
			errConn := conn.Close()
			if errConn != nil {
				log.Fatalln("Error closing connection:", errConn)
			}
		} else {
			go handleConnection(conn)
		}
	}
}
