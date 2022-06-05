package store

import (
	"context"
	"douyin-app/pkg/model"
)

type BearerAuthStoreInterface interface {
	SetJWTToken(ctx context.Context, jwtClaims *model.JWTClaims, jwtToken string) error
	GetJWTToken(ctx context.Context, jwtClaims *model.JWTClaims) (string, error)
	RenewJWTToken(ctx context.Context, jwtClaims *model.JWTClaims) error
	DeleteJWTToken(ctx context.Context, jwtClaims *model.JWTClaims) error
}
