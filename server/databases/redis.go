package databases

import (
	"context"
	"flag"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

var (
	redisHost     = flag.String("redis_host", "localhost", "Redis host")
	redisPort     = flag.Int("redis_port", 6379, "Redis port")
	redisPassword = flag.String("redis_pass", "", "Redis password")
	redisUsername = flag.String("redis_username", "", "Redis username")
	redisDB       = flag.Int("redis_db", 0, "Redis db")
)

type RedisClientProvider struct {
	RedisClient *redis.Client
}

func NewRedisClientProvider() *RedisClientProvider {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     *redisHost + ":" + strconv.Itoa(*redisPort),
		Username: *redisUsername,
		Password: *redisPassword,
		DB:       *redisDB,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	pong := redisClient.Ping(ctx)
	if pong == nil {
		panic("redis ping failed")
	}
	return &RedisClientProvider{
		RedisClient: redisClient,
	}
}
