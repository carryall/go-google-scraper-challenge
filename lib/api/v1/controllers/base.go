package controllers

import (
	"go-google-scraper-challenge/constants"
	"go-google-scraper-challenge/lib/models"

	"github.com/gin-gonic/gin"
)

type BaseController struct{}

func (c *BaseController) GetCurrentUser(ctx *gin.Context) *models.User {
	currentUser := ctx.MustGet(constants.ContextCurrentUser)
	if currentUser == nil {
		return nil
	}

	return currentUser.(*models.User)
}
