package controllers

import (
	"go-google-scraper-challenge/helpers/log"
	"go-google-scraper-challenge/view"
	"net/http"

	"github.com/foolin/goview"
	"github.com/gin-gonic/gin"
)

type SessionsController struct {
	BaseWebController
}

func (c SessionsController) New(ctx *gin.Context) {
	view.SetLayout("authentication")

	err := goview.Render(ctx.Writer, http.StatusOK, "sessions/new", c.Data(ctx, gin.H{"Title": "Sign In"}))
	if err != nil {
		log.Info("Error", err.Error())
	}
}
