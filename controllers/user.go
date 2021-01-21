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
	c.Mapping("Post", c.Post)
}

// New handle new user action
// @Title New
// @Description new User
// @Success 200
// @router / [get]
func (c *UserController) New() {
	c.Data["Form"] = &forms.RegistrationForm{}
	c.Data["Alert"] = ""
	c.Data["Title"] = "Signup"

	c.Layout = "layouts/default.tpl"
	c.TplName = "users/new.tpl"

	web.ReadFromRequest(&c.Controller)
}

// Post handle create user action
// @Title Post
// @Description create User
// @Param	body		body 	forms.Registration	true		"body for Registration form"
// @Success 302 redirect to signup with success message
// @Failure 302 redirect to signup with error message
// @router / [post]
func (c *UserController) Post() {
	flash := web.NewFlash()
	form := forms.RegistrationForm{}

	err := c.ParseForm(&form)
	if err != nil {
		flash.Error(err.Error())
	}

	_, errors := form.Save()
	if len(errors) > 0 {
		for _, err := range errors {
			flash.Error(err.Error())
		}
	} else {
		flash.Success("The user was successfully created")
	}

	flash.Store(&c.Controller)
	c.Redirect("/signup", http.StatusFound)
}
