package helpers

import (
	"go-google-scraper-challenge/errors"
	"go-google-scraper-challenge/helpers/log"
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
	title := err.Error()
	detail := oauth_errors.Descriptions[err]

	log.Infoln("RenderOAuthJSONError: ", err)
	log.Infoln("RenderOAuthJSONError: ", title, " Detail: ", detail)

	payload := payloadFromError(errors.ErrInvalidCredentials, title, detail)

	ctx.AbortWithStatusJSON(http.StatusUnauthorized, payload)
}

func RenderJSONError(ctx *gin.Context, err error, detail string) {
	payload := payloadFromError(err, "", detail)

	ctx.AbortWithStatusJSON(errors.StatusCodes[err], payload)
}

func payloadFromError(err error, title string, detail string) (errorPayload *jsonapi.ErrorsPayload) {
	if len(title) == 0 {
		title = errors.Titles[err]
	}

	if len(detail) == 0 {
		detail = errors.Descriptions[err]
	}

	errorObjs := []*jsonapi.ErrorObject{{
		Title:  errors.Titles[err],
		Detail: detail,
		Code:   err.Error(),
	}}

	errorPayload = &jsonapi.ErrorsPayload{
		Errors: errorObjs,
	}

	return
}
