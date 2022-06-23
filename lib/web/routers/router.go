package routers

import (
	"go-google-scraper-challenge/lib/web/controllers"

	"github.com/gin-gonic/gin"
)

const ROOT_VIEW_PATH = "lib/web/views"
const PARTIAL_PATH = ROOT_VIEW_PATH + "/partials"

func ComebineRoutes(engine *gin.Engine) {
	// Assets
	router := engine.Group("/")
	router.Static("/static", "./static")
	router.Static("/assets/images", "./lib/web/assets/images")

	// Routes
	router.GET("/", controllers.HomeController{}.Index)
}
