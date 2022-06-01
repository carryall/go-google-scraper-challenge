package controllers

import (
	"github.com/gin-gonic/gin"
)

type ResultsController struct {
	BaseController
}

func (c *ResultsController) Create(ctx *gin.Context) {
	c.EnsureAuthenticatedUser(ctx)
}
