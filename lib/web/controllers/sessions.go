package controllers

import (
	"github.com/gin-gonic/gin"
)

type SessionsController struct {
	BaseWebController
}

func (c SessionsController) New(ctx *gin.Context) {
	// ctx.HTML(http.StatusOK, "sessions/new", c.Data(ctx, gin.H{
	// 	"Title": "Sign In",
	// }))
}
