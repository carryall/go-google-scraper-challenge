package controllers

import (
	"math/rand"
	"net/http"

	"go-google-scraper-challenge/errors"
	. "go-google-scraper-challenge/helpers"
	. "go-google-scraper-challenge/helpers/api"
	apiforms "go-google-scraper-challenge/lib/api/v1/forms"
	"go-google-scraper-challenge/lib/api/v1/serializers"
	"go-google-scraper-challenge/lib/services/oauth"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type AuthenticationController struct {
	BaseController
}

func (c *AuthenticationController) Login(ctx *gin.Context) {
	authenticationForm := &apiforms.AuthenticationForm{}

	err := ctx.ShouldBindWith(authenticationForm, binding.Form)
	if err != nil {
		RenderJSONError(ctx, errors.ErrInvalidRequest, err.Error())
		return
	}

	_, err = authenticationForm.Validate()
	if err != nil {
		RenderJSONError(ctx, errors.ErrInvalidRequest, err.Error())
		return
	}

	err = authenticationForm.ValidateUser()
	if err != nil {
		RenderJSONError(ctx, errors.ErrInvalidCredentials, err.Error())
		return
	}

	tokenData, err := oauth.HandleTokenRequest(ctx)
	if err != nil {
		RenderOAuthJSONError(ctx, err)
		return
	}

	tokenInfo, err := GetTokenInfo(tokenData)
	if err != nil {
		RenderJSONError(ctx, errors.ErrInvalidCredentials, err.Error())
		return
	}

	response := &serializers.AuthenticationResponse{
		ID:           int64(rand.Uint64()),
		AccessToken:  tokenInfo.AccessToken,
		RefreshToken: tokenInfo.RefreshToken,
		ExpiresIn:    tokenInfo.ExpiresIn,
		TokenType:    tokenInfo.TokenType,
	}

	RenderJSON(ctx, http.StatusOK, response)
}
