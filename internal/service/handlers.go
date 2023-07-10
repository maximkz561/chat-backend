package service

import (
	"chat-backend/internal/domain"
	"chat-backend/internal/store"
	"github.com/google/uuid"
	"log"
)

func addChat(
	chat domain.Chat,
	uow store.UnitOfWork,
) error {
	if err := uow.Begin(); err != nil {
		log.Fatal(err)
	}
	chatRepo := uow.Chat()
	if err := chatRepo.Create(&chat); err != nil {
		_ = uow.Rollback()
		return err
	}
	if err := uow.Commit(); err != nil {
		return err
	}
	return nil
}

func getChat(
	id uuid.UUID,
	uow store.UnitOfWork,
) (*domain.Chat, error) {
	chatRepo := uow.Chat()
	chat, err := chatRepo.Find(id)
	if err != nil {
		return nil, err
	}
	return chat, nil
}

func getUserChats(
	userId uuid.UUID,
	uow store.UnitOfWork,
) ([]*domain.Chat, error) {
	chatRepo := uow.Chat()
	chats, err := chatRepo.FindByUser(userId)
	if err != nil {
		return nil, err
	}
	return chats, nil
}
