package v1

import (
	"context"
	"douyin-app/internal/douyin-auth/store"
	"douyin-app/pkg/model"
	"fmt"
	"github.com/golang-jwt/jwt"
)

// BearerAuthSrvInterface 对应 StoreInterface 那边的所有动作
type BearerAuthSrvInterface interface {
	SetJWTToken(ctx context.Context, user *model.User) error
	CompareJWTToken(ctx context.Context, user *model.User, token string) (bool, error)
	RenewJWTToken(ctx context.Context, user *model.User) error
	DeleteJWTToken(ctx context.Context, user *model.User) error
}

type bearerAuthService struct {
	storeFactory store.Factory
}

func (b bearerAuthService) SetJWTToken(ctx context.Context, user *model.User) error {
	token, err := b.storeFactory.BearerAuthorization().GetJWTToken(ctx, user.Id)
	// 如果已经存在 token，直接返回
	if err == nil {
		return fmt.Errorf("user_id 已经存在 Redis 中，无法创建")
	}
	myClaim, err := model.NewJWTClaims(user.Id)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodES256, myClaim)
	token, err = jwtToken.SigningString()
	if err != nil {
		return fmt.Errorf("jwt token 无法转换成 string 类型")
	}
	err = b.storeFactory.BearerAuthorization().SetJWTToken(ctx, user.Id, token)
	if err != nil {
		return err
	}
	return nil
}

func (b bearerAuthService) CompareJWTToken(ctx context.Context, user *model.User, token string) (bool, error) {
	tokenInRedis, err := b.storeFactory.BearerAuthorization().GetJWTToken(ctx, user.Id)
	if err != nil {
		return false, err
	}
	if tokenInRedis != token {
		return false, nil
	}
	return true, nil
}

func (b bearerAuthService) RenewJWTToken(ctx context.Context, user *model.User) error {
	if err := b.storeFactory.BearerAuthorization().RenewJWTToken(ctx, user.Id); err != nil {
		return err
	}
	return nil
}

func (b bearerAuthService) DeleteJWTToken(ctx context.Context, user *model.User) error {
	if _, err := b.storeFactory.BearerAuthorization().GetJWTToken(ctx, user.Id); err != nil {
		return nil
	}
	err := b.storeFactory.BearerAuthorization().DeleteJWTToken(ctx, user.Id)
	if err != nil {
		return err
	}
	return nil
}

// todo auth Service 层待实现
var _ BearerAuthSrvInterface = (*bearerAuthService)(nil)

func newBearerAuth(srv *service) *bearerAuthService {
	return &bearerAuthService{storeFactory: srv.store}
}
