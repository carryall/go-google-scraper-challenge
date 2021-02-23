package controllers

import (
	"net/http"

	"go-google-scraper-challenge/helpers"

	"github.com/beego/beego/v2/server/web"
)

// SearchResultController operations for User
type SearchResultController struct {
	BaseController
}

// URLMapping map user controller actions to functions
func (c *SearchResultController) URLMapping() {
	c.Mapping("List", c.List)
	c.Mapping("Create", c.Create)
}

func (c *SearchResultController) List() {
	c.EnsureAuthenticatedUser(true)

	c.Layout = "layouts/default.html"
	c.TplName = "search_results/list.html"

	web.ReadFromRequest(&c.Controller)
}

func (c *SearchResultController) Create() {
	c.EnsureAuthenticatedUser(true)

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
	helpers.Search(keywords)

	c.Redirect("/", http.StatusFound)
}