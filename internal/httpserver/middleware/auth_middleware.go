package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/alifnh/bjb-auction-backend/internal/dto"
	"github.com/alifnh/bjb-auction-backend/internal/pkg/ctxutils"
	"github.com/alifnh/bjb-auction-backend/internal/pkg/jwtutils"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	jwtUtil jwtutils.JwtUtil
}

func NewAuthMiddleware(jwtUtil jwtutils.JwtUtil) *AuthMiddleware {
	return &AuthMiddleware{
		jwtUtil: jwtUtil,
	}
}

func (m *AuthMiddleware) RequireToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authStr := c.Request.Header.Get("Authorization")
		if authStr == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
				Message: "please provide the token",
			})
			return
		}

		authStrs := strings.Split(authStr, " ")
		if len(authStrs) != 2 || authStrs[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
				Message: "invalid token - malformed",
			})
			return
		}

		tokenString := authStrs[1]

		claims, err := m.jwtUtil.Parse(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
				Message: fmt.Sprintf("invalid token - %s", err),
			})
		}

		userInfo := ctxutils.UserInfo{
			ID:   claims.UserId,
			Role: claims.Role,
		}
		c.Request = c.Request.WithContext(ctxutils.SetUserInfo(c.Request.Context(), userInfo))
		c.Next()
	}
}
