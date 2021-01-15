package controllers

import (
	"go-google-scraper-challenge/models/forms"

	"github.com/beego/beego/v2/server/web"
)

//  UserController operations for User
type UserController struct {
	web.Controller
}

// URLMapping ...
func (c *UserController) URLMapping() {
	c.Mapping("New", c.New)
	c.Mapping("Post", c.Post)
}

// New ...
// @Title New
// @Description new User
// @Success 200
// @router / [get]
func (c *UserController) New() {
	c.Data["Form"] = &forms.RegistrationForm{}
	c.Data["Alert"] = ""
	c.TplName = "users/new.tpl"

	web.ReadFromRequest(&c.Controller)
}

// Post ...
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

	userID, errors := form.Save()
	if len(errors) > 0 {
		for _, err := range errors {
			flash.Error(err.Error())
		}
	} else {
		flash.Success("New User created with ID: %d", userID)
	}

	flash.Store(&c.Controller)
	c.Redirect("/signup", 302)
}
