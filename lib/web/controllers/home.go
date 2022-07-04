package controllers

import (
	"net/http"

	"go-google-scraper-challenge/helpers/log"
	"go-google-scraper-challenge/view"

	"github.com/foolin/goview"
	"github.com/gin-gonic/gin"
)

type HomeController struct {
	BaseWebController
}

func (c *HomeController) Index(ctx *gin.Context) {
	view.SetLayout("default")

	err := goview.Render(ctx.Writer, http.StatusOK, "home/index", c.Data(ctx, gin.H{}))
	if err != nil {
		log.Info("Error", err.Error())
	}
}
