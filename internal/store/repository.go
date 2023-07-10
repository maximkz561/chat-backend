package store

import (
	"chat-backend/internal/domain"
	"github.com/google/uuid"
)

type ChatRepository interface {
	Create(*domain.Chat) error
	Find(uuid uuid.UUID) (*domain.Chat, error)
	FindByUser(uuid uuid.UUID) ([]*domain.Chat, error)
}

type UnitOfWork interface {
	Begin() error
	Commit() error
	Rollback() error
	Chat() ChatRepository
}
