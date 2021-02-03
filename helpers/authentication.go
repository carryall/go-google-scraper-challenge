package helpers

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword hash given password, returns hashed password and error
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)

		return "", err
	}

	return string(hash), nil
}
