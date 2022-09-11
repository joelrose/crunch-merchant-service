package config

import (
	"log"
	"os"
)

type Config struct {
	InternalAuthToken string
	DatabaseUrl       string
}

func mustGetEnv(env string) string {
	val, exists := os.LookupEnv(env)

	if !exists {
		log.Fatalf("%s environment variable is not set", env)
	}

	return val
}

func LoadConfig() Config {
	config := Config{
		InternalAuthToken: mustGetEnv("INTERNAL_AUTH_TOKEN"),
		DatabaseUrl:       mustGetEnv("DATABASE_URL"),
	}

	return config
}
