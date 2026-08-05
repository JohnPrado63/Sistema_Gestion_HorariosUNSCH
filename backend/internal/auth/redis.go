package auth

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client
var redisEnabled = true

func InitRedis(addr, password string, db int) error {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Printf("auth: Redis connection failed: %v - token blacklist disabled", err)
		redisEnabled = false
		return nil
	}

	log.Printf("auth: Redis connected successfully")
	redisEnabled = true
	return nil
}

func IsRedisEnabled() bool {
	return redisEnabled && redisClient != nil
}

func BlacklistToken(ctx context.Context, jti string, expiry time.Duration) error {
	if !redisEnabled || redisClient == nil {
		log.Printf("auth: Redis disabled, skipping blacklist for JTI: %s", jti)
		return nil
	}

	key := fmt.Sprintf("blacklist:%s", jti)
	err := redisClient.Set(ctx, key, "1", expiry).Err()
	if err != nil {
		log.Printf("auth: Failed to blacklist token: %v", err)
		return err
	}

	log.Printf("auth: Token blacklisted: %s", jti)
	return nil
}

func IsTokenBlacklisted(ctx context.Context, jti string) (bool, error) {
	if !redisEnabled || redisClient == nil {
		return false, nil
	}

	key := fmt.Sprintf("blacklist:%s", jti)
	result, err := redisClient.Exists(ctx, key).Result()
	if err != nil {
		log.Printf("auth: Failed to check blacklist: %v", err)
		return false, err
	}

	return result > 0, nil
}

func CloseRedis() error {
	if redisClient != nil {
		return redisClient.Close()
	}
	return nil
}
