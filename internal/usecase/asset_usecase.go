package usecase

import (
	"context"
	"log"

	"github.com/alifnh/bjb-auction-backend/internal/config"
	"github.com/alifnh/bjb-auction-backend/internal/model"
	"github.com/alifnh/bjb-auction-backend/internal/pkg/apperror"
	"github.com/alifnh/bjb-auction-backend/internal/pkg/database"
	"github.com/alifnh/bjb-auction-backend/internal/pkg/encryptutils"
	"github.com/alifnh/bjb-auction-backend/internal/pkg/jwtutils"
	"github.com/alifnh/bjb-auction-backend/internal/pkg/randutils"
	"github.com/alifnh/bjb-auction-backend/internal/repository"
)

type AssetUsecase interface {
	GetAssetByID(ctx context.Context, id int64) (*model.Asset, error)
}

type assetUsecase struct {
	assetRepository   repository.AssetRepository
	transactor        database.Transactor
	passwordEncryptor encryptutils.PasswordEncryptor
	jwtUtil           jwtutils.JwtUtil
	config            *config.Config
	randutils         randutils.RandomUtil
}

func NewAssetUsecase(ur repository.AssetRepository, transactor database.Transactor, passwordEncryptor encryptutils.PasswordEncryptor, jwtUtil jwtutils.JwtUtil, cfg *config.Config, randutils randutils.RandomUtil) *assetUsecase {
	return &assetUsecase{
		assetRepository:   ur,
		transactor:        transactor,
		passwordEncryptor: passwordEncryptor,
		jwtUtil:           jwtUtil,
		config:            cfg,
		randutils:         randutils,
	}
}

func (u *assetUsecase) GetAssetByID(ctx context.Context, id int64) (*model.Asset, error) {
	asset, err := u.assetRepository.GetAssetById(ctx, id)
	if err != nil {
		log.Printf("failed to get asset by ID: %v", err)
		return nil, apperror.ErrFailedToGetAssetInfo
	}
	return asset, nil
}
