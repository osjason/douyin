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

func (authStore BearerAuthStore) SetJWTToken(ctx context.Context, jwtClaims *model.JWTClaims, jwtToken string) error {
	if err := authStore.redisClient.Set(ctx, strconv.Itoa(jwtClaims.UserId), jwtToken, model.JWTTokenDuration).Err(); err != nil {
		return err
	}
	return nil
}

func (authStore BearerAuthStore) GetJWTToken(ctx context.Context, jwtClaims *model.JWTClaims) (string, error) {
	token, err := authStore.redisClient.Get(ctx, strconv.Itoa(jwtClaims.UserId)).Result()
	if err != nil {
		return "", err
	}
	return token, nil
}

func (authStore BearerAuthStore) RenewJWTToken(ctx context.Context, jwtClaims *model.JWTClaims) error {
	if err := authStore.redisClient.Expire(ctx, strconv.Itoa(jwtClaims.UserId), model.JWTTokenDuration).Err(); err != nil {
		return err
	}
	return nil
}

func (authStore BearerAuthStore) DeleteJWTToken(ctx context.Context, jwtClaims *model.JWTClaims) error {
	if err := authStore.redisClient.Del(ctx, strconv.Itoa(jwtClaims.UserId)).Err(); err != nil {
		return err
	}
	return nil
}

func newBearerAuthStore(rdStore *redisStore) *BearerAuthStore {
	return &BearerAuthStore{redisClient: rdStore.redisClient}
}
