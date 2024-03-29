package controllers

import (
	"net/http"

	"go-google-scraper-challenge/errors"
	. "go-google-scraper-challenge/helpers/api"
	helpers "go-google-scraper-challenge/helpers/api"
	"go-google-scraper-challenge/lib/api/v1/controllers"
	"go-google-scraper-challenge/lib/api/v1/serializers"
	"go-google-scraper-challenge/lib/services/oauth"

	"github.com/gin-gonic/gin"
)

type OAuthClientsController struct {
	controllers.BaseController
}

func (c *OAuthClientsController) Create(ctx *gin.Context) {
	oauthClient, err := oauth.GenerateClient()

	if err != nil {
		RenderJSONError(ctx, errors.ErrUnProcessableEntity, err.Error())
	}

	response := &serializers.OAuthClientResponse{
		ID:           oauthClient.ClientID,
		ClientID:     oauthClient.ClientID,
		ClientSecret: oauthClient.ClientSecret,
	}

	helpers.RenderJSON(ctx, http.StatusOK, response)
}
