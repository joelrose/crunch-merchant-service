package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

type Deliverect struct {
	BaseUrl      string
	ChannelName  string
	ClientId     string
	ClientSecret string
}

type Stripe struct {
	SecretKey     string
	WebhookSecret string
}

type Auth0 struct {
	Authority string
	Audience  string
}

type Config struct {
	FirebaseConfig string
	DatabaseUrl    string
	RedisUrl       string
	Stripe         Stripe
	Deliverect     Deliverect
	Auth0          Auth0
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
			ChannelName:  "crunch",
			ClientId:     mustGetEnv("DELIVERECT_CLIENT_ID"),
			ClientSecret: mustGetEnv("DELIVERECT_CLIENT_SECRET"),
		},
		Auth0: Auth0{
			Authority: mustGetEnv("AUTH0_AUTHORITY"),
			Audience:  mustGetEnv("AUTH0_AUDIENCE"),
		},
	}

	return config
}
