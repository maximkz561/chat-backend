package http_types

import (
	"chat-backend/internal/domain"
	"github.com/google/uuid"
	"time"
)

type CreateChat struct {
	UserId uuid.UUID `json:"user_id"`
	Title  string    `json:"title"`
}

type ReadChat struct {
	Id        string `json:"id"`
	CreatedAt string `json:"created_at"`
	CreateChat
}

func NewReadChatFromDomain(chat *domain.Chat) *ReadChat {
	return &ReadChat{
		Id:        chat.Id.String(),
		CreatedAt: chat.CreatedAt.Format(time.RFC3339),
		CreateChat: CreateChat{
			UserId: chat.UserId,
			Title:  chat.Title,
		},
	}
}
