package controllers

import (
	"net/http"
	"strconv"

	"go-google-scraper-challenge/errors"
	. "go-google-scraper-challenge/helpers/api"
	"go-google-scraper-challenge/lib/api/v1/forms"
	"go-google-scraper-challenge/lib/api/v1/serializers"
	"go-google-scraper-challenge/lib/models"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type ResultsController struct {
	BaseController
}

func (c *ResultsController) List(ctx *gin.Context) {
	if c.EnsureAuthenticatedUser(ctx) != nil {
		return
	}

	resultSearchForm := &forms.ResultSearchForm{}
	_ = ctx.ShouldBindWith(resultSearchForm, binding.JSON)

	results, err := models.GetUserResults(c.CurrentUser.ID, []string{"User", "AdLinks", "Links"}, resultSearchForm.Keyword)
	if err != nil {
		RenderJSONError(ctx, errors.ErrServerError, err.Error())

		return
	}

	response := []*serializers.ResultResponse{}
	for _, result := range results {
		response = append(response, serializers.ResultSerializer{Result: result}.Response())
	}

	RenderJSON(ctx, http.StatusOK, response)
}

func (c *ResultsController) Show(ctx *gin.Context) {
	if c.EnsureAuthenticatedUser(ctx) != nil {
		return
	}

	resultID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		RenderJSONError(ctx, errors.ErrInvalidRequest, err.Error())

		return
	}

	result, err := models.GetResultByID(int64(resultID), c.CurrentUser, []string{"User", "AdLinks", "Links"})
	if err != nil {
		RenderJSONError(ctx, errors.ErrNotFound, err.Error())

		return
	}

	response := serializers.ResultSerializer{Result: result}
	RenderJSON(ctx, http.StatusOK, response.DetailResponse())
}

func (c *ResultsController) Create(ctx *gin.Context) {
	if c.EnsureAuthenticatedUser(ctx) != nil {
		return
	}

	file, fileHeader, err := ctx.Request.FormFile("file")
	if err != nil {
		RenderJSONError(ctx, errors.ErrInvalidRequest, err.Error())

		return
	}

	uploadForm := &forms.UploadForm{
		File:       file,
		FileHeader: fileHeader,
		User:       c.CurrentUser,
	}

	resultIDs, err := uploadForm.Save()
	if err != nil {
		RenderJSONError(ctx, errors.ErrInvalidRequest, err.Error())

		return
	}

	results, err := models.GetResultsByIDs(resultIDs)
	if err != nil {
		RenderJSONError(ctx, errors.ErrInvalidRequest, err.Error())

		return
	}

	response := []*serializers.ResultResponse{}
	for _, result := range *results {
		response = append(response, serializers.ResultSerializer{Result: &result}.Response())
	}

	RenderJSON(ctx, http.StatusCreated, response)
}
