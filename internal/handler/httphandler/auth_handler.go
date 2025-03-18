package httphandler

import (
	"net/http"

	"github.com/alifnh/bjb-auction-backend/internal/constant"
	"github.com/alifnh/bjb-auction-backend/internal/dto"
	"github.com/alifnh/bjb-auction-backend/internal/pkg/ctxutils"
	"github.com/alifnh/bjb-auction-backend/internal/pkg/ginutils"
	"github.com/alifnh/bjb-auction-backend/internal/usecase"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthHandler(u usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		authUsecase: u,
	}
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	var req dto.RegisterUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		return
	}

	req.ToLower()

	err := h.authUsecase.Register(ctx, &req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ginutils.ResponseSuccessJSON(ctx, http.StatusCreated, constant.ResponseMsgSuccessRegister, nil)
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	var req dto.LoginUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		return
	}

	token, err := h.authUsecase.Login(ctx, &req)
	if err != nil {
		ctx.Error(err)
		return
	}

	data := gin.H{
		"access_token": token,
	}

	ginutils.ResponseSuccessJSON(ctx, http.StatusOK, constant.ResponseMsgSuccessLogin, data)
}

func (h *AuthHandler) GetProfileByID(ctx *gin.Context) {
	userIDInt, _ := ctxutils.GetUserId(ctx)

	user, err := h.authUsecase.GetProfileByID(ctx.Request.Context(), userIDInt)
	if err != nil {
		ctx.Error(err)
		return
	}
	response := dto.EntityToUserResponse(user)

	ginutils.ResponseSuccessJSON(ctx, http.StatusOK, "success get profile", response)
}

func (h *AuthHandler) UpdateProfile(ctx *gin.Context) {
	userID, _ := ctxutils.GetUserId(ctx)

	var req dto.UpdateProfileRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Panggil usecase untuk update profile dan mendapatkan data terbaru
	user, err := h.authUsecase.UpdateProfile(ctx.Request.Context(), userID, &req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ginutils.ResponseSuccessJSON(ctx, http.StatusOK, "success update profile", user)

}
