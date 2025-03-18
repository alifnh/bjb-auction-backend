package dto

import (
	"github.com/alifnh/bjb-auction-backend/internal/model"
	"github.com/alifnh/bjb-auction-backend/internal/pkg/dateutils"
)

// Struktur untuk response ketika mengambil data asset
type AssetResponse struct {
	ID          int64   `json:"id"`
	Category    string  `json:"category"`
	ImgURL      string  `json:"img_url"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	City        string  `json:"city"`
	Address     string  `json:"address"`
	MapsURL     string  `json:"maps_url"`
	StartDate   string  `json:"start_date"`
	EndDate     string  `json:"end_date"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

func AssetEntityToResponse(asset *model.Asset) *AssetResponse {
	return &AssetResponse{
		ID:          asset.ID,
		Category:    asset.Category,
		ImgURL:      asset.ImgUrl,
		Name:        asset.Name,
		Price:       asset.Price,
		Description: *asset.Description,
		City:        asset.City,
		Address:     asset.Address,
		MapsURL:     *asset.MapsUrl,
		StartDate:   dateutils.TimestampToDate(asset.StartDate),
		EndDate:     dateutils.TimestampToDate(asset.EndDate),
		CreatedAt:   dateutils.TimestampToDate(asset.CreatedAt),
		UpdatedAt:   dateutils.TimestampToDate(asset.UpdatedAt),
	}
}
