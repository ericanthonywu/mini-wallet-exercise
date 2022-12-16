package Config

import (
	"github.com/go-redis/redis"
	"os"
)

func RedisClientConfig() *redis.Options {
	return &redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	}
}
