package randutils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type RandomUtil interface {
	StringAlphaNum(len int) string
	GenerateVerificationToken() string
}

type stdLibRandomUtil struct{}

func NewStdLibRandomUtil() *stdLibRandomUtil {
	// rand instance can be improved by creating a new instance instead of using the package level one
	// or even better, use the cryptographically safe random generator
	rand.Seed(time.Now().Unix())

	return &stdLibRandomUtil{}
}

var alphanumSet = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func (u *stdLibRandomUtil) StringAlphaNum(n int) string {
	var res strings.Builder

	lenCharSet := len(alphanumSet)

	for i := 0; i < n; i++ {
		res.WriteByte(alphanumSet[rand.Intn(lenCharSet)])
	}

	return res.String()
}

func (u *stdLibRandomUtil) GenerateVerificationToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
