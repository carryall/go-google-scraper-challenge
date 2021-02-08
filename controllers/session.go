package controllers

import (
	"go-google-scraper-challenge/forms"
	"net/http"

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
}

// New handle new session action
// @Title New
// @Description new session
// @Success 200
// @router / [get]
func (c *SessionController) New() {
	c.Data["Title"] = "Sign In"

	c.Layout = "layouts/authentication.tpl"
	c.TplName = "sessions/new.tpl"

	web.ReadFromRequest(&c.Controller)
}

// Create handle create session action
// @Title Create
// @Description create session
// @Success 302 redirect to root path with success message
// @Failure 302 redirect to login path with error message
// @router / [post]
func (c *SessionController) Create() {
	flash := web.NewFlash()
	form := forms.LoginForm{}
	redirectPath := ""

	err := c.ParseForm(&form)
	if err != nil {
		flash.Error(err.Error())
		redirectPath = "/login"
	}

	errs := form.Save()
	if len(errs) > 0 {
		flash.Error(err.Error())
		redirectPath = "/login"
	} else {
		flash.Success("Successfully logged in")
		redirectPath = "/"
	}

	flash.Store(&c.Controller)
	c.Redirect(redirectPath, http.StatusFound)
}
