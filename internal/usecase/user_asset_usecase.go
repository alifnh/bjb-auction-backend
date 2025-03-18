package usecase

import (
	"context"

	"github.com/alifnh/bjb-auction-backend/internal/repository"
)

type UserAssetUsecase interface {
	AddFavorite(ctx context.Context, userID int64, assetID int64) error
	RemoveFavorite(ctx context.Context, userID int64, assetID int64) error
}

type userAssetUsecase struct {
	userAssetRepository repository.UserAssetRepository
}

func NewUserAssetUsecase(uar repository.UserAssetRepository) *userAssetUsecase {
	return &userAssetUsecase{
		userAssetRepository: uar,
	}
}

func (u *userAssetUsecase) AddFavorite(ctx context.Context, userID int64, assetID int64) error {
	return u.userAssetRepository.AddFavorite(ctx, userID, assetID)
}

func (u *userAssetUsecase) RemoveFavorite(ctx context.Context, userID int64, assetID int64) error {
	return u.userAssetRepository.RemoveFavorite(ctx, userID, assetID)
}
