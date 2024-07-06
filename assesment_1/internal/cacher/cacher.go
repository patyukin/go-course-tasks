package cacher

import (
	"github.com/patyukin/go-course-tasks/assesment_1/internal/model"
	"sync"
)

type Cacher struct {
	mu sync.Mutex
	c  model.Cache
}

func New() *Cacher {
	return &Cacher{
		mu: sync.Mutex{},
		c:  make(model.Cache),
	}
}

func (c *Cacher) Cache(msg model.Message) {
	c.mu.Lock()
	c.c[msg.FileID] = append(c.c[msg.FileID], msg)
	c.mu.Unlock()
}
