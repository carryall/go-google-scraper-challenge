package controllers

import (
	"net/http"

	"go-google-scraper-challenge/services/scraper"

	"github.com/beego/beego/v2/server/web"
)

// ResultController operations for User
type ResultController struct {
	BaseController
}

// URLMapping map user controller actions to functions
func (c *ResultController) URLMapping() {
	c.Mapping("List", c.List)
	c.Mapping("Create", c.Create)
}

func (c *ResultController) List() {
	c.EnsureAuthenticatedUser(true)

	c.Layout = "layouts/default.html"
	c.TplName = "search_results/list.html"

	web.ReadFromRequest(&c.Controller)
}

func (c *ResultController) Create() {
	c.EnsureAuthenticatedUser(false)

	keywords := []string{
		"ergonomic chair",
		"cloud storage service",
		"cloud computing service",
		"crypto currency",
		"เตา balmuda",
	}
	scraper.Search(keywords)

	c.Redirect("/", http.StatusFound)
}
