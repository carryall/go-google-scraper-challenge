package webcontrollers

import (
	"net/http"

	"go-google-scraper-challenge/constants"
	"go-google-scraper-challenge/helpers/log"
	"go-google-scraper-challenge/lib/sessions"
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
	view.SetLayout("authentication")

	data := c.Data(ctx, gin.H{
		"Title":   "Sign In",
		"flashes": sessions.GetFlash(ctx),
	})

	err := goview.Render(ctx.Writer, http.StatusOK, "sessions/new", data)
	if err != nil {
		log.Info("Error", err.Error())
	}
}

func (c *SessionsController) Create(ctx *gin.Context) {
	authenticationForm := &webforms.AuthenticationForm{}
	redirectURL := constants.WebRoutes["sessions"]["new"]

	err := ctx.ShouldBindWith(authenticationForm, binding.Form)
	if err != nil {
		sessions.SetFlash(ctx, sessions.FlashTypeError, err.Error())

		ctx.Redirect(http.StatusFound, redirectURL)

		return
	}

	_, err = authenticationForm.Validate()
	if err != nil {
		sessions.SetFlash(ctx, sessions.FlashTypeError, err.Error())

		ctx.Redirect(http.StatusFound, redirectURL)

		return
	}

	user, err := authenticationForm.Save()
	if err != nil {
		sessions.SetFlash(ctx, sessions.FlashTypeError, err.Error())
	} else {
		sessions.SetCurrentUser(ctx, user.ID)
		sessions.SetFlash(ctx, sessions.FlashTypeSuccess, "Successfully signed in")
		redirectURL = constants.WebRoutes["result"]["index"]
	}

	ctx.Redirect(http.StatusFound, redirectURL)
}
