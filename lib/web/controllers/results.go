package webcontrollers

import (
	"net/http"

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
		c.RenderError(ctx, err.Error())
		ctx.Abort()
	}

	data := c.Data(ctx, gin.H{
		"Flashes": sessions.GetFlash(ctx),
		"Results": results,
	})

	err = goview.Render(ctx.Writer, http.StatusOK, "results/index", data)
	if err != nil {
		c.RenderError(ctx, err.Error())
		ctx.Abort()
	}
}
