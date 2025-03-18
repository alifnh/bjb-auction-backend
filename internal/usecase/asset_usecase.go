package usecase

import (
	"context"

	"github.com/alifnh/bjb-auction-backend/internal/dto"
	"github.com/alifnh/bjb-auction-backend/internal/model"
	"github.com/alifnh/bjb-auction-backend/internal/repository"
)

type AssetUsecase interface {
	CreateAsset(ctx context.Context, req *dto.CreateAssetRequest, img string) (*model.Asset, error)
}

type assetUsecase struct {
	assetRepository repository.AssetRepository
}

func NewAssetUsecase(ar repository.AssetRepository) *assetUsecase {
	return &assetUsecase{
		assetRepository: ar,
	}
}

func (u *assetUsecase) CreateAsset(ctx context.Context, req *dto.CreateAssetRequest, img string) (*model.Asset, error) {

	asset, err := dto.CreateAssetReqToEntity(req, img)
	if err != nil {
		return nil, err
	}
	return u.assetRepository.CreateAsset(ctx, asset)
}
