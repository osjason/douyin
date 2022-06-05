package redis

import (
	"douyin-app/internal/douyin-api/store"
	"github.com/go-redis/redis/v8"
)

type redisStore struct {
	redisClient *redis.Client
}

func (redisClient *redisStore) BearerAuthorization() store.BearerAuthStoreInterface {
	return newBearerAuthStore(redisClient)
}
