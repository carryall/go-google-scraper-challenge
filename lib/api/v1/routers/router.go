package routers

import (
	"go-google-scraper-challenge/lib/api/v1/controllers"

	"github.com/gin-gonic/gin"
)

func ComebineRoutes(engine *gin.Engine) {
	v1 := engine.Group("/api/v1")

	v1.GET("/health", controllers.HealthController{}.HealthStatus)
}
