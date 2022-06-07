package store

import (
	"context"
	"douyin-app/internal/pkg/model"
)

type BasicAuthStoreInterface interface {
	Get(ctx context.Context, userName string) (*model.UserSecret, error)
	Create(ctx context.Context, userSecret *model.UserSecret) error
	Update(ctx context.Context, userSecret *model.UserSecret) error
	Delete(ctx context.Context, userName string) error
}
