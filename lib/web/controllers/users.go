package controllers

import (
	"net/http"

	"go-google-scraper-challenge/helpers/log"
	"go-google-scraper-challenge/view"

	"github.com/foolin/goview"
	"github.com/gin-gonic/gin"
)

type UsersController struct {
	BaseWebController
}

func (c *UsersController) New(ctx *gin.Context) {
	view.SetLayout("authentication")

	err := goview.Render(ctx.Writer, http.StatusOK, "users/new", c.Data(ctx, gin.H{"Title": "Sign Up"}))
	if err != nil {
		log.Info("Error", err.Error())
	}
}
