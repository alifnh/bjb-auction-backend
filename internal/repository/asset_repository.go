package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/alifnh/bjb-auction-backend/internal/model"
	"github.com/alifnh/bjb-auction-backend/internal/pkg/database"
)

type AssetRepository interface {
	GetAssetById(ctx context.Context, id int64) (*model.Asset, error)
	CreateAsset(ctx context.Context, asset *model.Asset) (*model.Asset, error)
}

type assetRepository struct {
	db *database.PostgresWrapper
}

func NewAssetRepository(db *database.PostgresWrapper) *assetRepository {
	return &assetRepository{db: db}
}

func (r *assetRepository) GetAssetById(ctx context.Context, id int64) (*model.Asset, error) {
	q := `
		SELECT id, category, img_url, name, price, description, city, address, maps_url, 
		       start_date, end_date, created_at, updated_at, deleted_at
		FROM assets
		WHERE id = $1
	`
	var asset model.Asset
	err := r.db.Start(ctx).QueryRowContext(ctx, q, id).Scan(&asset.ID, &asset.Category, &asset.ImgUrl, &asset.Name, &asset.Price,
		&asset.Description, &asset.City, &asset.Address, &asset.MapsUrl,
		&asset.StartDate, &asset.EndDate, &asset.CreatedAt, &asset.UpdatedAt, &asset.DeletedAt)
	if err != nil {
		log.Printf("failed to get user by id: %v", err)
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &asset, nil
func (r *assetRepository) CreateAsset(ctx context.Context, asset *model.Asset) (*model.Asset, error) {
	query := `INSERT INTO assets (category, img_url, name, price, description, city, address, maps_url, start_date, end_date, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, NOW(), NOW()) RETURNING id`
	err := r.db.Start(ctx).QueryRowContext(ctx, query, asset.Category, asset.ImgUrl, asset.Name, asset.Price, asset.Description, asset.City, asset.Address, asset.MapsUrl, asset.StartDate, asset.EndDate).Scan(&asset.ID)
	if err != nil {
		return nil, err
	}
	return asset, nil
}
