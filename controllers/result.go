package controllers

import (
	"net/http"

	"go-google-scraper-challenge/forms"
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
	c.EnsureAuthenticatedUser()

	c.TplName = "results/list.html"

	web.ReadFromRequest(&c.Controller)
}

func (c *ResultController) Create() {
	c.EnsureAuthenticatedUser()
	flash := web.NewFlash()

	file, fileHeader, err := c.GetFile("file")
	if err != nil {
		flash.Error("Failed to upload file, please make sure the file is not corrupted")
	} else {
		uploadForm := forms.UploadForm{
			File: file,
			FileHeader: fileHeader,
			User: c.CurrentUser,
		}
		keywords, errs := uploadForm.Save()
		if len(errs) > 0 {
			for _, err := range errs {
				flash.Error(err.Error())
			}
		} else {
			flash.Success("Successfully upload file, the result status would be update soon")
			scraper.Search(keywords, c.CurrentUser)
		}
	}

	flash.Store(&c.Controller)
	c.Redirect("/", http.StatusFound)
}
