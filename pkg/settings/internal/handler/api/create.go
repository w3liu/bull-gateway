package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/w3liu/bull-gateway/pkg/response"
	"github.com/w3liu/bull-gateway/pkg/settings/options"
	"github.com/w3liu/bull-gateway/pkg/settings/types"
)

func (h *Handler) Create(c *gin.Context) {
	var r *types.Api
	if err := c.ShouldBindJSON(&r); err != nil {
		response.Write(c, err, nil)
		return
	}
	if err := h.srv.ApiSrv().Create(context.Background(), r, options.CreateOptions{}); err != nil {
		response.Write(c, err, nil)
	}
	response.Write(c, nil, nil)
}
