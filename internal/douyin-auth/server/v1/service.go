package v1

import "douyin-app/internal/douyin-auth/store"

type ServiceInterface interface {
	BearerAuth() BearerAuthSrvInterface
	BasicAuth() BasicAuthSrvInterface
}

type service struct {
	store store.Factory
}

func (s *service) BearerAuth() BearerAuthSrvInterface {
	return newBearerAuth(s)
}

func (s *service) BasicAuth() BasicAuthSrvInterface {
	return newBasicAuth(s)
}

func NewService(store store.Factory) ServiceInterface {
	return &service{store: store}
}
