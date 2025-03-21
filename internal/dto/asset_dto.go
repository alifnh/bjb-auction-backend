package dto

import (
	"github.com/alifnh/bjb-auction-backend/internal/model"
	"github.com/alifnh/bjb-auction-backend/internal/pkg/dateutils"
)

// type CreateAssetRequest struct {
// 	Category    string  `json:"category" binding:"required"`
// 	Img_url     string  `json:"img_url" binding:"required"`
// 	Name        string  `json:"name" binding:"required"`
// 	Price       float64 `json:"price" binding:"required"`
// 	Description string  `json:"description"`
// 	City        string  `json:"city" binding:"required"`
// 	Address     string  `json:"address" binding:"required"`
// 	Maps_url    string  `json:"maps_url"`
// 	Start_date  string  `json:"start_date" binding:"required"`
// 	End_date    string  `json:"end_date" binding:"required"`
// }

// func CreateAssetReqToEntity(r *CreateAssetRequest) (*model.Asset, error) {
// 	startDate, err := dateutils.DateToTimestamp(r.Start_date)
// 	if err != nil {
// 		return nil, err
// 	}
// 	endDate, err := dateutils.DateToTimestamp(r.End_date)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &model.Asset{
// 		Category:    r.Category,
// 		Img_url:     r.Img_url,
// 		Name:        r.Name,
// 		Price:       r.Price,
// 		Description: &r.Description,
// 		City:        r.City,
// 		Address:     r.Address,
// 		Maps_url:    &r.Maps_url,
// 		Start_date:  startDate,
// 		End_date:    endDate,
// 	}, nil
// }

type CreateAssetRequest struct {
	Category    string  `form:"category" binding:"required"`
	Name        string  `form:"name" binding:"required"`
	Price       float64 `form:"price" binding:"required"`
	Description string  `form:"description"`
	City        string  `form:"city" binding:"required"`
	Address     string  `form:"address" binding:"required"`
	MapsURL     string  `form:"maps_url"`
	StartDate   string  `form:"start_date" binding:"required"`
	EndDate     string  `form:"end_date" binding:"required"`
}

type UpdateAssetRequest struct {
	Category    string  `json:"category" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Description string  `json:"description"`
	City        string  `json:"city" binding:"required"`
	Address     string  `json:"address" binding:"required"`
	MapsURL     string  `json:"maps_url"`
	StartDate   string  `json:"start_date" binding:"required"`
	EndDate     string  `json:"end_date" binding:"required"`
}

func CreateAssetReqToEntity(r *CreateAssetRequest, img string) (*model.Asset, error) {
	startDate, err := dateutils.DateToTimestamp(r.StartDate)
	if err != nil {
		return nil, err
	}
	endDate, err := dateutils.DateToTimestamp(r.EndDate)
	if err != nil {
		return nil, err
	}

	return &model.Asset{
		Category:    r.Category,
		ImgUrl:      img,
		Name:        r.Name,
		Price:       r.Price,
		Description: &r.Description,
		City:        r.City,
		Address:     r.Address,
		MapsUrl:     &r.MapsURL,
		StartDate:   startDate,
		EndDate:     endDate,
	}, nil
}

type GetAssetResponse struct {
	Category    string  `json:"category"`
	Img_url     string  `json:"img_url"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	City        string  `json:"city"`
	Address     string  `json:"address"`
	MapsUrl     string  `json:"maps_url"`
	StartDate   string  `json:"start_date"`
	EndDate     string  `json:"end_date"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

func EntityToGetAssetResponse(asset *model.Asset) *GetAssetResponse {
	startDate, err := dateutils.TimestampToDate(asset.StartDate)
	if err != nil {
		return nil
	}
	endDate, err := dateutils.TimestampToDate(asset.EndDate)
	if err != nil {
		return nil
	}
	createdAt, err := dateutils.TimestampToDateTime(asset.CreatedAt)
	if err != nil {
		return nil
	}
	updatedAt, err := dateutils.TimestampToDateTime(asset.UpdatedAt)
	if err != nil {
		return nil
	}

	return &GetAssetResponse{
		Category:    asset.Category,
		Img_url:     asset.ImgUrl,
		Name:        asset.Name,
		Price:       asset.Price,
		Description: *asset.Description,
		City:        asset.City,
		Address:     asset.Address,
		MapsUrl:     *asset.MapsUrl,
		StartDate:   startDate,
		EndDate:     endDate,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}

// Struktur untuk response ketika mengambil data asset
type AssetResponse struct {
	ID          int64   `json:"id"`
	Category    string  `json:"category"`
	ImgURL      string  `json:"img_url"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description,omitempty"`
	City        string  `json:"city"`
	Address     string  `json:"address"`
	MapsURL     string  `json:"maps_url,omitempty"`
	StartDate   string  `json:"start_date"`
	EndDate     string  `json:"end_date"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	IsFavorite  bool    `json:"is_favorite"`
}

func AssetEntityToResponse(asset *model.Asset, isFavorite bool) *AssetResponse {
	startDate, err := dateutils.TimestampToDate(asset.StartDate)
	if err != nil {
		return nil
	}
	endDate, err := dateutils.TimestampToDate(asset.EndDate)
	if err != nil {
		return nil
	}
	createdAt, err := dateutils.TimestampToDateTime(asset.CreatedAt)
	if err != nil {
		return nil
	}
	updatedAt, err := dateutils.TimestampToDateTime(asset.UpdatedAt)
	if err != nil {
		return nil
	}

	var MapsURL string
	if asset.MapsUrl != nil {
		MapsURL = *asset.MapsUrl
	}
	var Description string
	if asset.Description != nil {
		Description = *asset.Description
	}

	return &AssetResponse{
		ID:          asset.ID,
		Category:    asset.Category,
		ImgURL:      asset.ImgUrl,
		Name:        asset.Name,
		Price:       asset.Price,
		Description: Description,
		City:        asset.City,
		Address:     asset.Address,
		MapsURL:     MapsURL,
		StartDate:   startDate,
		EndDate:     endDate,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		IsFavorite:  isFavorite,
	}
}

type GetAssetListRequest struct {
	Limit    int    `form:"limit" binding:"numeric,gte=0"`
	Category string `form:"category"`
}

type SumAssetResponse struct {
	Id        int64   `json:"id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	ImgUrl    string  `json:"img_url"`
	City      string  `json:"city"`
	Category  string  `json:"category"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

type GetAllAssetsResponse struct {
	Data []*SumAssetResponse
}

func ConvertAssetsToSumAssetResponses(assets []*model.Asset) []*SumAssetResponse {
	var response []*SumAssetResponse

	for _, asset := range assets {
		response = append(response, &SumAssetResponse{
			Id:        asset.ID,
			Name:      asset.Name,
			Price:     asset.Price,
			ImgUrl:    asset.ImgUrl,
			City:      asset.City,
			Category:  asset.Category,
			CreatedAt: asset.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: asset.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return response
}

func UpdateAssetReqToEntity(r *UpdateAssetRequest) (*model.Asset, error) {
	startDate, err := dateutils.DateToTimestamp(r.StartDate)
	if err != nil {
		return nil, err
	}
	endDate, err := dateutils.DateToTimestamp(r.EndDate)
	if err != nil {
		return nil, err
	}

	return &model.Asset{
		Category:    r.Category,
		Name:        r.Name,
		Price:       r.Price,
		Description: &r.Description,
		City:        r.City,
		Address:     r.Address,
		MapsUrl:     &r.MapsURL,
		StartDate:   startDate,
		EndDate:     endDate,
	}, nil
}
