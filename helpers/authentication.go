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

// CompareHashedPasswords compare given hashed password and password
func CompareHashedPasswords(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		log.Println(err)

		return false
	}

	return true
}
