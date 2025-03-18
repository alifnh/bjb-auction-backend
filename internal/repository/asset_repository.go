package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/alifnh/bjb-auction-backend/internal/model"
	"github.com/alifnh/bjb-auction-backend/internal/pkg/database"
)

type AssetRepository interface {
	GetAssetById(ctx context.Context, id int64) (*model.Asset, error)
	CreateAsset(ctx context.Context, asset *model.Asset) (*model.Asset, error)
	GetAllAssets(ctx context.Context, category string, limit int) ([]*model.Asset, error)
	GetAllAssetsFavorite(ctx context.Context, category string, limit int, userId int64) ([]*model.Asset, error)
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
	log.Print(id)
	var asset model.Asset
	err := r.db.Start(ctx).QueryRowContext(ctx, q, id).Scan(&asset.ID, &asset.Category, &asset.ImgUrl, &asset.Name, &asset.Price,
		&asset.Description, &asset.City, &asset.Address, &asset.MapsUrl,
		&asset.StartDate, &asset.EndDate, &asset.CreatedAt, &asset.UpdatedAt, &asset.DeletedAt)
	if err != nil {
		log.Printf("failed to get assets by id: %v", err)
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &asset, nil
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
	var query string
	if category == "" && limit == 0 {
		query = `SELECT * FROM assets WHERE deleted_at IS NULL ORDER BY created_at`
	} else if category != "" && limit == 0 {
		query = fmt.Sprintf(`SELECT * FROM assets WHERE category = '%s' AND deleted_at IS NULL ORDER BY created_at`, category)
	} else if category == "" && limit > 0 {
		query = fmt.Sprintf(`SELECT * FROM assets WHERE deleted_at IS NULL ORDER BY created_at DESC LIMIT %d`, limit)
	} else if category != "" && limit > 0 {
		query = fmt.Sprintf(`SELECT * FROM assets WHERE category = '%s' AND deleted_at IS NULL ORDER BY created_at DESC LIMIT %d`, category, limit)
	}
	var assets []*model.Asset
	rows, err := r.db.Start(ctx).QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var asset model.Asset
		err := rows.Scan(&asset.ID, &asset.Category, &asset.ImgUrl, &asset.Name, &asset.Price,
			&asset.Description, &asset.City, &asset.Address, &asset.MapsUrl,
			&asset.StartDate, &asset.EndDate, &asset.CreatedAt, &asset.DeletedAt, &asset.UpdatedAt)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}
	return assets, nil
}

func (r *assetRepository) GetAllAssetsFavorite(ctx context.Context, category string, limit int, userId int64) ([]*model.Asset, error) {
	var query string
	if category == "" && limit == 0 {
		query = fmt.Sprintf(`select * from assets a where deleted_at isnull and id in (select ua.asset_id from users_assets ua where ua.user_id = %d) ORDER BY created_at DESC`, userId)
	} else if category != "" && limit == 0 {
		query = fmt.Sprintf(`select * from assets a where category = '%s' and deleted_at isnull and id in (select ua.asset_id from users_assets ua where ua.user_id = %d) ORDER BY created_at DESC`, category, userId)
	} else if category == "" && limit > 0 {
		query = fmt.Sprintf(`select * from assets a where deleted_at isnull and id in (select ua.asset_id from users_assets ua where ua.user_id = %d) ORDER BY created_at DESC LIMIT %d`, userId, limit)
	} else if category != "" && limit > 0 {
		query = fmt.Sprintf(`select * from assets a where category = '%s' and deleted_at isnull and id in (select ua.asset_id from users_assets ua where ua.user_id = %d) ORDER BY created_at DESC LIMIT %d`, category, userId, limit)
	}
	var assets []*model.Asset
	rows, err := r.db.Start(ctx).QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var asset model.Asset
		err := rows.Scan(&asset.ID, &asset.Category, &asset.ImgUrl, &asset.Name, &asset.Price,
			&asset.Description, &asset.City, &asset.Address, &asset.MapsUrl,
			&asset.StartDate, &asset.EndDate, &asset.CreatedAt, &asset.DeletedAt, &asset.UpdatedAt)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}
	return assets, nil
}
