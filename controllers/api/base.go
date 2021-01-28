package api_controllers

import (
	"fmt"

	"github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	web.Controller
}

func (this *BaseController) ResponseWithError(errors []error, status int) {
	errorMessages := []string{}
	for _, err := range errors {
		errorMessages = append(errorMessages, err.Error())
	}

	this.Data["jsonp"] = &ErrorResponse{
		ErrorMessages: errorMessages,
		ErrorStatus:   status,
	}

	err := this.ServeJSON()
	if err != nil {
		fmt.Println("Failed to serve JSON response")
	}
}
