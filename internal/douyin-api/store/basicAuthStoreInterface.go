package store

import (
	"context"
	"douyin-app/pkg/model"
)

type BasicAuthStoreInterface interface {
	Get(ctx context.Context, userId int) (*model.UserSecret, error)
	Create(ctx context.Context, userSecret *model.UserSecret) error
	Update(ctx context.Context, userSecret *model.UserSecret) error
	Delete(ctx context.Context, userID int) error
}
