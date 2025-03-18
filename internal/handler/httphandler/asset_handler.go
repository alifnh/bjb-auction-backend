package httphandler

import (
	"net/http"
	"strconv"

	"os"

	"github.com/alifnh/bjb-auction-backend/internal/constant"
	"github.com/alifnh/bjb-auction-backend/internal/dto"
	"github.com/alifnh/bjb-auction-backend/internal/pkg/cloudinaryutils"
	"github.com/alifnh/bjb-auction-backend/internal/pkg/ginutils"
	"github.com/alifnh/bjb-auction-backend/internal/usecase"
	"github.com/gin-gonic/gin"
)

type AssetHandler struct {
	assetUsecase usecase.AssetUsecase
}

func NewAssetHandler(assetUsecase usecase.AssetUsecase) *AssetHandler {
	return &AssetHandler{assetUsecase: assetUsecase}
}

// Get Asset by ID
func (h *AssetHandler) GetAssetByID(ctx *gin.Context) {

	assetIDParam := ctx.Param("id")
	assetID, err := strconv.ParseInt(assetIDParam, 10, 64)
	if err != nil {
		ctx.Error(err)
		return
	}

	asset, isFavorite, err := h.assetUsecase.GetAssetByID(ctx.Request.Context(), assetID)
	if err != nil {
		ctx.Error(err)
		return
	}
	response := dto.AssetEntityToResponse(asset, isFavorite)
	ginutils.ResponseSuccessJSON(ctx, http.StatusOK, "Succesfully get asset by id", response)

}

func (h *AssetHandler) CreateAsset(ctx *gin.Context) {
	var req dto.CreateAssetRequest

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.Error(err)
		return
	}

	file, err := ctx.FormFile("image_file")
	if err != nil {
		ctx.Error(err)
		return
	}
	err = ctx.SaveUploadedFile(file, file.Filename)
	if err != nil {
		ctx.Error(err)
		return
	}
	img, err := cloudinaryutils.UploadImage(ctx, file.Filename, constant.CloudinaryFolderImageAsset)
	if err != nil {
		ctx.Error(err)
		return
	}
	err = os.Remove(file.Filename)
	if err != nil {
		ctx.Error(err)
		return
	}
	asset, err := h.assetUsecase.CreateAsset(ctx, &req, img)
	if err != nil {
		ctx.Error(err)
		return
	}

	response := dto.EntityToGetAssetResponse(asset)

	ginutils.ResponseSuccessJSON(ctx, http.StatusCreated, "Asset created successfully", response)
}

func (h *AssetHandler) GetAllAssets(ctx *gin.Context) {
	var req dto.GetAssetListRequest

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.Error(err)
		return
	}

	assets, err := h.assetUsecase.GetAllAssets(ctx, req.Category, req.Limit)
	if err != nil {
		ctx.Error(err)
		return
	}
	ginutils.ResponseSuccessJSON(ctx, http.StatusOK, "success get all assets", assets)
}

func (h *AssetHandler) GetAllFavoriteAssets(ctx *gin.Context) {
	var req dto.GetAssetListRequest

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.Error(err)
		return
	}

	assets, err := h.assetUsecase.GetAllFavoriteAssets(ctx, req.Category, req.Limit)
	if err != nil {
		ctx.Error(err)
		return
	}
	ginutils.ResponseSuccessJSON(ctx, http.StatusOK, "success get all favorite assets", assets)
}

func (h *AssetHandler) DeleteAssetByID(ctx *gin.Context) {

	assetIDParam := ctx.Param("id")
	assetID, err := strconv.ParseInt(assetIDParam, 10, 64)
	if err != nil {
		ctx.Error(err)
		return
	}

	err = h.assetUsecase.DeleteAssetByID(ctx.Request.Context(), assetID)
	if err != nil {
		ctx.Error(err)
		return
	}
	ginutils.ResponseSuccessJSON(ctx, http.StatusOK, "Success delete asset", nil)
}

func (h *AssetHandler) UpdateAsset(ctx *gin.Context) {
	var req dto.UpdateAssetRequest
	assetIDParam := ctx.Param("id")
	assetID, err := strconv.ParseInt(assetIDParam, 10, 64)
	if err != nil {
		ctx.Error(err)
		return
	}

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.Error(err)
		return
	}

	err = h.assetUsecase.UpdateAsset(ctx, &req, assetID)
	if err != nil {
		ctx.Error(err)
		return
	}
	ginutils.ResponseSuccessJSON(ctx, http.StatusOK, "success update assets", req)
}
