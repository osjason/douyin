package v1

import (
	"context"
	"crypto/sha512"
	"douyin-app/internal/douyin-auth/store"
	"douyin-app/pkg/model"
	"fmt"
)

type BasicAuthSrvInterface interface {
	Get(ctx context.Context, userId int) (*model.UserSecret, error)
	Create(ctx context.Context, userID int, clearPasswd string) error
	Update(ctx context.Context, userID int, clearPasswd string) error
	Delete(ctx context.Context, userID int) error
}

type basicAuthService struct {
	storeFactory store.Factory
}

func (b basicAuthService) Get(ctx context.Context, userId int) (*model.UserSecret, error) {
	var userSec *model.UserSecret
	userSec, err := b.storeFactory.BasicAuthorization().Get(ctx, userId)
	if err != nil {
		return nil, err
	}
	return userSec, nil
}

func (b basicAuthService) Create(ctx context.Context, userID int, clearPasswd string) error {

	userSecret, err := b.storeFactory.BasicAuthorization().Get(ctx, userID)
	if userSecret != nil && err == nil {
		// 数据库中已经存在该条目
		return nil
	}
	// 使用 sha-512 加密密码
	userSecret = &model.UserSecret{UserId: userID}
	var passwd [32]byte = sha512.Sum512_256([]byte(clearPasswd))
	userSecret.Password = fmt.Sprintf("%x", passwd)
	// 创建
	err = b.storeFactory.BasicAuthorization().Create(ctx, userSecret)
	if err != nil {
		return err
	}
	return nil
}

func (b basicAuthService) Update(ctx context.Context, userID int, clearPasswd string) error {
	if _, err := b.storeFactory.BasicAuthorization().Get(ctx, userID); err != nil {
		return err
	}
	var userSecret *model.UserSecret
	// 使用 sha-512 加密密码
	userSecret = &model.UserSecret{UserId: userID}
	var passwd [32]byte = sha512.Sum512_256([]byte(clearPasswd))
	userSecret.Password = fmt.Sprintf("%x", passwd)
	// 更新
	err := b.storeFactory.BasicAuthorization().Update(ctx, userSecret)
	if err != nil {
		return err
	}
	return nil
}

func (b basicAuthService) Delete(ctx context.Context, userID int) error {
	if _, err := b.storeFactory.BasicAuthorization().Get(ctx, userID); err != nil {
		// 压根不存在
		return nil
	}
	err := b.storeFactory.BasicAuthorization().Delete(ctx, userID)
	if err != nil {
		return err
	}
	return nil
}

var _ BasicAuthSrvInterface = (*basicAuthService)(nil)

func newBasicAuth(srv *service) *basicAuthService {
	return &basicAuthService{storeFactory: srv.store}
}
