package model

import (
	"time"
)

type User struct {
	ID          int64
	Name        string
	Password    string
	Email       string
	Role        string
	PhoneNumber *string
	Nik         string
	City        string
	Gender      string
	DateOfBirth time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
