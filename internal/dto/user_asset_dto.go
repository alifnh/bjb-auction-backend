package dto

type FavoriteRequest struct {
	AssetID int64 `json:"asset_id" binding:"required"`
}
