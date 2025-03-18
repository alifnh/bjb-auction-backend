package model

import "time"

type UsersAssets struct {
	ID        int64
	UserID    int64
	AssetID   int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
