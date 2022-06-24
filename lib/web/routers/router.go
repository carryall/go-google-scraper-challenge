package routers

import (
	"go-google-scraper-challenge/lib/web/controllers"

	"github.com/gin-gonic/gin"
)

var webRoutes = map[string]map[string]string{
	"home": {
		"index": "/",
	},
	"session": {
		"new": "/signin",
	},
}

func ComebineRoutes(engine *gin.Engine) {
	// Assets
	router := engine.Group("/")
	router.Static("/static", "./static")
	router.Static("/assets/images", "./lib/web/assets/images")

	// Routes
	router.GET(webRoutes["home"]["index"], controllers.HomeController{}.Index)
	router.GET(webRoutes["session"]["new"], controllers.SessionsController{}.New)
}
