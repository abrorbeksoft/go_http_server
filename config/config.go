package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	HttpPort string
}

func Load() Config {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	cfg := Config{}
	cfg.HttpPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8000"))
	return cfg
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	if val, exists := os.LookupEnv(key); exists {
		return val
	} else {
		return defaultValue
	}
}
