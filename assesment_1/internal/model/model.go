package model

type Message struct {
	Token  string
	FileID string
	Data   string
}

type Cache map[string][]Message
