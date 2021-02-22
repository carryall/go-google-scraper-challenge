package controllers

import (
	"github.com/beego/beego/v2/server/web"
	"go-google-scraper-challenge/helpers"
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

	//var keyword string
	//c.Ctx.Input.Bind(&keyword, "keyword")
	//if len(keyword) > 0 {
	//	helpers.Scrape(keyword)
	//}

	keywords := []string{
		"ergonomic chair",
		"cloud storage service",
		"cloud computing service",
		"crypto currency",
	}
	helpers.Scrape(keywords)
}
