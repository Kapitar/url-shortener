package store

import (
	"context"
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

func InitializeStore() *StorageService {
	redisURL := os.Getenv("REDIS_URL")

	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		log.Fatalf("invalid redis url: %v", err)
	}

	redisClient := redis.NewClient(opt)

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
