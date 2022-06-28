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

	homeController := controllers.HomeController{}
	sessionsController := controllers.SessionsController{}
	usersController := controllers.UsersController{}

	// Routes
	router.GET(constants.WebRoutes["home"]["index"], homeController.Index)
	router.GET(constants.WebRoutes["sessions"]["new"], sessionsController.New)
	router.GET(constants.WebRoutes["users"]["new"], usersController.New)
}
