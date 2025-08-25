package database

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var RedisDB *redis.Client
var Ctx = context.Background()

func InitRedisBD() {
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
}
