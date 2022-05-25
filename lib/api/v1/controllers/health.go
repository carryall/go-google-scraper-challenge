package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct {
	BaseController
}

func (c *HealthController) HealthStatus(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "alive",
	})
}
