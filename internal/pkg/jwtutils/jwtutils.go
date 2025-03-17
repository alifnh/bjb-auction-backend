package jwtutils

import (
	"errors"
	"time"

	"github.com/alifnh/bjb-auction-backend/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

type JwtUtil interface {
	Sign(userId int64, role string) (string, error)
	Parse(tokenString string) (*MyAuthClaims, error)
}

type jwtUtil struct {
	config config.JwtConfig
}

func NewJwtUtil(jwtConfig config.JwtConfig) *jwtUtil {
	return &jwtUtil{
		config: jwtConfig,
	}
}

type MyAuthClaims struct {
	UserId int64  `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func (h *jwtUtil) Sign(userId int64, role string) (string, error) {
	currentTime := time.Now()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyAuthClaims{
		UserId: userId,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(currentTime),
			ExpiresAt: jwt.NewNumericDate(currentTime.Add(time.Duration(h.config.TokenDuration) * time.Minute)),
			Issuer:    h.config.Issuer,
		},
	})

	s, err := token.SignedString([]byte(h.config.SecretKey))
	if err != nil {
		return "", err
	}

	return s, nil
}

func (h *jwtUtil) Parse(tokenString string) (*MyAuthClaims, error) {
	res := MyAuthClaims{}

	parser := jwt.NewParser(
		jwt.WithValidMethods(h.config.AllowedAlgs),
		jwt.WithIssuer(h.config.Issuer),
		jwt.WithIssuedAt(),
	)

	token, err := parser.ParseWithClaims(tokenString, &MyAuthClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(h.config.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*MyAuthClaims); ok && token.Valid {
		res = *claims
	} else {
		return nil, errors.New("token not valid")
	}

	return &res, nil
}
