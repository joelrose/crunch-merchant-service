package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

type Config struct {
	FirebaseConfig string
	DatabaseUrl    string
	StripeKey      string
}

func mustGetEnv(env string) string {
	val, exists := os.LookupEnv(env)

	if !exists {
		log.Fatalf("%s environment variable is not set", env)
	}

	return val
}

func LoadConfig() Config {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Debugf("Could not load .env file")
	}

	config := Config{
		FirebaseConfig: mustGetEnv("FIREBASE_CONFIG"),
		DatabaseUrl:    mustGetEnv("DATABASE_URL"),
		StripeKey:      mustGetEnv("STRIPE_KEY"),
	}

	return config
}
