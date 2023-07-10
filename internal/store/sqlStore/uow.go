package sqlstore

import (
	"chat-backend/internal/database"
	"chat-backend/internal/store"
	"github.com/jmoiron/sqlx"
)

type UnitOfWorkSql struct {
	db             *sqlx.DB
	transaction    *sqlx.Tx
	chatRepository *ChatRepositorySql
}

func NewUnitOfWorkSql() *UnitOfWorkSql {
	db := database.DB
	return &UnitOfWorkSql{
		db: db,
	}
}

func (uow *UnitOfWorkSql) Begin() error {
	tx, err := uow.db.Beginx()
	if err != nil {
		return err
	}

	uow.transaction = tx
	uow.chatRepository = &ChatRepositorySql{
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
