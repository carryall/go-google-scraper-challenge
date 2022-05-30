package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"go-google-scraper-challenge/constants"
	. "go-google-scraper-challenge/helpers/api"
	"go-google-scraper-challenge/lib/api/v1/forms"
	"go-google-scraper-challenge/lib/api/v1/serializers"
	"go-google-scraper-challenge/lib/models"
	"go-google-scraper-challenge/lib/services/oauth"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type UsersController struct {
	BaseController
}

func (c *UsersController) Register(ctx *gin.Context) {
	registrationForm := &forms.RegistrationForm{}

	err := ctx.ShouldBindWith(registrationForm, binding.Form)
	if err != nil {
		ResponseWithError(ctx, http.StatusBadRequest, err)
		return
	}

	_, err = registrationForm.Validate()
	if err != nil {
		ResponseWithError(ctx, http.StatusBadRequest, err)
		return
	}

	userID, err := registrationForm.Save()
	if err != nil {
		ResponseWithError(ctx, http.StatusUnprocessableEntity, err)
		return
	}

	tokenRequest := &oauth.TokenRequest{
		ClientID:     registrationForm.ClientID,
		ClientSecret: registrationForm.ClientSecret,
		UserID:       fmt.Sprint(userID),
	}

	tokenInfo, err := oauth.GenerateToken(tokenRequest)
	if err != nil {
		_ = models.DeleteUser(userID)
		ResponseWithError(ctx, http.StatusUnauthorized, errors.New(constants.OAuthClientInvalid))
		return
	}

	response := serializers.RegistrationResponse{
		UserID:       userID,
		AccessToken:  tokenInfo.GetAccess(),
		RefreshToken: tokenInfo.GetRefresh(),
	}

	ctx.JSON(http.StatusOK, response)
}
