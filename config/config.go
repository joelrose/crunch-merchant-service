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

type Stripe struct {
	SecretKey     string
	WebhookSecret string
}

type Config struct {
	FirebaseConfig string
	DatabaseUrl    string
	RedisUrl       string
	Stripe         Stripe
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
		RedisUrl:       mustGetEnv("REDISCLOUD_URL"),
		Stripe: Stripe{
			SecretKey:     mustGetEnv("STRIPE_SECRET_KEY"),
			WebhookSecret: mustGetEnv("STRIPE_WEBHOOK_SIGNATURE"),
		},
		Deliverect: Deliverect{
			BaseUrl:      mustGetEnv("DELIVERECT_BASE_URL"),
			ClientId:     mustGetEnv("DELIVERECT_CLIENT_ID"),
			ClientSecret: mustGetEnv("DELIVERECT_CLIENT_SECRET"),
		},
	}

	return config
}
