package mysql

import (
	"github.com/w3liu/bull-gateway/pkg/settings/internal/store"
	"github.com/w3liu/bull-gateway/tools/mysql"
)

type myFactory struct {
	db *mysql.Store
}

func newFactory(s *mysql.Store) store.Factory {
	return &myFactory{db: s}
}

func (s *myFactory) ApiStore() store.ApiStore {
	return newApiStore(s.db)
}

func (s *myFactory) Close() error {
	return nil
}
