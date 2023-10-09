package main

import (
	"os"
	"fmt"
	"context"
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/go-redis/redis/v8"
)

func setupRedis() (*redis.Client, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file:", err)
	}
	redisAddr := os.Getenv("REDIS_URL")
    redisClient := redis.NewClient(&redis.Options{
        Addr: redisAddr,
    })
    ctx := context.Background()
    pong, err := redisClient.Ping(ctx).Result()
    if err != nil {
        return nil, fmt.Errorf("Failed to connect to Redis: %s", err.Error())
    }

    log.Println("Connected to Redis: %s", pong)
    return redisClient, nil
}

func main(){
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file:", err)
	}

	redisClient, err := setupRedis()
    if err != nil {
        log.Println("Redis setup failed: %s", err.Error())
        time.Sleep(2 * time.Second)
        return
    }

	member_key := "online_users"

	ctx := context.Background()

	start := 0
    end := -1 // To get all elements

    // Use LRange to retrieve the list data
    result, err := client.LRange(ctx, member_key, int64(start), int64(end)).Result()
    if err != nil {
        return nil, err
    }
	log.Println("==> ", result)

    defer redisClient.Close()
}