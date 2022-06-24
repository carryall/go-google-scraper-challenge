package routers

import (
	"go-google-scraper-challenge/constants"
	"go-google-scraper-challenge/lib/web/controllers"

	"github.com/gin-gonic/gin"
)

func ComebineRoutes(engine *gin.Engine) {
	// Assets
	router := engine.Group("/")
	router.Static("/static", "./static")
	router.Static("/assets/images", "./lib/web/assets/images")

	// Routes
	router.GET(constants.WebRoutes["home"]["index"], controllers.HomeController{}.Index)
	router.GET(constants.WebRoutes["session"]["new"], controllers.SessionsController{}.New)
	router.GET(constants.WebRoutes["users"]["new"], controllers.UsersController{}.New)
}
