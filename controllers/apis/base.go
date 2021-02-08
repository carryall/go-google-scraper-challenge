package apis

import (
	"log"
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	web.Controller
}

type ErrorResponse struct {
	ErrorType        string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

// ResponseWithError response with JSON error data
func (this *BaseController) ResponseWithError(errorMessage string, status int) {
	controller := &this.Controller

	controller.Data["json"] = &ErrorResponse{
		ErrorType:        http.StatusText(status),
		ErrorDescription: errorMessage,
	}
	controller.Ctx.Output.Status = status

	err := controller.ServeJSON()
	if err != nil {
		log.Println("Failed to serve JSON response", err.Error())
	}
}
