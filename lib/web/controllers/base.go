package controllers

import (
	"runtime"
	"strings"

	api_controllers "go-google-scraper-challenge/lib/api/v1/controllers"

	"github.com/gin-gonic/gin"
)

type BaseWebController struct {
	api_controllers.BaseController
}

func (c BaseWebController) Data(ctx *gin.Context, data gin.H) gin.H {
	data["CurrentPath"] = ctx.Request.URL
	controllerName, actionName := getControllerAndActionName()
	data["ControllerName"] = strings.ToLower(controllerName)
	data["ActionName"] = strings.ToLower(actionName)

	return data
}

func getControllerAndActionName() (controllerName string, actionName string) {
	// Get the second caller program on the stack e.g. the controller that call the Data()
	programCounter, _, _, _ := runtime.Caller(2)
	// Get the caller name, it wll be in format go-google-scraper-challenge/lib/web/controllers.SessionsController.New
	callerName := runtime.FuncForPC(programCounter).Name()
	callerHierarchy := strings.Split(callerName, "/")
	previousCallerName := callerHierarchy[len(callerHierarchy)-1]
	callerElements := strings.Split(previousCallerName, ".")
	controllerName = strings.Replace(callerElements[1], "Controller", "", 1)
	actionName = callerElements[2]

	return controllerName, actionName
}
