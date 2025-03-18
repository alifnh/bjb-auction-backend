package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/alifnh/bjb-auction-backend/internal/dto"
	"github.com/alifnh/bjb-auction-backend/internal/model"
	"github.com/alifnh/bjb-auction-backend/internal/pkg/database"
)

type AuthRepository interface {
	Register(ctx context.Context, user *model.User) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	GetById(ctx context.Context, id int64) (*model.User, error)
	UserExists(ctx context.Context, email string) (bool, error)
	UpdateProfile(ctx context.Context, userID int64, req *dto.UpdateProfileRequest) error
}

type authRepository struct {
	db *database.PostgresWrapper
}

func NewAuthRepository(db *database.PostgresWrapper) *authRepository {
	return &authRepository{db: db}
}

func (r *authRepository) Register(ctx context.Context, user *model.User) (*model.User, error) {
	query := `INSERT INTO users (email, password, name, nik, gender, date_of_birth, role, city, phone_number, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6, 'user', $7, $8, NOW(), NOW()) RETURNING id`
	err := r.db.Start(ctx).QueryRowContext(ctx, query, user.Email, user.Password, user.Name, user.Nik, user.Gender, user.DateOfBirth, user.City, user.PhoneNumber).Scan(&user.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *authRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	query := `SELECT id, email, role, password FROM users WHERE email = $1 and deleted_at isnull`
	var user model.User
	err := r.db.Start(ctx).QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Email, &user.Role, &user.Password)

	if err != nil {
		log.Printf("failed to get user by email: %v", err)
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *authRepository) GetById(ctx context.Context, id int64) (*model.User, error) {
	q := `select id, email, password, name, nik, gender, phone_number, date_of_birth, role, city, created_at, updated_at from users where id = $1 and deleted_at isnull `
	var user model.User
	err := r.db.Start(ctx).QueryRowContext(ctx, q, id).Scan(&user.ID, &user.Email, &user.Password, &user.Name, &user.Nik, &user.Gender, &user.PhoneNumber, &user.DateOfBirth, &user.Role, &user.City, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		log.Printf("failed to get user by id: %v", err)
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *authRepository) UserExists(ctx context.Context, email string) (bool, error) {
	query := `SELECT EXISTS (SELECT 1 FROM users WHERE email = $1 and deleted_at isnull)`
	var exists bool
	err := r.db.Start(ctx).QueryRowContext(ctx, query, email).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (r *authRepository) UpdateProfile(ctx context.Context, userID int64, req *dto.UpdateProfileRequest) error {
	query := `
		UPDATE users
		SET name = $1, email = $2, phone_number = $3, nik = $4, 
		    date_of_birth = $5, gender = $6, city = $7, updated_at = NOW()
		WHERE id = $8
	`

	_, err := r.db.Start(ctx).ExecContext(ctx, query,
		req.Name, req.Email, req.PhoneNumber, req.Nik,
		req.DateOfBirth, req.Gender, req.City, userID,
	)
	if err != nil {
		return err
	}

	return nil
}
