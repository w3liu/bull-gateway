package api

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/w3liu/bull-gateway/infra/code"
	"github.com/w3liu/bull-gateway/infra/response"
	"github.com/w3liu/bull-gateway/pkg/settings/options"
	"github.com/w3liu/bull-gateway/pkg/settings/types"
)

func (h *Handler) Create(c *gin.Context) {
	var r *types.Api
	if err := c.ShouldBindJSON(&r); err != nil {
		response.FailedWithC(c, code.ErrForbidden, err)
		return
	}
	if err := h.srv.ApiSrv().Create(context.Background(), r, options.CreateOptions{}); err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, nil)
}
