package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/w3liu/bull-gateway/infra/utils"
	"github.com/w3liu/bull-gateway/pkg/settings/internal/handler/api"
	"github.com/w3liu/bull-gateway/pkg/settings/internal/store"
)

const basePath = "/setting"

func Register(v1 *gin.RouterGroup, s store.Factory) {
	apiHandler := api.New(s)
	v1.Handle(http.MethodPost, utils.JoinPath(basePath, api.Api), apiHandler.Create)
}
