package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

// KeywordController operations for User
type KeywordController struct {
	BaseController
}

// URLMapping map user controller actions to functions
func (c *KeywordController) URLMapping() {
	c.Mapping("List", c.List)
}

func (c *KeywordController) List() {
	c.EnsureAuthenticatedUser(true)

	c.Layout = "layouts/default.html"
	c.TplName = "keywords/list.html"

	web.ReadFromRequest(&c.Controller)
}
