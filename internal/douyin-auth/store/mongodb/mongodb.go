package mongodb

import (
	"douyin-app/internal/douyin-auth/store"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoStore struct {
	mgoClient *mongo.Client
}

func (m mongoStore) BasicAuthorization() store.BasicAuthStoreInterface {
	return newBasicAuthStore(m)
}
