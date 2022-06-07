package helpers

import (
	"encoding/json"
	"go-google-scraper-challenge/helpers/log"
	"go-google-scraper-challenge/lib/api/v1/serializers"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword hash given password, returns hashed password and error
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Info(err)

		return "", err
	}

	return string(hash), nil
}

// CompareHashWithPassword compare given hashed password and password
func CompareHashWithPassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		log.Info(err)

		return false
	}

	return true
}

func GetTokenInfo(tokenData map[string]interface{}) (tokenInfo *serializers.AuthenticationToken, err error) {
	tokenInfo = &serializers.AuthenticationToken{}
	tokenDataString, err := json.Marshal(tokenData)
	if err != nil {
		return
	}

	err = json.Unmarshal(tokenDataString, tokenInfo)

	return
}
