package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port   string
	DBHost string
	DBPort string
	DBName string
	DBUser string
	DBPass string
	DBSSL  string
}

func Load() Config {
	return Config{
		Port:   env("APP_PORT", "8081"),
		DBHost: env("DB_HOST", "127.0.0.1"),
		DBPort: env("DB_PORT", "5433"),
		DBName: env("DB_NAME", "unsch_horarios"),
		DBUser: env("DB_USER", "postgres"),
		DBPass: os.Getenv("DB_PASSWORD"),
		DBSSL:  env("DB_SSLMODE", "disable"),
	}
}

func (c Config) DatabaseURL() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.DBUser,
		c.DBPass,
		c.DBHost,
		c.DBPort,
		c.DBName,
		c.DBSSL,
	)
}

func env(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}


