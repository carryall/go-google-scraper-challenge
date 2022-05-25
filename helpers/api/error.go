package api_helpers

import (
	"fmt"
	"go-google-scraper-challenge/lib/api/v1/serializers"

	"github.com/gin-gonic/gin"
)

func ResponseWithError(ctx *gin.Context, code int, err error) {
	errorResponse := serializers.ErrorResponse{
		Error:       fmt.Sprint(code),
		ErrorDetail: err.Error(),
	}

	ctx.JSON(code, errorResponse)
}
