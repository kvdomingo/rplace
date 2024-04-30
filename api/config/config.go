package config

import (
	"fmt"
	"os"
)

func GetEnv(key string) string {
	env := os.Getenv(key)
	if env == "" {
		panic(fmt.Sprintf("Environment variable %s not set", key))
	}
	return env
}

func RedisConnectionString() string {
	return GetEnv("REDIS_URL")
}
