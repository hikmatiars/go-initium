package server

import (
	"context"
	"fmt"
	"go-initium/internal/config"
	"log"

	"github.com/go-redis/redis/v8"
)

func New(ctx context.Context, cfg config.Config) {
	// connect to redis
	cache := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.Db,
	})
	if _, err := cache.Ping(ctx).Result(); err != nil {
		log.Fatalf("failed connect to redis: %s", err)
	}
}
