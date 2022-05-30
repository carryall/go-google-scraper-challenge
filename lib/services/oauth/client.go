package oauth

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/spf13/viper"
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
		Domain: fmt.Sprintf("%s:%s", viper.GetString("APP_HOST"), viper.GetString("PORT")),
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
