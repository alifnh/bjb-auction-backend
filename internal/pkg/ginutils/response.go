package ginutils

import (
	"net/http"

	"github.com/alifnh/bjb-auction-backend/internal/constant"
	"github.com/alifnh/bjb-auction-backend/internal/dto"
	"github.com/gin-gonic/gin"
)

// convenience method, errors should be handled implicitly by the middleware
func ResponseOKPlain(c *gin.Context) {
	ResponseOKData(c, nil)
}

func ResponseOKData(c *gin.Context, data interface{}) {
	ResponseOK(c, constant.MessageResponseSuccess, data)
}

func ResponseOK(c *gin.Context, message string, data interface{}) {
	ResponseSuccessJSON(c, http.StatusOK, message, data)
}

func ResponseSuccessJSON(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, dto.SuccessResponse{
		Message: message,
		Data:    data,
	})
}
