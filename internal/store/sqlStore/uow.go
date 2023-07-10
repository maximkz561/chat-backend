package sqlstore

import (
	"chat-backend/internal/store"
	"database/sql"
)

type UnitOfWorkSql struct {
	db             *sql.DB
	transaction    *sql.Tx
	chatRepository *ChatRepository
}

func (uow *UnitOfWorkSql) Begin() error {
	tx, err := uow.db.Begin()
	if err != nil {
		return err
	}

	uow.transaction = tx
	uow.chatRepository = &ChatRepository{
		store: uow,
	}

	return nil
}

func (uow *UnitOfWorkSql) Commit() error {
	return uow.transaction.Commit()
}

func (uow *UnitOfWorkSql) Rollback() error {
	return uow.transaction.Rollback()
}

func (uow *UnitOfWorkSql) Chat() store.ChatRepository {
	return uow.chatRepository
}
