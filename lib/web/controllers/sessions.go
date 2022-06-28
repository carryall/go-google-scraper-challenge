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

	data := c.Data(ctx, gin.H{
		"Title":   "Sign In",
		"flashes": helpers.GetFlash(ctx),
	})

	err := goview.Render(ctx.Writer, http.StatusOK, "sessions/new", data)
	if err != nil {
		log.Info("Error", err.Error())
	}
}

func (c *SessionsController) Create(ctx *gin.Context) {
	c.EnsureGuestUser(ctx)

	authenticationForm := &webforms.AuthenticationForm{}
	redirectURL := constants.WebRoutes["sessions"]["new"]

	err := ctx.ShouldBindWith(authenticationForm, binding.Form)
	if err != nil {
		helpers.SetFlash(ctx, helpers.FlashTypeError, err.Error())
	}

	_, err = authenticationForm.Validate()
	if err != nil {
		helpers.SetFlash(ctx, helpers.FlashTypeError, err.Error())
	}

	user, err := authenticationForm.Save()
	if err != nil {
		helpers.SetFlash(ctx, helpers.FlashTypeError, err.Error())
	} else {
		helpers.SetCurrentUser(ctx, user.ID)
		redirectURL = constants.WebRoutes["result"]["index"]
	}

	ctx.Redirect(http.StatusFound, redirectURL)
}
