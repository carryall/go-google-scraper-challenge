package routers

import (
	"go-google-scraper-challenge/lib/web/controllers"

	"github.com/gin-gonic/gin"
)

func ComebineRoutes(engine *gin.Engine) {
	// Assets
	router := engine.Group("/")
	router.Static("/static", "./static")
	router.Static("/assets/images", "./lib/web/assets/images")

	// Routes
	router.GET("/", controllers.HomeController{}.Index)
}
