package webcontrollers

import (
	"net/http"

	"go-google-scraper-challenge/helpers/log"
	"go-google-scraper-challenge/view"

	"github.com/foolin/goview"
	"github.com/gin-gonic/gin"
)

type UsersController struct {
	BaseController
}

func (c *UsersController) New(ctx *gin.Context) {
	c.EnsureGuestUser(ctx)
	view.SetLayout("authentication")

	err := goview.Render(ctx.Writer, http.StatusOK, "users/new", c.Data(ctx, gin.H{"Title": "Sign Up"}))
	if err != nil {
		log.Info("Error", err.Error())
	}
}

func (c UsersController) Create(ctx *gin.Context) {
	c.EnsureGuestUser(ctx)
}
