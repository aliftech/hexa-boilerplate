package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func Load() (*Config, error) {
	_ = godotenv.Load() // ignore error if .env missing

	return &Config{
		Server: ServerConfig{
			Port: getEnv("PORT", "3000"),
			Env:  getEnv("APP_ENV", "development"),
		},
		DB: DBConfig{
			Host:            getEnv("DB_HOST", "localhost"),
			Port:            getEnvInt("DB_PORT", 3306),
			User:            getEnv("DB_USER", "root"),
			Password:        getEnv("DB_PASSWORD", ""),
			DBName:          getEnv("DB_NAME", "hexa_sso"),
			TimeZone:        getEnv("DB_TIMEZONE", "Local"),
			MaxOpenConns:    getEnvInt("DB_MAX_OPEN_CONNS", 25),
			MaxIdleConns:    getEnvInt("DB_MAX_IDLE_CONNS", 5),
			ConnMaxLifetime: time.Duration(getEnvInt("DB_CONN_MAX_LIFETIME", 300)) * time.Second,
			Env:             getEnv("APP_ENV", "development"),
		},
	}, nil
}

// --- Helpers ---
func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if v := os.Getenv(key); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			return i
		}
	}
	return fallback
}
