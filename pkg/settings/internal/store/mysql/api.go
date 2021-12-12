package mysql

import (
	"context"
	"github.com/w3liu/bull-gateway/pkg/settings/options"
	"github.com/w3liu/bull-gateway/pkg/settings/types"
	"github.com/w3liu/bull-gateway/tools/mysql"
)

type apiStore struct {
	db *mysql.Store
}

func newApiStore(db *mysql.Store) *apiStore {
	return &apiStore{db: db}
}

func (s *apiStore) Create(ctx context.Context, api *types.Api, opts options.CreateOptions) error {
	return nil
}
func (s *apiStore) Update(ctx context.Context, api *types.Api, opts options.UpdateOptions) error {
	return nil
}
func (s *apiStore) Delete(ctx context.Context, id int64, opts options.DeleteOptions) error {
	return nil
}
func (s *apiStore) Page(ctx context.Context, api *types.Api, opts options.PageOptions) error {
	return nil
}
