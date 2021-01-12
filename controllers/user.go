package controllers

import (
	"strings"

	"go-google-scraper-challenge/models"
	"go-google-scraper-challenge/models/forms"

	"github.com/beego/beego/v2/core/validation"
	beego "github.com/beego/beego/v2/server/web"
)

//  UserController operations for User
type UserController struct {
	beego.Controller
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
	c.Data["Form"] = &forms.Registration{}
	c.Data["Alert"] = ""
	c.TplName = "users/new.tpl"

	flash := beego.ReadFromRequest(&c.Controller)
	if n, ok := flash.Data["notice"]; ok {
		// Display settings successful
		c.Data["Alert"] = n
	} else if n, ok = flash.Data["error"]; ok {
		// Display error messages
		c.Data["Alert"] = n
	}
}

// Post ...
// @Title Post
// @Description create User
// @Param	body		body 	forms.Registration	true		"body for Registration form"
// @Success 302 redirect to signup with success message
// @Failure 302 redirect to signup with error message
// @router / [post]
func (c *UserController) Post() {
	valid := validation.Validation{}
	var form forms.Registration

	err := c.ParseForm(&form)
	if err != nil {
		c.SetFlash("error", err.Error())
		c.Redirect("/signup", 302)
	}

	validForm, err := valid.Valid(&form)
	if err != nil {
		c.SetFlash("error", err.Error())
		c.Redirect("/signup", 302)
	}

	if !validForm {
		errors := []string{}
		for _, err := range valid.Errors {
			errors = append(errors, err.Message)
		}

		c.SetFlash("error", strings.Join(errors, ","))
		c.Redirect("/signup", 302)
	} else {
		user := models.User{
			Email: form.Email,
		}

		if _, err := models.AddUser(&user); err == nil {
			c.SetFlash("success", "New User created with email "+user.Email)
			c.Redirect("/signup", 302)
		} else {
			c.SetFlash("error", err.Error())
			c.Redirect("/signup", 302)
		}
	}
}

func (c *UserController) SetFlash(flashType string, message string) {
	flash := beego.NewFlash()

	switch flashType {
	case "success":
		flash.Success(message)
	case "notice":
		flash.Notice(message)
	case "error":
		flash.Error(message)
	}
	flash.Store(&c.Controller)
}
