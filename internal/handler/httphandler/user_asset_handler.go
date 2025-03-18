package httphandler

import (
	"net/http"
	"strconv"

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

func (h *UserAssetHandler) AddFavorite(ctx *gin.Context) {
	userID, ok := ctxutils.GetUserId(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	// Ambil assetID dari parameter URL
	assetIDParam := ctx.Param("id")
	assetID, err := strconv.ParseInt(assetIDParam, 10, 64)

	if err != nil {
		ctx.Error(err)
		return
	}

	// Tambahkan asset ke daftar favorit
	err = h.userAssetUsecase.AddFavorite(ctx, userID, assetID)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success add asset to favorite"})
}

func (h *UserAssetHandler) RemoveFavorite(ctx *gin.Context) {
	userID, ok := ctxutils.GetUserId(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	// Ambil assetID dari parameter URL
	assetIDParam := ctx.Param("id")
	assetID, err := strconv.ParseInt(assetIDParam, 10, 64)
	if err != nil {
		ctx.Error(err)
		return
	}

	// Hapus asset dari daftar favorit
	err = h.userAssetUsecase.RemoveFavorite(ctx.Request.Context(), userID, assetID)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success remove asset from favorite"})
}
