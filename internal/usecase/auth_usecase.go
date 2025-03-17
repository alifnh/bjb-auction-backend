package usecase

import (
	"context"
	"log"

	"github.com/alifnh/bjb-auction-backend/internal/config"
	"github.com/alifnh/bjb-auction-backend/internal/dto"
	"github.com/alifnh/bjb-auction-backend/internal/pkg/apperror"
	"github.com/alifnh/bjb-auction-backend/internal/pkg/database"
	"github.com/alifnh/bjb-auction-backend/internal/pkg/encryptutils"
	"github.com/alifnh/bjb-auction-backend/internal/pkg/jwtutils"
	"github.com/alifnh/bjb-auction-backend/internal/pkg/randutils"
	"github.com/alifnh/bjb-auction-backend/internal/repository"
)

type AuthUsecase interface {
	Register(ctx context.Context, req *dto.RegisterUserRequest) error
	Login(ctx context.Context, req *dto.LoginUserRequest) (string, error)
}

type authUsecase struct {
	authRepository    repository.AuthRepository
	transactor        database.Transactor
	passwordEncryptor encryptutils.PasswordEncryptor
	jwtUtil           jwtutils.JwtUtil
	config            *config.Config
	randutils         randutils.RandomUtil
}

func NewAuthUsecase(ur repository.AuthRepository, transactor database.Transactor, passwordEncryptor encryptutils.PasswordEncryptor, jwtUtil jwtutils.JwtUtil, cfg *config.Config, randutils randutils.RandomUtil) *authUsecase {
	return &authUsecase{
		authRepository:    ur,
		transactor:        transactor,
		passwordEncryptor: passwordEncryptor,
		jwtUtil:           jwtUtil,
		config:            cfg,
		randutils:         randutils,
	}
}

func (u *authUsecase) Register(ctx context.Context, req *dto.RegisterUserRequest) error {
	return u.transactor.Transaction(ctx, func(txCtx context.Context) error {
		user, err := dto.RegisterReqToUserEntity(req)
		if err != nil {
			return err
		}
		exists, err := u.authRepository.UserExists(txCtx, user.Email)
		if err != nil {
			log.Printf("failed to check if user exists: %v", err)
			return apperror.ErrFailedToRegisterUser
		}
		if exists {
			return apperror.ErrEmailAlreadyRegistered
		}
		hashedPassword, err := u.passwordEncryptor.Hash(user.Password)
		if err != nil {
			log.Printf("failed to hash password: %v", err)
			return apperror.ErrFailedToHashPassword
		}
		user.Password = hashedPassword
		user.Role = "user"
		_, err = u.authRepository.Register(txCtx, user)
		if err != nil {
			log.Printf("failed to register user: %v", err)
			return apperror.ErrFailedToRegisterUser
		}
		return nil
	})
}

func (u *authUsecase) Login(ctx context.Context, req *dto.LoginUserRequest) (string, error) {
	user, err := u.authRepository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		log.Printf("failed to get user by email: %v", err)
		return "", apperror.ErrFailedToLogin
	}

	if user == nil {
		return "", apperror.ErrInvalidCredentials
	}

	isMatch := u.passwordEncryptor.Check(req.Password, user.Password)
	if !isMatch {
		return "", apperror.ErrInvalidCredentials
	}

	token, err := u.jwtUtil.Sign(user.ID, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}
