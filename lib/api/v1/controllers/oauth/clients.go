package controllers

import (
	"net/http"

	. "go-google-scraper-challenge/helpers/api"
	"go-google-scraper-challenge/lib/api/v1/controllers"
	"go-google-scraper-challenge/lib/services/oauth"

	"github.com/gin-gonic/gin"
)

type OAuthClientsController struct {
	controllers.BaseController
}

func (c *OAuthClientsController) Create(ctx *gin.Context) {
	oauthClient, err := oauth.GenerateClient()

	if err != nil {
		ResponseWithError(ctx, http.StatusUnprocessableEntity, err)
	}

	ctx.JSON(http.StatusOK, oauthClient)
}
