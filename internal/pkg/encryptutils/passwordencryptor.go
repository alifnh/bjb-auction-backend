package encryptutils

import (
	"encoding/base64"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type PasswordEncryptor interface {
	Hash(password string) (string, error)
	Check(password, hash string) bool
	GenerateResetPasswordToken(email string) string
	SetTokenExpiry(duration time.Duration) time.Time
}

type bcryptPasswordEncryptor struct {
	cost int
}

func NewBcryptPasswordEncryptor(cost int) *bcryptPasswordEncryptor {
	return &bcryptPasswordEncryptor{
		cost: cost,
	}
}

func (e *bcryptPasswordEncryptor) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), e.cost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (e *bcryptPasswordEncryptor) Check(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (e *bcryptPasswordEncryptor) GenerateResetPasswordToken(email string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(email), e.cost)
	if err != nil {
		return err.Error()
	}

	return base64.StdEncoding.EncodeToString(hash)
}

func (e *bcryptPasswordEncryptor) SetTokenExpiry(duration time.Duration) time.Time {
	return time.Now().Add(duration)
}
