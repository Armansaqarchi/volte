package models

import (
	"flag"
	"fmt"
	"github.com/minio/sha256-simd"
)

var passwordHashAlg = flag.String("password_hash_alg", "sha256", "Password hash algorithm")

const (
	sha256Hash string = "sha256"
)

type User struct {
	Commitment string
	Username   string
	Password   string
}

func (u *User) HashPassword() (string, error) {
	switch *passwordHashAlg {
	case sha256Hash:
		return string(sha256.New().Sum([]byte(u.Password))), nil
	}
	return "", fmt.Errorf(fmt.Sprintf("Unknown password hash algorithm: %s", *passwordHashAlg))
}
