package apimiddlewares

import (
	"strconv"

	"go-google-scraper-challenge/constants"
	"go-google-scraper-challenge/errors"
	helpers "go-google-scraper-challenge/helpers/api"
	"go-google-scraper-challenge/lib/models"
	"go-google-scraper-challenge/lib/services/oauth"

	"github.com/gin-gonic/gin"
)

func CurrentUser(ctx *gin.Context) {
	tokenInfo, err := oauth.ValidateToken(ctx.Request)
	if err != nil {
		ctx.Set(constants.ContextCurrentUser, nil)
		ctx.Next()

		return
	}

	userID, err := strconv.Atoi(tokenInfo.GetUserID())
	if err != nil {
		ctx.Set(constants.ContextCurrentUser, nil)
		ctx.Next()

		return
	}

	user, err := models.GetUserByID(int64(userID))
	if err != nil {
		ctx.Set(constants.ContextCurrentUser, nil)
		ctx.Next()

		return
	}

	ctx.Set(constants.ContextCurrentUser, user)
	ctx.Next()
}

func EnsureAuthenticatedUser(ctx *gin.Context) {
	currentUser := ctx.MustGet(constants.ContextCurrentUser)

	if currentUser == nil {
		helpers.RenderJSONError(ctx, errors.ErrUnauthorizedUser, "")
		ctx.Abort()
	}
}
