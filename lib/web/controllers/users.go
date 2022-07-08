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

type UsersController struct {
	BaseController
}

func (c *UsersController) New(ctx *gin.Context) {
	c.EnsureGuestUser(ctx)
	view.SetLayout("authentication")

	data := c.Data(ctx, gin.H{
		"Title":   "Sign Up",
		"flashes": sessions.GetFlash(ctx),
	})

	err := goview.Render(ctx.Writer, http.StatusOK, "users/new", data)
	if err != nil {
		log.Info("Error", err.Error())
	}
}

func (c UsersController) Create(ctx *gin.Context) {
	c.EnsureGuestUser(ctx)

	registrationForm := &webforms.RegistrationForm{}
	redirectURL := constants.WebRoutes["users"]["new"]

	err := ctx.ShouldBindWith(registrationForm, binding.Form)
	if err != nil {
		sessions.SetFlash(ctx, sessions.FlashTypeError, err.Error())

		ctx.Redirect(http.StatusFound, redirectURL)

		return
	}

	_, err = registrationForm.Validate()
	if err != nil {
		sessions.SetFlash(ctx, sessions.FlashTypeError, err.Error())

		ctx.Redirect(http.StatusFound, redirectURL)

		return
	}

	userID, err := registrationForm.Save()
	if err != nil {
		sessions.SetFlash(ctx, sessions.FlashTypeError, err.Error())
	} else {
		sessions.SetCurrentUser(ctx, *userID)
		sessions.SetFlash(ctx, sessions.FlashTypeSuccess, "Successfully signed up")
		redirectURL = constants.WebRoutes["result"]["index"]
	}

	ctx.Redirect(http.StatusFound, redirectURL)
}
