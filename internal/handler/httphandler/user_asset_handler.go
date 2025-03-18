package httphandler

import (
	"net/http"
	"strconv"

	"github.com/alifnh/bjb-auction-backend/internal/pkg/ctxutils"
	"github.com/alifnh/bjb-auction-backend/internal/pkg/ginutils"
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
	userID, _ := ctxutils.GetUserId(ctx)

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

	ginutils.ResponseSuccessJSON(ctx, http.StatusOK, "success add asset to favorite", nil)
}

func (h *UserAssetHandler) RemoveFavorite(ctx *gin.Context) {
	userID, _ := ctxutils.GetUserId(ctx)

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
	ginutils.ResponseSuccessJSON(ctx, http.StatusOK, "success remove asset from favorite", nil)
}
