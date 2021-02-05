package controllers

import "github.com/beego/beego/v2/server/web"

// AuthController operations for Auth
type AuthController struct {
	BaseController
}

// URLMapping ...
func (c *AuthController) URLMapping() {
	c.Mapping("New", c.New)
}

// New handle user login
// @Title New
// @Success 200
// @router / [get]
func (c *AuthController) New() {
	c.Data["Title"] = "Login"

	c.Layout = "layouts/default.tpl"
	c.TplName = "auth/login.html"

	web.ReadFromRequest(&c.Controller)
}
