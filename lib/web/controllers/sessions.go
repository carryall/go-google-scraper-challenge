package webcontrollers

import (
	"net/http"

	"go-google-scraper-challenge/constants"
	"go-google-scraper-challenge/helpers"
	"go-google-scraper-challenge/helpers/log"
	webforms "go-google-scraper-challenge/lib/web/forms"
	"go-google-scraper-challenge/view"

	"github.com/foolin/goview"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type SessionsController struct {
	BaseController
}

func (c *SessionsController) New(ctx *gin.Context) {
	c.EnsureGuestUser(ctx)
	view.SetLayout("authentication")

	err := goview.Render(ctx.Writer, http.StatusOK, "sessions/new", c.Data(ctx, gin.H{"Title": "Sign In"}))
	if err != nil {
		log.Info("Error", err.Error())
	}
}

func (c *SessionsController) Create(ctx *gin.Context) {
	c.EnsureGuestUser(ctx)

	authenticationForm := &webforms.AuthenticationForm{}

	err := ctx.ShouldBindWith(authenticationForm, binding.Form)
	if err != nil {
		// TODO: Display error
		return
	}

	_, err = authenticationForm.Validate()
	if err != nil {
		// TODO: Display error
		return
	}

	user, err := authenticationForm.Save()
	if err != nil {
		// TODO: Display error
		return
	}

	helpers.SetCurrentUser(ctx, user.ID)

	ctx.Redirect(http.StatusFound, constants.WebRoutes["result"]["index"])
}
