package controllers

import (
	"net/http"

	"go-google-scraper-challenge/constants"
	"go-google-scraper-challenge/forms"

	"github.com/beego/beego/v2/server/web"
)

// SessionController operations for User
type SessionController struct {
	BaseController
}

// URLMapping map user controller actions to functions
func (c *SessionController) URLMapping() {
	c.Mapping("New", c.New)
	c.Mapping("Create", c.Create)
	c.Mapping("Delete", c.Delete)
}

// New handle new session action
// @Title New
// @Description new session
// @Success 200
// @router / [get]
func (c *SessionController) New() {
	c.EnsureGuestUser()

	c.Data["Title"] = "Sign In"

	c.Layout = "layouts/authentication.html"
	c.TplName = "sessions/new.html"

	web.ReadFromRequest(&c.Controller)
}

// Create handle create session action
// @Title Create
// @Description create session
// @Success 302 redirect to root path with success message
// @Failure 405 response with method not allowed when user already signed in
// @Failure 302 redirect to sign in path with error message
// @router / [post]
func (c *SessionController) Create() {
	c.EnsureGuestUser()

	flash := web.NewFlash()
	form := forms.SessionForm{}
	redirectPath := ""

	err := c.ParseForm(&form)
	if err != nil {
		flash.Error(err.Error())
	}

	user, err := form.Save()
	if err != nil {
		flash.Error(err.Error())
		redirectPath = "/signin"
	} else {
		c.SetCurrentUser(user)

		flash.Success(constants.SignInSuccess)
		redirectPath = "/"
	}

	flash.Store(&c.Controller)
	c.Redirect(redirectPath, http.StatusFound)
}

// Delete handle delete session action
// @Title Delete
// @Description delete session
// @Success 302 redirect to sign in path with success message
// @Failure 302 redirect to root path with error message
// @router / [get]
func (c *SessionController) Delete() {
	c.EnsureAuthenticatedUser()

	flash := web.NewFlash()
	redirectPath := ""

	err := c.ClearCurrentUser()
	if err != nil {
		flash.Error(constants.SignOutFail)
		redirectPath = "/"
	} else {
		flash.Success(constants.SignOutSuccess)
		redirectPath = "/signin"
	}

	flash.Store(&c.Controller)
	c.Redirect(redirectPath, http.StatusFound)
}
