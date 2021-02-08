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

func (c *SessionController) New() {
	c.Data["Title"] = "Sign In"

	c.Layout = "layouts/authentication.tpl"
	c.TplName = "sessions/new.tpl"

	web.ReadFromRequest(&c.Controller)
}

// Login provide user login feature
// @Title Login
// @Description User login
// @Success 302 redirect to home page with success message
// @Failure 302 redirect to login page with error message
// @router / [post]
func (c *SessionController) Create() {
	flash := web.NewFlash()
	form := forms.LoginForm{}

	err := c.ParseForm(&form)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/login", http.StatusFound)
	}

	errs := form.Save()
	if len(errs) > 0 {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/login", http.StatusFound)
	} else {
		// TODO: generate token and log user in
		flash.Success("Successfully logged in")
		flash.Store(&c.Controller)
		c.Redirect("/login", http.StatusFound)
	}
}
