package domain

import (
	"github.com/google/uuid"
	"time"
)

type Chat struct {
	Id        uuid.UUID
	Messages  []Message
	UserId    uuid.UUID
	Title     string
	CreatedAt time.Time
}

func newChat(userId uuid.UUID, title string) *Chat {
	return &Chat{
		Id:        uuid.New(),
		Messages:  []Message{},
		UserId:    userId,
		Title:     title,
		CreatedAt: time.Now(),
	}
}
