package core

import "os"

type settings struct {
	HttpPort string
	LogLevel string
}

var Settings *settings

func initSettings() {
	Settings = &settings{
		HttpPort: os.Getenv("HTTP_PORT"),
		LogLevel: os.Getenv("LOG_LEVEL"),
	}
}
