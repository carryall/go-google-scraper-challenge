package helpers

import (
	"go-google-scraper-challenge/constants"
	"go-google-scraper-challenge/lib/api/v1/serializers"

	"github.com/gin-gonic/gin"
)

func ResponseWithError(ctx *gin.Context, code int, err error) {
	errorResponse := serializers.ErrorResponse{
		Error:            constants.Errors[code],
		ErrorDescription: err.Error(),
	}

	ctx.JSON(code, errorResponse)
}
