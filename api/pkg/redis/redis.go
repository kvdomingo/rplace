package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/kvdomingo/rplace/config"
	"github.com/redis/go-redis/v9"
)

var Client = NewClient(0)

func NewClient(db int) *redis.Client {
	return redis.NewClient(
		&redis.Options{
			Addr:     config.RedisConnectionString(),
			Password: "",
			DB:       db,
		},
	)
}

func Health() string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	_, err := Client.Ping(ctx).Result()
	if err != nil {
		return fmt.Sprint(err)
	}
	return "ok"
}

func FlushDB() error {
	return Client.FlushDB(context.Background()).Err()
}
