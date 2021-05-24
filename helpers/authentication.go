package helpers

import (
	"github.com/beego/beego/v2/core/logs"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hash given password, returns hashed password and error
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		logs.Info(err)

		return "", err
	}

	return string(hash), nil
}

// CompareHashWithPassword compare given hashed password and password
func CompareHashWithPassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		logs.Info(err)

		return false
	}

	return true
}
