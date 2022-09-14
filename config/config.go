package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

type Deliverect struct {
	BaseUrl      string
	ClientId     string
	ClientSecret string
}

type Config struct {
	FirebaseConfig string
	DatabaseUrl    string
	StripeKey      string
	RedisUrl       string
	Deliverect     Deliverect
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
		RedisUrl:       mustGetEnv("REDISCLOUD_URL"),
		Deliverect: Deliverect{
			BaseUrl:      mustGetEnv("DELIVERECT_BASE_URL"),
			ClientId:     mustGetEnv("DELIVERECT_CLIENT_ID"),
			ClientSecret: mustGetEnv("DELIVERECT_CLIENT_SECRET"),
		},
	}

	return config
}
