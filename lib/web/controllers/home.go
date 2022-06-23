package controllers

import (
	"net/http"

	"go-google-scraper-challenge/helpers/log"

	"github.com/foolin/goview"
	"github.com/gin-gonic/gin"
)

type HomeController struct {
	BaseWebController
}

func (c HomeController) Index(ctx *gin.Context) {
	// ctx.HTML(http.StatusOK, "home/index", c.Data(ctx, gin.H{
	// 	"BodyClass": "home index",
	// }))

	err := goview.Render(ctx.Writer, http.StatusOK, "home/index", goview.M{})
	if err != nil {
		log.Info("Error", err.Error())
	}
}
