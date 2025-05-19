package config

import (
	"os"
)

const (
	// DefaultPort - порт по умолчанию
	DefaultPort = "7540"
	// WebDir - директория с фронтендом
	WebDir = "./web"
)

// GetPort возвращает порт из переменной окружения или значение по умолчанию
func GetPort() string {
	if port := os.Getenv("TODO_PORT"); port != "" {
		return port
	}
	return DefaultPort
}
