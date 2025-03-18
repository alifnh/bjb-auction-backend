package httphandler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/alifnh/bjb-auction-backend/internal/dto"
	"github.com/alifnh/bjb-auction-backend/internal/pkg/ctxutils"
	"github.com/alifnh/bjb-auction-backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

type UserAssetHandler struct {
	userAssetUsecase usecase.UserAssetUsecase
}

func NewUserAssetHandler(userAssetUsecase usecase.UserAssetUsecase) *UserAssetHandler {
	return &UserAssetHandler{userAssetUsecase: userAssetUsecase}
}

func (h *UserAssetHandler) AddFavorite(c *gin.Context) {
	userID, ok := ctxutils.GetUserId(c)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	// Ambil assetID dari parameter URL
	assetIDParam := c.Param("id")
	assetID, err := strconv.ParseInt(assetIDParam, 10, 64)
	log.Println(assetID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid asset ID"})
		return
	}

	var req dto.FavoriteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Tambahkan asset ke daftar favorit
	err = h.userAssetUsecase.AddFavorite(c.Request.Context(), userID, assetID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add asset to favorite"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success add asset to favorite"})
}

func (h *UserAssetHandler) RemoveFavorite(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userIDInt, ok := userID.(int64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	// Ambil assetID dari parameter URL
	assetIDParam := c.Param("id")
	assetID, err := strconv.ParseInt(assetIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid asset ID"})
		return
	}

	// Validasi request body jika dibutuhkan
	var req dto.FavoriteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hapus asset dari daftar favorit
	err = h.userAssetUsecase.RemoveFavorite(c.Request.Context(), userIDInt, assetID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove asset from favorite"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success remove asset from favorite"})
}
