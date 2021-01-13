package helpers

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}
