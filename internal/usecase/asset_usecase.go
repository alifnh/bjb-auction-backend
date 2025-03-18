package usecase

import (
	"context"
	"log"

	"github.com/alifnh/bjb-auction-backend/internal/dto"
	"github.com/alifnh/bjb-auction-backend/internal/model"
	"github.com/alifnh/bjb-auction-backend/internal/pkg/apperror"
	"github.com/alifnh/bjb-auction-backend/internal/pkg/ctxutils"
	"github.com/alifnh/bjb-auction-backend/internal/repository"
)

type AssetUsecase interface {
	CreateAsset(ctx context.Context, req *dto.CreateAssetRequest, img string) (*model.Asset, error)
	GetAssetByID(ctx context.Context, id int64) (*model.Asset, bool, error)
	GetAllAssets(ctx context.Context, category string, limit int) ([]*dto.SumAssetResponse, error)
}

type assetUsecase struct {
	assetRepository     repository.AssetRepository
	userAssetRepository repository.UserAssetRepository
}

func NewAssetUsecase(ar repository.AssetRepository, uar repository.UserAssetRepository) *assetUsecase {
	return &assetUsecase{
		assetRepository:     ar,
		userAssetRepository: uar,
	}
}

func (u *assetUsecase) CreateAsset(ctx context.Context, req *dto.CreateAssetRequest, img string) (*model.Asset, error) {

	asset, err := dto.CreateAssetReqToEntity(req, img)
	if err != nil {
		return nil, err
	}
	return u.assetRepository.CreateAsset(ctx, asset)
}

func (u *assetUsecase) GetAssetByID(ctx context.Context, id int64) (*model.Asset, bool, error) {
	asset, err := u.assetRepository.GetAssetById(ctx, id)
	if err != nil {
		log.Printf("failed to get asset by ID: %v", err)
		return nil, false, apperror.ErrFailedToGetAssetInfo
	}
	userId, _ := ctxutils.GetUserId(ctx)
	if u.userAssetRepository == nil {
		log.Println("ERROR: userAssetRepository is nil")
		return nil, false, nil
	}
	isFavorite, err := u.userAssetRepository.IsFavorite(ctx, userId, id)
	if err != nil {
		log.Printf("failed to check favorite status: %v", err)
		return asset, false, nil
	}

	return asset, isFavorite, nil
}

func (u *assetUsecase) GetAllAssets(ctx context.Context, category string, limit int) ([]*dto.SumAssetResponse, error) {
	assets, err := u.assetRepository.GetAllAssets(ctx, category, limit)
	if err != nil {
		log.Printf("failed to get all assets: %v", err)
		return nil, err
	}
	result := dto.ConvertAssetsToSumAssetResponses(assets)
	return result, nil
}
