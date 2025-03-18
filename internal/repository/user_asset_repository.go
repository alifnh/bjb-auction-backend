package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/alifnh/bjb-auction-backend/internal/pkg/database"
)

type UserAssetRepository interface {
	AddFavorite(ctx context.Context, userID int64, assetID int64) error
	RemoveFavorite(ctx context.Context, userID int64, assetID int64) error
	IsFavorite(ctx context.Context, userID int64, assetID int64) (bool, error)
}

type userAssetRepository struct {
	db *database.PostgresWrapper
}

func NewUserAssetRepository(db *database.PostgresWrapper) *userAssetRepository {
	return &userAssetRepository{db: db}
}

func (r *userAssetRepository) AddFavorite(ctx context.Context, userID int64, assetID int64) error {
	query := `
		INSERT INTO users_assets (user_id, asset_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (user_id, asset_id) DO NOTHING;
	`

	now := time.Now()
	_, err := r.db.Start(ctx).ExecContext(ctx, query, userID, assetID, now, now)
	if err != nil {
		return fmt.Errorf("failed to add favorite asset: %w", err)
	}

	return nil
}

// Menghapus Asset dari Favorite
func (r *userAssetRepository) RemoveFavorite(ctx context.Context, userID int64, assetID int64) error {
	query := `DELETE FROM users_assets WHERE user_id = $1 AND asset_id = $2`
	res, err := r.db.Start(ctx).ExecContext(ctx, query, userID, assetID)
	if err != nil {
		return fmt.Errorf("failed to remove favorite asset: %w", err)
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("favorite asset not found")
	}

	return nil
}

// Mengecek Apakah Asset Sudah Difavoritkan oleh User
func (r *userAssetRepository) IsFavorite(ctx context.Context, userID int64, assetID int64) (bool, error) {
	log.Println("555555555555555555555555555555555")
	query := `SELECT EXISTS(SELECT 1 FROM users_assets WHERE user_id = $1 AND asset_id = $2)`
	log.Println(query)
	var exists bool
	err := r.db.Start(ctx).QueryRowContext(ctx, query, userID, assetID).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check favorite status: %w", err)
	}
	return exists, nil
}
