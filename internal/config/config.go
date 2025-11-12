package config

import (
	"os"
)

type Config struct {
	Env         string
	Port        string
	DatabaseURL string
}

func Load() Config {
	return Config{
		Env:         getEnv("APP_ENV", "dev"),
		Port:        getEnv("PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", "postgresql://root@localhost:26257/app?sslmode=disable"),
	}
}

func getEnv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
