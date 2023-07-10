package core

import "os"

type settings struct {
	HttpPort string
	LogLevel string

	DbHost string
	DbPort string
	DbUser string
	DbPass string
	DbName string
}

var Settings *settings

func initSettings() {
	Settings = &settings{
		HttpPort: os.Getenv("HTTP_PORT"),
		LogLevel: os.Getenv("LOG_LEVEL"),

		DbHost: os.Getenv("DB_HOST"),
		DbPort: os.Getenv("DB_PORT"),
		DbUser: os.Getenv("DB_USER"),
		DbPass: os.Getenv("DB_PASS"),
		DbName: os.Getenv("DB_NAME"),
	}
}
