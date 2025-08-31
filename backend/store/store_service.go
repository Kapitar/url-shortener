package store

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

type StorageService struct {
	redisClient *redis.Client
}

var (
	storeService = &StorageService{}
	ctx          = context.Background()
)

func newRedis() *redis.Client {
	url := os.Getenv("REDIS_TLS_URL")
	if url == "" {
		url = os.Getenv("REDIS_URL")
	}
	opt, err := redis.ParseURL(url)
	if err != nil {
		log.Fatalf("invalid redis url: %v", err)
	}

	if opt.TLSConfig == nil {
		opt.TLSConfig = &tls.Config{}
	}
	opt.TLSConfig.ServerName = os.Getenv("REDIS_TLS_SERVER_NAME")
	opt.TLSConfig.MinVersion = tls.VersionTLS12

	return redis.NewClient(opt)
}

func InitializeStore() *StorageService {
	redisClient := newRedis()

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}

	fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)
	storeService.redisClient = redisClient
	return storeService
}

func SaveUrlMapping(shortUrl string, originalUrl string) {
	err := storeService.redisClient.Set(ctx, shortUrl, originalUrl, 0).Err()

	if err != nil {
		panic(fmt.Sprintf("Failed to save URL mapping for shortUrl %s and originalUrl %s: %v", shortUrl, originalUrl, err))
	}
}

func GetOriginalUrl(shortUrl string) string {
	result, err := storeService.redisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to get original URL for shortUrl %s: %v", shortUrl, err))
	}
	return result
}
