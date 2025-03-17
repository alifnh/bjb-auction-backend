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
