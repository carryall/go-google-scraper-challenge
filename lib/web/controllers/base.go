package webcontrollers

import (
	"net/http"
	"regexp"
	"runtime"
	"strings"

	"go-google-scraper-challenge/helpers"
	"go-google-scraper-challenge/helpers/log"
	api_controllers "go-google-scraper-challenge/lib/api/v1/controllers"

	"github.com/foolin/goview"
	"github.com/gin-gonic/gin"
)

type BaseController struct {
	api_controllers.BaseController
}

func (c *BaseController) RenderError(ctx *gin.Context, errorMessage string) {
	err := goview.Render(ctx.Writer, http.StatusOK, "shared/error", gin.H{"Message": errorMessage})
	if err != nil {
		log.Info("Error", err.Error())
	}
}

func (c *BaseController) Data(ctx *gin.Context, data gin.H) gin.H {
	data["CurrentPath"] = ctx.Request.URL
	data["CurrentUser"] = c.GetCurrentUser(ctx)
	controllerName, actionName := getControllerAndActionName()
	data["ControllerName"] = helpers.ToSnakeCase(controllerName)
	data["ActionName"] = helpers.ToSnakeCase(actionName)

	return data
}

func getControllerAndActionName() (controllerName string, actionName string) {
	// Get the second caller program on the stack e.g. the controller that call the Data()
	programCounter, _, _, _ := runtime.Caller(2)
	// Get the caller name, it wll be in format go-google-scraper-challenge/lib/web/controllers.(*SessionsController).New
	callerName := runtime.FuncForPC(programCounter).Name()
	callerHierarchy := strings.Split(callerName, "/")
	previousCallerName := callerHierarchy[len(callerHierarchy)-1]
	callerElements := strings.Split(previousCallerName, ".")

	re := regexp.MustCompile(`\(\*(.*)Controller\)`)
	controllerNameParts := re.FindStringSubmatch(callerElements[1])
	if len(controllerNameParts) > 0 {
		controllerName = controllerNameParts[1]
	}
	actionName = callerElements[2]

	return controllerName, actionName
}
