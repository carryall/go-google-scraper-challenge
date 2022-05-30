package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/jsonapi"
)

func RenderJSON(ctx *gin.Context, status_code int, data interface{}) {
	payload, err := jsonapi.Marshal(data)
	if err != nil {

	}

	ctx.JSON(status_code, payload)
}

func RenderJSONError(ctx *gin.Context, status_code int, err error, title string, code string) {
	payload := payloadFromError(err, title, code)

	ctx.JSON(status_code, payload)
}

func payloadFromError(err error, title string, code string) jsonapi.ErrorObject {
	if len(title) == 0 {
		title = err.Error()
	}

	payload := jsonapi.ErrorObject{
		Title:  title,
		Detail: err.Error(),
		Code:   code,
	}

	return payload
}
