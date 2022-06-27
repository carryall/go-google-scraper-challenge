package helpers

import (
	"net/http"

	"go-google-scraper-challenge/constants"

	"github.com/gin-gonic/gin"
)

var ActionsWithGetMethod = []string{"index", "new", "show", "edit", "delete", "cache"}

func HandleUnauthorizeRequest(ctx *gin.Context, actionName string) {
	isGetMethod := false
	for _, action := range ActionsWithGetMethod {
		if action == actionName {
			isGetMethod = true
		}
	}

	if isGetMethod {
		signInPath := constants.WebRoutes["session"]["new"]
		ctx.Redirect(http.StatusFound, signInPath)
	} else {
		ctx.AbortWithStatus(http.StatusMethodNotAllowed)
	}
}

func RedirectToDashboard(ctx *gin.Context) {
	dashboardPath := constants.WebRoutes["results"]["index"]
	ctx.Redirect(http.StatusFound, dashboardPath)
}
