package apperror

import (
	"net/http"

	"github.com/alifnh/bjb-auction-backend/internal/constant"
)

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (ce *AppError) Error() string {
	return ce.Message
}

var (
	ErrInternalServerError           = &AppError{http.StatusInternalServerError, constant.InternalServerErrorMessage}
	ErrFailedToRegisterUser          = &AppError{http.StatusInternalServerError, constant.FailedToRegister}
	ErrEmailAlreadyRegistered        = &AppError{http.StatusBadRequest, constant.UserAlreadyRegisteredError}
	ErrFailedToHashPassword          = &AppError{http.StatusInternalServerError, constant.FailedToHashPasword}
	ErrFailedToSaveVerificationToken = &AppError{http.StatusInternalServerError, constant.FailedToSaveVerificationToken}
	ErrInvalidVerificationToken      = &AppError{http.StatusBadRequest, constant.InvalidVerificationToken}
	ErrFailedToSendVerificationEmail = &AppError{http.StatusInternalServerError, constant.FailedToSendVerificationEmail}
	ErrPasswordNotMatch              = &AppError{http.StatusBadRequest, constant.PasswordNotMatch}
	ErrInvalidCredentials            = &AppError{http.StatusBadRequest, constant.InvalidCredentials}
	ErrFailedToLogin                 = &AppError{http.StatusInternalServerError, constant.FailedToLogin}
	ErrFailedToForgotPassword        = &AppError{http.StatusInternalServerError, constant.FailedToForgotPassword}
	ErrFailedToResetPassword         = &AppError{http.StatusInternalServerError, constant.FailedToResetPassword}
	ErrInvalidOrExpiredToken         = &AppError{http.StatusBadRequest, constant.InvalidOrExpiredToken}
	ErrUnauthorized                  = &AppError{http.StatusUnauthorized, constant.Unauthorized}
	ErrUserAlreadyVerified           = &AppError{http.StatusBadRequest, constant.UserAlreadyVerified}
	ErrFailedToExchangeToken         = &AppError{http.StatusInternalServerError, constant.FailedToExchangeToken}
	ErrFailedToGetUserInfo           = &AppError{http.StatusInternalServerError, constant.FailedToGetUserInfo}
	ErrFailedToDecodeUserInfo        = &AppError{http.StatusInternalServerError, constant.FailedToDecodeUserInfo}
	ErrRoleNotPermitted              = &AppError{http.StatusBadRequest, constant.RoleNotPermitted}
	ErrFailedToGetAssetInfo          = &AppError{http.StatusInternalServerError, constant.FailedToGetAssetInfo}
	ErrFailedToCreateAsset           = &AppError{http.StatusInternalServerError, constant.FailedToCreateAsset}
)
