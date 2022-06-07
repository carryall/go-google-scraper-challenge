package helpers

import (
	"net/http"

	"go-google-scraper-challenge/constants"

	"github.com/gin-gonic/gin"
	"github.com/google/jsonapi"
)

func RenderJSON(ctx *gin.Context, status_code int, data interface{}) {
	payload, err := jsonapi.Marshal(data)
	if err != nil {
		RenderJSONError(ctx, http.StatusInternalServerError, err, constants.Errors[http.StatusInternalServerError], "")
	}

	ctx.JSON(status_code, payload)
}

func RenderJSONError(ctx *gin.Context, status_code int, err error, title string, code string) {
	payload := payloadFromError(err, title, code)

	ctx.AbortWithStatusJSON(status_code, payload)
}

func payloadFromError(err error, title string, code string) (errorPayload *jsonapi.ErrorsPayload) {
	if len(title) == 0 {
		title = err.Error()
	}

	errorObjs := []*jsonapi.ErrorObject{{
		Title:  title,
		Detail: err.Error(),
		Code:   code,
	}}

	errorPayload = &jsonapi.ErrorsPayload{
		Errors: errorObjs,
	}

	return
}
