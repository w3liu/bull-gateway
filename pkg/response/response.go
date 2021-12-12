package response

import (
	"github.com/gin-gonic/gin"
	"github.com/w3liu/bull-gateway/pkg/errors"
	"github.com/w3liu/bull-gateway/tools/log"
	"go.uber.org/zap"
	"net/http"
)

type ErrResponse struct {
	// Code defines the business error code.
	Code int `json:"code"`

	// Message contains the detail of this message.
	// This message is suitable to be exposed to external
	Message string `json:"message"`

	// Reference returns the reference document which maybe useful to solve this error.
	Reference string `json:"reference,omitempty"`
}

func Write(c *gin.Context, err error, data interface{}) {
	if err != nil {
		log.Error("error", zap.Error(err))
		coder := errors.ParseCoder(err)
		c.JSON(coder.HTTPStatus(), ErrResponse{
			Code:      coder.Code(),
			Message:   coder.String(),
			Reference: coder.Reference(),
		})

		return
	}

	c.JSON(http.StatusOK, data)
}
