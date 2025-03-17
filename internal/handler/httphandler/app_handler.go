package httphandler

import (
	"net/http"

	"github.com/alifnh/bjb-auction-backend/internal/dto"
	"github.com/gin-gonic/gin"
)

type AppHandler struct {
}

func NewAppHandler() *AppHandler {
	return &AppHandler{}
}

func (h *AppHandler) RouteNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, dto.ErrorResponse{
		Message: "Route not found",
	})
}

func (h *AppHandler) Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome To SeeDoco"})
}
