package v1

import (
	"context"
	"crypto/sha512"
	"douyin-app/internal/douyin-auth/store"
	"douyin-app/internal/pkg/model"
	"fmt"
)

type BasicAuthSrvInterface interface {
	Get(ctx context.Context, userName string) (*model.UserSecret, error)
	Create(ctx context.Context, userName string, clearPasswd string) error
	Update(ctx context.Context, userName string, clearPasswd string) error
	Delete(ctx context.Context, userName string) error
}

func EncryptClearPasswd(clearPasswd string) string {
	return fmt.Sprintf("%x", sha512.Sum512_256([]byte(clearPasswd)))
}

type basicAuthService struct {
	storeFactory store.Factory
}

func (b basicAuthService) Get(ctx context.Context, userName string) (*model.UserSecret, error) {
	var userSec *model.UserSecret
	userSec, err := b.storeFactory.BasicAuthorization().Get(ctx, userName)
	if err != nil {
		return nil, err
	}
	return userSec, nil
}

func (b basicAuthService) Create(ctx context.Context, userName string, clearPasswd string) error {

	userSecret, err := b.storeFactory.BasicAuthorization().Get(ctx, userName)
	if userSecret != nil && err == nil {
		// 数据库中已经存在该条目
		return nil
	}
	// 使用 sha-512 加密密码
	userSecret = &model.UserSecret{UserName: userName}
	userSecret.Password = EncryptClearPasswd(clearPasswd)
	// 创建
	err = b.storeFactory.BasicAuthorization().Create(ctx, userSecret)
	if err != nil {
		return err
	}
	return nil
}

func (b basicAuthService) Update(ctx context.Context, userName string, clearPasswd string) error {
	if _, err := b.storeFactory.BasicAuthorization().Get(ctx, userName); err != nil {
		return err
	}
	var userSecret *model.UserSecret
	// 使用 sha-512 加密密码
	userSecret = &model.UserSecret{UserName: userName}
	userSecret.Password = EncryptClearPasswd(clearPasswd)
	// 更新
	err := b.storeFactory.BasicAuthorization().Update(ctx, userSecret)
	if err != nil {
		return err
	}
	return nil
}

func (b basicAuthService) Delete(ctx context.Context, userName string) error {
	if _, err := b.storeFactory.BasicAuthorization().Get(ctx, userName); err != nil {
		// 压根不存在
		return nil
	}
	err := b.storeFactory.BasicAuthorization().Delete(ctx, userName)
	if err != nil {
		return err
	}
	return nil
}

var _ BasicAuthSrvInterface = (*basicAuthService)(nil)

func newBasicAuth(srv *service) *basicAuthService {
	return &basicAuthService{storeFactory: srv.store}
}
