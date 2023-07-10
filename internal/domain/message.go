package domain

import (
	"github.com/google/uuid"
	"time"
)

type author int

const (
	user = iota + 1
	system
)

type Message struct {
	Id        uuid.UUID
	Content   string
	ChatID    uuid.UUID
	CreatedAt time.Time
	Author    author
}

func newMessage(content string, chatID uuid.UUID, author author) Message {
	return Message{
		Id:        uuid.New(),
		Content:   content,
		ChatID:    chatID,
		CreatedAt: time.Now(),
		Author:    author,
	}
}
