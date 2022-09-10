package config

import (
	"log"
	"os"
)

type Config struct {
	InternalAuthToken string
}

func mustGetEnv(env string) string {
	env, exists := os.LookupEnv(env)

	if !exists {
		log.Fatalf("%v environment variable is not set", env)
	}

	return env
}

func LoadConfig() (Config) {
	config := Config{
		InternalAuthToken: mustGetEnv("INTERNAL_AUTH_TOKEN"),
	}

	return config
}
