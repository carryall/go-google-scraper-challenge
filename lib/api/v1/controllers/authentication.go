package controllers

import (
	"errors"
	"net/http"

	"go-google-scraper-challenge/constants"
	. "go-google-scraper-challenge/helpers/api"
	"go-google-scraper-challenge/lib/api/v1/forms"
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
		ResponseWithError(ctx, http.StatusBadRequest, err)
		return
	}

	_, err = authenticationForm.Validate()
	if err != nil {
		ResponseWithError(ctx, http.StatusBadRequest, err)
		return
	}

	err = authenticationForm.ValidateUser()
	if err != nil {
		ResponseWithError(ctx, http.StatusUnauthorized, err)
		return
	}

	err = oauth.HandleTokenRequest(ctx)
	if err != nil {
		ResponseWithError(ctx, http.StatusUnauthorized, errors.New(constants.OAuthClientInvalid))
		return
	}
}
