package controllers

import (
	api_controllers "go-google-scraper-challenge/lib/api/v1/controllers"

	"github.com/gin-gonic/gin"
)

type BaseWebController struct {
	api_controllers.BaseController
}

func (c BaseWebController) Data(ctx *gin.Context, data gin.H) gin.H {
	data["CurrentPath"] = ctx.Request.URL

	return data
}
