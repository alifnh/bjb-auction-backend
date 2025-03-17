package middleware

import (
	"github.com/alifnh/bjb-auction-backend/internal/pkg/apperror"
	"github.com/alifnh/bjb-auction-backend/internal/pkg/ctxutils"
	"github.com/gin-gonic/gin"
)

type AuthorizationMiddleware struct{}

func NewAuthorizationMiddleware() *AuthorizationMiddleware {
	return &AuthorizationMiddleware{}
}

func (m *AuthorizationMiddleware) RequireRole(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userRole, ok := ctxutils.GetUserRole(ctx.Request.Context())
		if !ok {
			ctx.Error(apperror.ErrUnauthorized)
			ctx.Abort()
			return
		}
		if !containsRole(userRole, roles) {
			ctx.Error(apperror.ErrRoleNotPermitted)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

func containsRole(role string, roles []string) bool {
	for _, r := range roles {
		if role == r {
			return true
		}
	}
	return false
}
