package config

import (
	"os"
)

type Config struct {
	Port      string
	JWTSecret string
	JWTExpire string // durasi, misal "1h"
}

func LoadConfig() *Config {
	return &Config{
		Port:      os.Getenv("PORT"),
		JWTSecret: os.Getenv("JWT_SECRET"),
		JWTExpire: os.Getenv("JWT_EXPIRE"),
	}
}
