package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/w3liu/bull-gateway/infra/code"
	"github.com/w3liu/bull-gateway/infra/errors"
	"github.com/w3liu/bull-gateway/infra/log"
	"go.uber.org/zap"
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

func Write(c *gin.Context, data interface{}, err error) {
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

func Failed(c *gin.Context, err error) {
	Write(c, nil, err)
}

func FailedWithC(c *gin.Context, code int, err error) {
	Failed(c, errors.WrapC(err, code, ""))
}

func Success(c *gin.Context, data interface{}) {
	Write(c, data, nil)
}
