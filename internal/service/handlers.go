package service

import (
	"chat-backend/internal/domain"
	sqlstore "chat-backend/internal/store/sqlStore"
	"github.com/google/uuid"
	"log"
)

func AddChat(
	title string,
	userId uuid.UUID,
) (*domain.Chat, error) {
	chat := domain.NewChat(userId, title)
	uow := sqlstore.NewUnitOfWorkSql()
	if err := uow.Begin(); err != nil {
		log.Fatal(err)
	}
	chatRepo := uow.Chat()
	if err := chatRepo.Create(chat); err != nil {
		_ = uow.Rollback()
		return nil, err
	}
	if err := uow.Commit(); err != nil {
		return nil, err
	}
	return chat, nil
}

func getChat(
	id uuid.UUID,
) (*domain.Chat, error) {
	uow := sqlstore.NewUnitOfWorkSql()
	if err := uow.Begin(); err != nil {
		log.Fatal(err)
	}
	chatRepo := uow.Chat()
	chat, err := chatRepo.Find(id)
	if err != nil {
		return nil, err
	}
	_ = uow.Commit()
	return chat, nil
}

func getUserChats(
	userId uuid.UUID,
) ([]*domain.Chat, error) {
	uow := sqlstore.NewUnitOfWorkSql()
	if err := uow.Begin(); err != nil {
		log.Fatal(err)
	}
	chatRepo := uow.Chat()
	chats, err := chatRepo.FindByUser(userId)
	if err != nil {
		return nil, err
	}
	_ = uow.Commit()
	return chats, nil
}
