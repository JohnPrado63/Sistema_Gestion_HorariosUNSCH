package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	Port   string
	DBHost string
	DBPort string
	DBName string
	DBUser string
	DBPass string
	DBSSL  string
	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDB       int
}

func Load() Config {
	loadEnvFile()
	return Config{
		Port:   env("APP_PORT", "8081"),
		DBHost: env("DB_HOST", "127.0.0.1"),
		DBPort: env("DB_PORT", "5433"),
		DBName: env("DB_NAME", "unsch_horarios"),
		DBUser: env("DB_USER", "postgres"),
		DBPass: os.Getenv("DB_PASSWORD"),
		DBSSL:  env("DB_SSLMODE", "disable"),
		RedisHost:     env("REDIS_HOST", "localhost"),
		RedisPort:     env("REDIS_PORT", "6379"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
		RedisDB:       0,
	}
}

func loadEnvFile() {
	cwd, _ := os.Getwd()

	candidates := []string{
		filepath.Join(cwd, ".env"),
		filepath.Join(cwd, "..", ".env"),
		filepath.Join(cwd, "..", "..", ".env"),
	}
	for _, path := range candidates {
		absPath, _ := filepath.Abs(path)
		log.Printf("config: checking .env at %s", absPath)
		if _, err := os.Stat(path); err == nil {
			log.Printf("config: .env found, attempting load...")
			err := godotenv.Load(path)
			if err != nil {
				log.Printf("config: godotenv.Load error: %v", err)
				continue
			}
			log.Printf("config: loaded .env from %s", path)
			return
		}
	}
	log.Printf("config: no .env file found in any candidate path")
}

func backendDir(file string) string {
	exe, err := os.Executable()
	if err != nil {
		return file
	}
	dir, _ := filepath.EvalSymlinks(filepath.Dir(exe))
	return filepath.Join(dir, "..", file)
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

func (c Config) RedisAddr() string {
	return fmt.Sprintf("%s:%s", c.RedisHost, c.RedisPort)
}

func env(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
