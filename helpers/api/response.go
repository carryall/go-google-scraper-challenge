package helpers

import (
	"go-google-scraper-challenge/errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/jsonapi"
	oauth_errors "gopkg.in/oauth2.v3/errors"
)

func RenderJSON(ctx *gin.Context, statusCode int, data interface{}) {
	payload, err := jsonapi.Marshal(data)
	if err != nil {
		RenderJSONError(ctx, errors.ErrServerError, err.Error())
	}

	ctx.JSON(statusCode, payload)
}

func RenderOAuthJSONError(ctx *gin.Context, err error) {
	detail := oauth_errors.Descriptions[err]

	payload := payloadFromError(errors.ErrInvalidCredentials, "", detail, "")

	ctx.AbortWithStatusJSON(http.StatusUnauthorized, payload)
}

func RenderJSONError(ctx *gin.Context, err error, detail string) {
	payload := payloadFromError(err, "", detail, "")

	ctx.AbortWithStatusJSON(errors.StatusCodes[err], payload)
}

func payloadFromError(err error, title string, detail string, code string) (errorPayload *jsonapi.ErrorsPayload) {
	if len(title) == 0 {
		title = errors.Titles[err]
	}

	if len(detail) == 0 {
		detail = errors.Descriptions[err]
	}

	if len(code) == 0 {
		code = err.Error()
	}

	errorObjs := []*jsonapi.ErrorObject{{
		Title:  title,
		Detail: detail,
		Code:   code,
	}}

	errorPayload = &jsonapi.ErrorsPayload{
		Errors: errorObjs,
	}

	return errorPayload
}
