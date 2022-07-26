package webcontrollers

import (
	"html/template"
	"net/http"
	"strconv"

	"go-google-scraper-challenge/constants"
	"go-google-scraper-challenge/lib/api/v1/forms"
	"go-google-scraper-challenge/lib/models"
	"go-google-scraper-challenge/lib/sessions"
	"go-google-scraper-challenge/view"

	"github.com/foolin/goview"
	"github.com/gin-gonic/gin"
)

type ResultsController struct {
	BaseController
}

func (c *ResultsController) Index(ctx *gin.Context) {
	view.SetLayout("default")

	currentUser := c.GetCurrentUser(ctx)
	results, err := models.GetUserResults(currentUser.ID, []string{}, "")
	if err != nil {
		c.RenderError(ctx, err.Error())
		ctx.Abort()
	}

	data := c.Data(ctx, gin.H{
		"Flashes": sessions.GetFlash(ctx),
		"Results": results,
	})

	err = goview.Render(ctx.Writer, http.StatusOK, "results/index", data)
	if err != nil {
		c.RenderError(ctx, err.Error())
		ctx.Abort()
	}
}

func (c *ResultsController) Create(ctx *gin.Context) {
	currentUser := c.GetCurrentUser(ctx)

	file, fileHeader, err := ctx.Request.FormFile("file")
	if err != nil {
		c.RenderError(ctx, err.Error())
		ctx.Abort()

		return
	}

	uploadForm := &forms.UploadForm{
		File:       file,
		FileHeader: fileHeader,
		User:       currentUser,
	}

	_, err = uploadForm.Save()
	if err != nil {
		c.RenderError(ctx, err.Error())
		ctx.Abort()

		return
	}

	ctx.Redirect(http.StatusFound, constants.WebRoutes["results"]["index"])
}

func (c *ResultsController) Show(ctx *gin.Context) {
	view.SetLayout("default")

	resultIDStr := ctx.Param("id")
	resultID, err := strconv.ParseInt(resultIDStr, 10, 0)
	if err != nil {
		c.renderNotFoundError(ctx)

		return
	}

	currentUser := c.GetCurrentUser(ctx)
	result, err := models.GetResultByID(resultID, currentUser, []string{"User", "AdLinks", "Links"})
	if err != nil {
		c.renderNotFoundError(ctx)

		return
	}

	data := c.Data(ctx, gin.H{
		"Result":         result,
		"TotalLinkCount": len(result.AdLinks) + len(result.Links),
		"GroupedAdLinks": models.GroupAdLinksByPosition(result.AdLinks),
	})

	err = goview.Render(ctx.Writer, http.StatusOK, "results/show", data)
	if err != nil {
		c.RenderError(ctx, err.Error())
		ctx.Abort()
	}
}

func (c *ResultsController) Cache(ctx *gin.Context) {
	view.SetLayout("default")

	resultIDStr := ctx.Param("id")
	resultID, err := strconv.ParseInt(resultIDStr, 10, 0)
	if err != nil {
		c.renderNotFoundError(ctx)
	}

	currentUser := c.GetCurrentUser(ctx)
	result, err := models.GetResultByID(resultID, currentUser, []string{})
	if err != nil {
		c.renderNotFoundError(ctx)
	}

	data := c.Data(ctx, gin.H{
		"PageCache": template.HTML(result.PageCache),
	})

	err = goview.Render(ctx.Writer, http.StatusOK, "results/cache", data)
	if err != nil {
		c.RenderError(ctx, err.Error())
		ctx.Abort()
	}
}

func (c *ResultsController) renderNotFoundError(ctx *gin.Context) {
	err := goview.Render(ctx.Writer, http.StatusOK, "results/not_found", c.Data(ctx, gin.H{}))
	if err != nil {
		c.RenderError(ctx, err.Error())
	}

	ctx.Abort()
}
