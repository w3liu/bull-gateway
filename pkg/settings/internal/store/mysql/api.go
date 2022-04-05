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
	_, err := s.db.Context(ctx).Insert(api)
	return err
}
func (s *apiStore) Update(ctx context.Context, api *types.Api, opts options.UpdateOptions) error {
	_, err := s.db.Context(ctx).Update(api)
	return err
}
func (s *apiStore) Delete(ctx context.Context, id int64, opts options.DeleteOptions) error {
	_, err := s.db.Context(ctx).Delete(&types.Api{Id: id})
	return err
}
func (s *apiStore) Page(ctx context.Context, apis []*types.Api, opts options.PageOptions) error {
	err := s.db.Context(ctx).Limit(opts.Limit(), opts.Start()).Find(&apis)
	return err
}
