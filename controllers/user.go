package controllers

import (
	"go-google-scraper-challenge/forms"
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

// UserController operations for User
type UserController struct {
	BaseController
}

// URLMapping map user controller actions to functions
func (c *UserController) URLMapping() {
	c.Mapping("New", c.New)
	c.Mapping("Post", c.Create)
}

// New handle new user action
// @Title New
// @Description new User
// @Success 200
// @router / [get]
func (c *UserController) New() {
	c.EnsureGuestUser()

	c.Data["Title"] = "Signup"

	c.Layout = "layouts/authentication.tpl"
	c.TplName = "users/new.tpl"

	web.ReadFromRequest(&c.Controller)
}

// Create handle create user action
// @Title Create
// @Description create User
// @Param	body		body 	forms.Registration	true		"body for Registration form"
// @Success 302 redirect to signup with success message
// @Failure 302 redirect to signup with error message
// @router / [post]
func (c *UserController) Create() {
	flash := web.NewFlash()
	form := forms.RegistrationForm{}
	redirectpath := ""

	err := c.ParseForm(&form)
	if err != nil {
		flash.Error(err.Error())
	}

	_, errors := form.Save()
	if len(errors) > 0 {
		for _, err := range errors {
			flash.Error(err.Error())
		}
		redirectpath = "/signup"
	} else {
		flash.Success("The user was successfully created")
		redirectpath = "/signin"
	}

	flash.Store(&c.Controller)
	c.Redirect(redirectpath, http.StatusFound)
}
