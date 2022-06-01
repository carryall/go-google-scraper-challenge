package controllers

import (
	. "go-google-scraper-challenge/helpers/api"
	"go-google-scraper-challenge/lib/models"
	"go-google-scraper-challenge/lib/services/oauth"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BaseController struct {
	CurrentUser *models.User
}

func (b *BaseController) EnsureAuthenticatedUser(ctx *gin.Context) {
	currentUser, err := b.GetCurrentUser(ctx)
	if err != nil {
		RenderOAuthJSONError(ctx, err)
	}

	b.CurrentUser = currentUser
}

func (b *BaseController) GetCurrentUser(ctx *gin.Context) (user *models.User, err error) {
	tokenInfo, err := oauth.ValidateToken(ctx.Request)
	if err != nil {
		return
	}

	userID, err := strconv.Atoi(tokenInfo.GetUserID())
	if err != nil {
		return
	}

	user, err = models.GetUserByID(int64(userID))

	return
}
