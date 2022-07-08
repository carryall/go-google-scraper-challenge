package controllers

import (
	"fmt"
	"math/rand"
	"net/http"

	"go-google-scraper-challenge/errors"
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
		RenderJSONError(ctx, errors.ErrInvalidRequest, err.Error())
		return
	}

	_, err = registrationForm.Validate()
	if err != nil {
		RenderJSONError(ctx, errors.ErrInvalidRequest, err.Error())
		return
	}

	userID, err := registrationForm.Save()
	if err != nil {
		RenderJSONError(ctx, errors.ErrInvalidRequest, err.Error())
		return
	}

	tokenRequest := &oauth.TokenRequest{
		ClientID:     registrationForm.ClientID,
		ClientSecret: registrationForm.ClientSecret,
		UserID:       fmt.Sprint(userID),
	}

	tokenInfo, err := oauth.GenerateToken(tokenRequest)
	if err != nil {
		_ = models.DeleteUser(*userID)
		RenderOAuthJSONError(ctx, err)
		return
	}

	response := &serializers.RegistrationResponse{
		ID:           int64(rand.Uint64()),
		UserID:       *userID,
		AccessToken:  tokenInfo.GetAccess(),
		RefreshToken: tokenInfo.GetRefresh(),
	}

	RenderJSON(ctx, http.StatusOK, response)
}
