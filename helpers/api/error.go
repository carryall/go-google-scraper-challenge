package helpers

import (
	"go-google-scraper-challenge/constants"

	"github.com/gin-gonic/gin"
)

func ResponseWithError(ctx *gin.Context, status int, err error, code string) {
	RenderJSONError(ctx, status, err, constants.Errors[status], code)
}
