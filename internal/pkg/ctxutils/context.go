package ctxutils

import (
	"context"

	"github.com/alifnh/bjb-auction-backend/internal/constant"
)

type UserInfo struct {
	ID   int64
	Role string
}

func SetUserInfo(ctx context.Context, userInfo UserInfo) context.Context {
	return context.WithValue(ctx, constant.ContextUserInfoKey, userInfo)
}

func GetUserInfo(ctx context.Context) (UserInfo, bool) {
	val, ok := ctx.Value(constant.ContextUserInfoKey).(UserInfo)
	return val, ok
}

func GetUserId(ctx context.Context) (int64, bool) {
	userInfo, ok := GetUserInfo(ctx)
	return userInfo.ID, ok
}

func GetUserRole(ctx context.Context) (string, bool) {
	userInfo, ok := GetUserInfo(ctx)
	return userInfo.Role, ok
}
