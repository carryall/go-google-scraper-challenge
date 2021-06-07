package controllers

import (
	"net/http"
	"strconv"

	"go-google-scraper-challenge/constants"
	"go-google-scraper-challenge/forms"
	"go-google-scraper-challenge/helpers"
	"go-google-scraper-challenge/models"
	"go-google-scraper-challenge/presenters"

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
func (rc *ResultController) URLMapping() {
	rc.Mapping("List", rc.List)
	rc.Mapping("Create", rc.Create)
	rc.Mapping("Show", rc.Show)
	rc.Mapping("Cache", rc.Cache)
}

func (rc *ResultController) List() {
	rc.EnsureAuthenticatedUser()
	rc.TplName = "results/list.html"
	web.ReadFromRequest(&rc.Controller)

	query := rc.queryFromParams()

	totalResultCount, err := models.CountResultsByUserId(rc.CurrentUser.Id)
	if err != nil {
		logs.Warn("Failed to count user results: ", err.Error())
		rc.Data["results"] = []*models.Result{}
	}

	perPage := helpers.GetPaginationPerPage()
	paginator := pagination.SetPaginator((*context.Context)(rc.Ctx), perPage, totalResultCount)

	query["limit"] = perPage
	query["offset"] = paginator.Offset()

	results, err := models.GetResultsBy(query)
	if err != nil {
		logs.Warn("Failed to get current user results: ", err.Error())
	}

	resultSets := presenters.PrepareResultSet(results)

	rc.Data["resultSets"] = resultSets
}

func (rc *ResultController) Create() {
	rc.EnsureAuthenticatedUser()
	flash := web.NewFlash()

	file, fileHeader, err := rc.GetFile("file")
	if err != nil {
		flash.Error(constants.FileUploadFail)
	} else {
		uploadForm := forms.UploadForm{
			File: file,
			FileHeader: fileHeader,
			User: rc.CurrentUser,
		}
		keywords, err := uploadForm.Save()
		if err != nil {
			flash.Error(err.Error())
		} else {
			rc.storeKeywords(keywords)

			flash.Success(constants.FileUploadSuccess)
		}
	}

	flash.Store(&rc.Controller)
	rc.Redirect("/", http.StatusFound)
}

func (rc *ResultController) Show() {
	rc.EnsureAuthenticatedUser()
	rc.TplName = "results/show.html"
	rc.Data["Title"] = "Result Detail"
	web.ReadFromRequest(&rc.Controller)

	resultID, err := rc.getResultID()
	if err == nil {
		result, err := models.GetResultByIdWithRelations(resultID)
		if err != nil {
			logs.Error("Failed to get result:", err.Error())
		}

		rc.Data["result"] = presenters.GetResultPresenter(result)
	}
}

func (rc *ResultController) Cache() {
	rc.EnsureAuthenticatedUser()
	rc.Layout = "layouts/blank.html"
	rc.TplName = "results/cache.html"
	rc.Data["Title"] = "Result Page Cache"
	web.ReadFromRequest(&rc.Controller)

	resultID, err := rc.getResultID()
	if err != nil {
		return
	}

	result, err := models.GetResultById(resultID)
	if err != nil {
		logs.Error("Failed to get result:", err.Error())
	} else {
		rc.Data["pageCache"] = result.PageCache
	}
}

func (rc *ResultController) getResultID() (int64, error) {
	resultIDParam := rc.Ctx.Input.Param(":id")
	resultID, err := strconv.ParseInt(resultIDParam, 0, 64)
	if err != nil {
		logs.Error("Failed to parse result ID params:", err.Error())

		return 0, err
	}

	return resultID, nil
}

func (rc *ResultController) queryFromParams() map[string]interface{} {
	searcheKeyword := rc.GetString("keyword")

	var query = map[string]interface{}{
		"user_id":            rc.CurrentUser.Id,
		"order":              "-created_at",
		"keyword__icontains": searcheKeyword,
	}

	return query
}

func (rc *ResultController) storeKeywords(keywords []string)  {
	for _, k := range keywords {
		result := &models.Result{
			User: rc.CurrentUser,
			Keyword: k,
		}
		_, err := models.CreateResult(result)
		if err != nil {
			logs.Error("Failed to create result:", err.Error())
		}
	}
}
