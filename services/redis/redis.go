package redis

import (
	"log"

	"github.com/go-redis/redis/v9"
)

func NewClient(redisUrl string) *redis.Client {
	opt, err := redis.ParseURL(redisUrl)
	if err != nil {
		log.Fatalf("failed to parse redis url: %v", err)
	}

	return redis.NewClient(opt)
}
