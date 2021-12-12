package store

import (
	"context"
	"github.com/w3liu/bull-gateway/pkg/settings/options"
	"github.com/w3liu/bull-gateway/pkg/settings/types"
)

type ApiStore interface {
	Create(ctx context.Context, api *types.Api, opts options.CreateOptions) error
	Update(ctx context.Context, api *types.Api, opts options.UpdateOptions) error
	Delete(ctx context.Context, id int64, opts options.DeleteOptions) error
	Page(ctx context.Context, api *types.Api, opts options.PageOptions) error
}
