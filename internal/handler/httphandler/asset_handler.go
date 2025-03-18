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
func (h *AssetHandler) GetAssetByID(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid asset ID"})
		return
	}

	asset, err := h.assetUsecase.GetAssetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, asset)
func NewAssetHandler(u usecase.AssetUsecase) *AssetHandler {
	return &AssetHandler{
		assetUsecase: u,
	}
}

func (h *AssetHandler) CreateAsset(ctx *gin.Context) {
	var req dto.CreateAssetRequest

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file, err := ctx.FormFile("img_file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file upload"})
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
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginutils.ResponseSuccessJSON(ctx, http.StatusCreated, constant.ResponseMsgSuccessRegister, asset)
}
