package mongodb

import (
	"context"
	"douyin-app/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	_ "go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
)

type BasicAuthStore struct {
	mgoClient *mongo.Client
}

func (b BasicAuthStore) Get(ctx context.Context, userId int) (*model.UserSecret, error) {
	coll := b.mgoClient.Database("user").Collection("secret")
	var userSecret *model.UserSecret
	err := coll.FindOne(ctx, bson.D{{"user_id", strconv.Itoa(userId)}}).Decode(userSecret)
	if err != nil {
		return nil, err
	}
	return userSecret, err

}

func (b BasicAuthStore) Create(ctx context.Context, userSecret *model.UserSecret) error {
	coll := b.mgoClient.Database("user").Collection("secret")
	_, err := coll.InsertOne(ctx, userSecret)
	if err != nil {
		return err
	}
	return nil
}

func (b BasicAuthStore) Update(ctx context.Context, newUserSecret *model.UserSecret) error {
	coll := b.mgoClient.Database("user").Collection("secret")
	updateBson := bson.D{{
		"$set", bson.D{{
			"password", newUserSecret.Password,
		}},
	}}
	_, err := coll.UpdateOne(ctx, bson.D{{"user_id", strconv.Itoa(newUserSecret.UserId)}}, updateBson)
	if err != nil {
		return err
	}
	return nil

}

func (b BasicAuthStore) Delete(ctx context.Context, userID int) error {
	coll := b.mgoClient.Database("user").Collection("secret")
	_, err := coll.DeleteOne(ctx, bson.D{{"user_id", strconv.Itoa(userID)}})
	if err != nil {
		return err
	}
	return nil
}

func newBasicAuthStore(store mongoStore) *BasicAuthStore {
	return &BasicAuthStore{mgoClient: store.mgoClient}
}
