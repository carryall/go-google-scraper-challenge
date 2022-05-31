package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"go-google-scraper-challenge/constants"
	. "go-google-scraper-challenge/helpers/api"
	helpers "go-google-scraper-challenge/helpers/api"
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
		ResponseWithError(ctx, http.StatusBadRequest, err, constants.ERROR_CODE_MALFORM_REQUEST)
		return
	}

	_, err = registrationForm.Validate()
	if err != nil {
		ResponseWithError(ctx, http.StatusBadRequest, err, constants.ERROR_CODE_INVALID_PARAM)
		return
	}

	userID, err := registrationForm.Save()
	if err != nil {
		ResponseWithError(ctx, http.StatusUnprocessableEntity, err, constants.ERROR_CODE_INVALID_PARAM)
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
		ResponseWithError(ctx, http.StatusUnauthorized, errors.New(constants.OAuthClientInvalid), constants.ERROR_CODE_INVALID_CREDENTIALS)
		return
	}

	response := &serializers.RegistrationResponse{
		UserID:       userID,
		AccessToken:  tokenInfo.GetAccess(),
		RefreshToken: tokenInfo.GetRefresh(),
	}

	helpers.RenderJSON(ctx, http.StatusOK, response)
}
