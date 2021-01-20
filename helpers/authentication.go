package helpers

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// EncryptPassword encrypt given password, returns encrypted password and error
func EncryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)

		return "", err
	}

	return string(hash), nil
}
