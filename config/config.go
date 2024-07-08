package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	DB_HOST     string
	DB_PORT     int
	DB_USER     string
	DB_NAME     string
	DB_PASSWORD string
	URL_PORT    string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}
	config.DB_HOST = cast.ToString(Coalasce("DB_HOST", "localhost"))
	config.DB_NAME = cast.ToString(Coalasce("DB_NAME", "postgres"))
	config.DB_PORT = cast.ToInt(Coalasce("DB_PORT", 5432))
	config.DB_USER = cast.ToString(Coalasce("DB_USER", "postgres"))
	config.DB_PASSWORD = cast.ToString(Coalasce("DB_PASSWORD", "0412"))
	config.URL_PORT = cast.ToString(Coalasce("URL_PORT", "50051"))

	return config
}

func Coalasce(key string, deafultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}
	return deafultValue
}
