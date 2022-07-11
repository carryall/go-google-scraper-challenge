package webmiddlewares

import (
	"net/http"

	"go-google-scraper-challenge/constants"
	"go-google-scraper-challenge/lib/models"
	"go-google-scraper-challenge/lib/sessions"

	"github.com/gin-gonic/gin"
)

func CurrentUser(ctx *gin.Context) {
	currentUserID := sessions.GetCurrentUserID(ctx)

	if currentUserID == nil {
		ctx.Set("CurrentUser", nil)
	} else {
		user, err := models.GetUserByID(*currentUserID)
		if err != nil {
			ctx.Set("CurrentUser", nil)
		}

		ctx.Set("CurrentUser", user)
	}

	ctx.Next()
}

func EnsureGuestUser(ctx *gin.Context) {
	currentUser := ctx.MustGet("CurrentUser")

	if currentUser != nil {
		dashboardPath := constants.WebRoutes["results"]["index"]
		ctx.Redirect(http.StatusFound, dashboardPath)
		ctx.Abort()
	}
}

func EnsureAuthenticatedUser(ctx *gin.Context) {
	currentUser := ctx.MustGet("CurrentUser")

	requestMethod := ctx.Request.Method

	if currentUser == nil {
		if requestMethod == "GET" {
			signInPath := constants.WebRoutes["sessions"]["new"]

			ctx.Redirect(http.StatusFound, signInPath)
			ctx.Abort()
		} else {
			ctx.AbortWithStatus(http.StatusMethodNotAllowed)
		}
	}
}
