package response

import (
	"github.com/gin-gonic/gin"
	"github.com/w3liu/bull-gateway/pkg/code"
	"github.com/w3liu/bull-gateway/pkg/errors"
	"github.com/w3liu/bull-gateway/tools/log"
	"go.uber.org/zap"
	"net/http"
)

type Response struct {
	// Code defines the business error code.
	Code int `json:"code"`
	// Message contains the detail of this message.
	// This message is suitable to be exposed to external
	Message string `json:"message"`

	// Reference returns the reference document which maybe useful to solve this error.
	Reference string `json:"reference,omitempty"`

	// Data returns the business data for indicated api.
	Data interface{} `json:"data"`
}

func Write(c *gin.Context, err error, data interface{}) {
	if err != nil {
		log.Error("error", zap.Error(err))
		coder := errors.ParseCoder(err)
		c.JSON(coder.HTTPStatus(), Response{
			Code:      coder.Code(),
			Message:   coder.String(),
			Reference: coder.Reference(),
		})

		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    code.Success,
		Message: "success",
		Data:    data,
	})
}
