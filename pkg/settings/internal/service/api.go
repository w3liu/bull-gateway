package service

import (
	"context"
	"github.com/w3liu/bull-gateway/pkg/settings/internal/store"
	"github.com/w3liu/bull-gateway/pkg/settings/options"
	"github.com/w3liu/bull-gateway/pkg/settings/types"
)

type ApiSrv interface {
	Create(ctx context.Context, api *types.Api, opts options.CreateOptions) error
	Update(ctx context.Context, api *types.Api, opts options.UpdateOptions) error
	Delete(ctx context.Context, id int64, opts options.DeleteOptions) error
	Page(ctx context.Context, apis []*types.Api, opts options.PageOptions) error
}

type apiService struct {
	store store.Factory
}

func newApiService(s store.Factory) *apiService {
	return &apiService{store: s}
}

func (s *apiService) Create(ctx context.Context, api *types.Api, opts options.CreateOptions) error {
	return s.store.ApiStore().Create(ctx, api, opts)
}

func (s *apiService) Update(ctx context.Context, api *types.Api, opts options.UpdateOptions) error {
	return s.store.ApiStore().Update(ctx, api, opts)
}

func (s *apiService) Delete(ctx context.Context, id int64, opts options.DeleteOptions) error {
	return s.store.ApiStore().Delete(ctx, id, opts)
}

func (s *apiService) Page(ctx context.Context, apis []*types.Api, opts options.PageOptions) error {
	return s.store.ApiStore().Page(ctx, apis, opts)
}
