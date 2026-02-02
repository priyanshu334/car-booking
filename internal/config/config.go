package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv  string
	AppPort string

	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
	DBSSL  string
}

var Cfg *Config

func Load() {
	_ = godotenv.Load()

	Cfg = &Config{
		AppEnv:  get("APP_ENV", "development"),
		AppPort: get("APP_PORT", "8080"),

		DBHost: get("DB_HOST", "localhost"),
		DBPort: get("DB_PORT", "5432"),
		DBUser: get("DB_USER", "postgres"),
		DBPass: get("DB_PASS", "postgres"),
		DBName: get("DB_NAME", "car_booking"),
		DBSSL:  get("DB_SSL", "disable"),
	}
}

func get(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}
