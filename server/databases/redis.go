package databases

//
//import (
//	"context"
//	"flag"
//	"fmt"
//	"github.com/redis/go-redis/v9"
//	"log/slog"
//	"strconv"
//	"time"
//)
//
//var (
//	redisHost     = flag.String("redisHost", "localhost", "Redis host")
//	redisPort     = flag.Int("redisPort", 6379, "Redis port")
//	redisPassword = flag.String("redisPass", "", "Redis password")
//	redisUsername = flag.String("redisUsername", "", "Redis username")
//	redisDB       = flag.Int("redisDB", 0, "Redis db")
//)
//
//type RedisClientProvider struct {
//	RedisClient *redis.Client
//}
//
//func (r RedisClientProvider) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
//	return r.RedisClient.Set(ctx, key, value, expiration).Err()
//}
//
//func (r RedisClientProvider) Get(ctx context.Context, key string) ([]byte, error) {
//	return r.RedisClient.Get(ctx, key).Bytes()
//}
//
//func NewRedisClientProvider() *RedisClientProvider {
//	redisClient := redis.NewClient(&redis.Options{
//		Addr:     *redisHost + ":" + strconv.Itoa(*redisPort),
//		Username: *redisUsername,
//		Password: *redisPassword,
//		DB:       *redisDB,
//	})
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//	if err := redisClient.Ping(ctx); err != nil {
//		slog.Error(fmt.Sprintf("Failed to initialize redis client. err : %s", err))
//		_ = redisClient.Close()
//		redisClient = nil
//		panic(err)
//	}
//	return &RedisClientProvider{
//		RedisClient: redisClient,
//	}
//}
