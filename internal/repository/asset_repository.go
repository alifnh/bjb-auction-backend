package repository

import (
	"context"

	"github.com/alifnh/bjb-auction-backend/internal/model"
	"github.com/alifnh/bjb-auction-backend/internal/pkg/database"
)

type AssetRepository interface {
	CreateAsset(ctx context.Context, asset *model.Asset) (*model.Asset, error)
}

type assetRepository struct {
	db *database.PostgresWrapper
}

func NewAssetRepository(db *database.PostgresWrapper) *assetRepository {
	return &assetRepository{db: db}
}

func (r *assetRepository) CreateAsset(ctx context.Context, asset *model.Asset) (*model.Asset, error) {
	query := `INSERT INTO assets (category, img_url, name, price, description, city, address, maps_url, start_date, end_date, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, NOW(), NOW()) RETURNING id`
	err := r.db.Start(ctx).QueryRowContext(ctx, query, asset.Category, asset.ImgUrl, asset.Name, asset.Price, asset.Description, asset.City, asset.Address, asset.MapsUrl, asset.StartDate, asset.EndDate).Scan(&asset.ID)
	if err != nil {
		return nil, err
	}
	return asset, nil
}

func (r *assetRepository) GetAllAssets(ctx context.Context, category string, limit int) ([]*model.Asset, error) {
	query := `SELECT * FROM assets WHERE category = $1 ORDER BY created_at DESC LIMIT $2`
	var assets []*model.Asset
	rows, err := r.db.Start(ctx).QueryContext(ctx, query, category, limit)
	if err != nil {
		return nil, err
	}
	return assets, nil
}
