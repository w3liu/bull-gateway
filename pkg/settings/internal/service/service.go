package service

import "github.com/w3liu/bull-gateway/pkg/settings/internal/store"

type Service interface {
	ApiSrv() ApiSrv
}

type service struct {
	store store.Factory
}

func NewService(s store.Factory) Service {
	return &service{store: s}
}

func (s *service) ApiSrv() ApiSrv {
	return newApiService(s.store)
}
