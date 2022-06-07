package controllers

import (
	"strconv"

	"go-google-scraper-challenge/errors"
	. "go-google-scraper-challenge/helpers/api"
	"go-google-scraper-challenge/lib/models"
	"go-google-scraper-challenge/lib/services/oauth"

	"github.com/gin-gonic/gin"
)

type BaseController struct {
	CurrentUser *models.User
}

func (b *BaseController) EnsureAuthenticatedUser(ctx *gin.Context) error {
	currentUser, err := b.GetCurrentUser(ctx)
	if err != nil {
		RenderJSONError(ctx, errors.ErrUnauthorizedUser, err.Error())

		return err
	}

	b.CurrentUser = currentUser

	return nil
}

func (b *BaseController) GetCurrentUser(ctx *gin.Context) (user *models.User, err error) {
	tokenInfo, err := oauth.ValidateToken(ctx.Request)
	if err != nil {
		return nil, err
	}

	userID, err := strconv.Atoi(tokenInfo.GetUserID())
	if err != nil {
		return nil, err
	}

	user, err = models.GetUserByID(int64(userID))
	if err != nil {
		return nil, err
	}

	return user, nil
}
