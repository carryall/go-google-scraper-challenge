package helpers

import (
	"encoding/json"

	"go-google-scraper-challenge/helpers/log"

	"golang.org/x/crypto/bcrypt"
)

type AuthenticationToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
	TokenType    string `json:"token_type"`
}

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

func GetTokenInfo(tokenData map[string]interface{}) (tokenInfo *AuthenticationToken, err error) {
	tokenInfo = &AuthenticationToken{}
	tokenDataString, err := json.Marshal(tokenData)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(tokenDataString, tokenInfo)
	if err != nil {
		return nil, err
	}

	return tokenInfo, nil
}
