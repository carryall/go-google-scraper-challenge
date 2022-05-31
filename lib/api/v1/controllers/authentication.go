package controllers

import (
	"errors"
	"net/http"

	"go-google-scraper-challenge/constants"
	. "go-google-scraper-challenge/helpers/api"
	helpers "go-google-scraper-challenge/helpers/api"
	"go-google-scraper-challenge/lib/api/v1/forms"
	"go-google-scraper-challenge/lib/api/v1/serializers"
	"go-google-scraper-challenge/lib/services/oauth"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type AuthenticationController struct {
	BaseController
}

func (c *AuthenticationController) Login(ctx *gin.Context) {
	authenticationForm := &forms.AuthenticationForm{}

	err := ctx.ShouldBindWith(authenticationForm, binding.Form)
	if err != nil {
		ResponseWithError(ctx, http.StatusBadRequest, err, constants.ERROR_CODE_MALFORM_REQUEST)
		return
	}

	_, err = authenticationForm.Validate()
	if err != nil {
		ResponseWithError(ctx, http.StatusBadRequest, err, constants.ERROR_CODE_INVALID_PARAM)
		return
	}

	err = authenticationForm.ValidateUser()
	if err != nil {
		ResponseWithError(ctx, http.StatusUnauthorized, err, constants.ERROR_CODE_INVALID_CREDENTIALS)
		return
	}

	tokenData, err := oauth.HandleTokenRequest(ctx)
	if err != nil {
		ResponseWithError(ctx, http.StatusUnauthorized, errors.New(constants.OAuthClientInvalid), constants.ERROR_CODE_INVALID_CREDENTIALS)
		return
	}

	response := &serializers.AuthenticationResponse{
		AccessToken:  tokenData["access_token"].(string),
		RefreshToken: tokenData["refresh_token"].(string),
		ExpiresIn:    tokenData["expires_in"].(int64),
		TokenType:    tokenData["token_type"].(string),
	}

	helpers.RenderJSON(ctx, http.StatusOK, response)
}
