package store

import (
	"context"
)

type BearerAuthStoreInterface interface {
	SetJWTToken(ctx context.Context, userID int, jwtToken string) error
	GetJWTToken(ctx context.Context, userID int) (string, error)
	RenewJWTToken(ctx context.Context, userID int) error
	DeleteJWTToken(ctx context.Context, userID int) error
}
