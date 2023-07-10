package database

import (
	"chat-backend/core"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

var DB *sqlx.DB

func Connect() {
	var err error
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		core.Settings.DbHost,
		core.Settings.DbPort,
		core.Settings.DbUser,
		core.Settings.DbPass,
		core.Settings.DbName)
	DB, err = sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}
}

func Ping() error {
	return DB.Ping()
}
