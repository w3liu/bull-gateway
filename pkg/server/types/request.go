package types

import (
	"context"
	"net/http"
)

type Request struct {
	ctx    context.Context
	origin *http.Request
	params map[string]interface{}
}
