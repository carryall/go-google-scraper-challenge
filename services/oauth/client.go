package oauth

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	"gopkg.in/oauth2.v3/models"
)

type OAuthClient struct {
	ClientID     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
}

func GenerateClient() (client OAuthClient, err error) {
	clientID := uuid.New().String()
	clientSecret := uuid.New().String()

	err = clientStore.Create(&models.Client{
		ID:     clientID,
		Secret: clientSecret,
		Domain: fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("PORT")),
	})
	if err != nil {
		return OAuthClient{}, err
	}

	client = OAuthClient{
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}

	return client, nil
}
