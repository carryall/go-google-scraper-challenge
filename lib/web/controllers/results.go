package webcontrollers

import (
	"net/http"

	"go-google-scraper-challenge/helpers/log"
	"go-google-scraper-challenge/lib/sessions"
	"go-google-scraper-challenge/view"

	"github.com/foolin/goview"
	"github.com/gin-gonic/gin"
)

type ResultsController struct {
	BaseController
}

func (c *ResultsController) Index(ctx *gin.Context) {
	view.SetLayout("default")

	data := c.Data(ctx, gin.H{
		"flashes": sessions.GetFlash(ctx),
	})

	err := goview.Render(ctx.Writer, http.StatusOK, "home/index", data)
	if err != nil {
		log.Info("Error", err.Error())
	}
}
