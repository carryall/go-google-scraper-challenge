package controllers

import (
	"net/http"

	"go-google-scraper-challenge/forms"
	"go-google-scraper-challenge/helpers"
	"go-google-scraper-challenge/models"
	"go-google-scraper-challenge/presenters"
	"go-google-scraper-challenge/services/scraper"

	"github.com/beego/beego/v2/adapter/context"
	"github.com/beego/beego/v2/adapter/utils/pagination"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

// ResultController operations for User
type ResultController struct {
	BaseController
}

const ITEMS_PER_PAGE = 20

// URLMapping map user controller actions to functions
func (c *ResultController) URLMapping() {
	c.Mapping("List", c.List)
	c.Mapping("Create", c.Create)
}

func (c *ResultController) List() {
	c.EnsureAuthenticatedUser()
	c.TplName = "results/list.html"
	web.ReadFromRequest(&c.Controller)

	totalResultCount, err := models.CountResultsByUserId(c.CurrentUser.Id)
	if err != nil {
		logs.Warn("Failed to count user results: ", err.Error())
		c.Data["results"] = []*models.Result{}
	}

	perPage := helpers.GetPaginationPerPage()
	paginator := pagination.SetPaginator((*context.Context)(c.Ctx), perPage, totalResultCount)

	results, err := models.GetPaginatedResultsByUserId(c.CurrentUser.Id, int64(perPage), int64(paginator.Offset()))
	if err != nil {
		logs.Warn("Failed to get current user results: ", err.Error())
	}

	resultSets := presenters.PrepareResultSet(results)

	c.Data["paginator"] = paginator
	c.Data["resultSets"] = resultSets
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
		keywords, err := uploadForm.Save()
		if err != nil {
			flash.Error(err.Error())
		} else {
			flash.Success("Successfully uploaded the file, the result status would be updated soon")
			scraper.Search(keywords, c.CurrentUser)
		}
	}

	flash.Store(&c.Controller)
	c.Redirect("/", http.StatusFound)
}
