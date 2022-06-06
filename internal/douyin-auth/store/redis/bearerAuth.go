package redis

import (
	"context"
	"douyin-app/pkg/model"
	"github.com/go-redis/redis/v8"
	"strconv"
)

type BearerAuthStore struct {
	redisClient *redis.Client
}

func (authStore BearerAuthStore) SetJWTToken(ctx context.Context, userID int, jwtToken string) error {
	if err := authStore.redisClient.Set(ctx, strconv.Itoa(userID), jwtToken, model.JWTTokenDuration).Err(); err != nil {
		return err
	}
	return nil
}

func (authStore BearerAuthStore) GetJWTToken(ctx context.Context, userID int) (string, error) {
	token, err := authStore.redisClient.Get(ctx, strconv.Itoa(userID)).Result()
	if err != nil {
		return "", err
	}
	return token, nil
}

// RenewJWTToken 延长 token 在 redis 中的时间
func (authStore BearerAuthStore) RenewJWTToken(ctx context.Context, userID int) error {
	if err := authStore.redisClient.Expire(ctx, strconv.Itoa(userID), model.JWTTokenDuration).Err(); err != nil {
		return err
	}
	return nil
}

func (authStore BearerAuthStore) DeleteJWTToken(ctx context.Context, userID int) error {
	if err := authStore.redisClient.Del(ctx, strconv.Itoa(userID)).Err(); err != nil {
		return err
	}
	return nil
}

func newBearerAuthStore(rdStore *redisStore) *BearerAuthStore {
	return &BearerAuthStore{redisClient: rdStore.redisClient}
}
