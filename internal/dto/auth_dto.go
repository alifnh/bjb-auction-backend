package dto

import (
	"strings"

	"github.com/alifnh/bjb-auction-backend/internal/model"
	"github.com/alifnh/bjb-auction-backend/internal/pkg/dateutils"
)

type RegisterUserRequest struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,alphanum"`
	PhoneNumber string `json:"phone_number" binding:"omitempty"`
	Nik         string `json:"nik" binding:"required"`
	City        string `json:"city" binding:"required"`
	Gender      string `json:"gender" binding:"required"`
	DateOfBirth string `json:"date_of_birth" binding:"required"`
}

func (req *RegisterUserRequest) ToLower() {
	req.Email = strings.ToLower(req.Email)
}

func RegisterReqToUserEntity(r *RegisterUserRequest) (*model.User, error) {
	dob, err := dateutils.DateToTimestamp(r.DateOfBirth)
	if err != nil {
		return nil, err
	}

	return &model.User{
		Name:        r.Name,
		Email:       r.Email,
		Password:    r.Password,
		PhoneNumber: &r.PhoneNumber,
		Nik:         r.Nik,
		City:        r.City,
		Gender:      r.Gender,
		DateOfBirth: dob,
	}, nil
}

type LoginUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,alphanum"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

func LoginReqToUserEntity(r *LoginUserRequest) *model.User {
	return &model.User{
		Email:    r.Email,
		Password: r.Password,
	}
}

func EntityToUserResponse(user *model.User) *UserResponse {
	DateOfBirth, err := dateutils.TimestampToDate(user.DateOfBirth)
	if err != nil {
		return nil
	}
	createdAt, err := dateutils.TimestampToDateTime(user.CreatedAt)
	if err != nil {
		return nil
	}
	updatedAt, err := dateutils.TimestampToDateTime(user.UpdatedAt)
	if err != nil {
		return nil
	}

	return &UserResponse{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Nik:         user.Nik,
		City:        user.City,
		Gender:      user.Gender,
		DateOfBirth: DateOfBirth,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}

type UpdateProfileRequest struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phone_number"`
	Nik         string `json:"nik" binding:"required"`
	DateOfBirth string `json:"date_of_birth" binding:"required"`
	Gender      string `json:"gender" binding:"required"`
	City        string `json:"city" binding:"required"`
}

type UserResponse struct {
	ID          int64   `json:"ID"`
	Name        string  `json:"Name"`
	Email       string  `json:"Email"`
	PhoneNumber *string `json:"PhoneNumber"`
	Nik         string  `json:"Nik"`
	City        string  `json:"City"`
	Gender      string  `json:"Gender"`
	DateOfBirth string  `json:"DateOfBirth"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
}
