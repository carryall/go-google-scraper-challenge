package webcontrollers

import (
	"net/http"

	"go-google-scraper-challenge/helpers/log"
	"go-google-scraper-challenge/lib/models"
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

	currentUser := c.GetCurrentUser(ctx)
	results, err := models.GetUserResults(currentUser.ID, []string{}, "")
	if err != nil {
		goview.Render(ctx.Writer, http.StatusOK, "shared/error", gin.H{"Message": err.Error()})
		ctx.Abort()
	}

	data := c.Data(ctx, gin.H{
		"Flashes": sessions.GetFlash(ctx),
		"Results": results,
	})

	err = goview.Render(ctx.Writer, http.StatusOK, "results/index", data)
	if err != nil {
		log.Info("Error", err.Error())
	}
}
