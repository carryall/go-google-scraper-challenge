package controllers

import (
	"net/http"

	"go-google-scraper-challenge/constants"
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

	c.Data["resultSets"] = resultSets
}

func (c *ResultController) Create() {
	c.EnsureAuthenticatedUser()
	flash := web.NewFlash()

	file, fileHeader, err := c.GetFile("file")
	if err != nil {
		flash.Error(constants.FileUploadFail)
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
			flash.Success(constants.FileUploadSuccess)
			scraper.Search(keywords, c.CurrentUser)
		}
	}

	flash.Store(&c.Controller)
	c.Redirect("/", http.StatusFound)
}
