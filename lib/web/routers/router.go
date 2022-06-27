package routers

import (
	"go-google-scraper-challenge/constants"
	webcontrollers "go-google-scraper-challenge/lib/web/controllers"

	"github.com/gin-gonic/gin"
)

func ComebineRoutes(engine *gin.Engine) {
	// Assets
	router := engine.Group("/")
	router.Static("/static", "./static")
	router.Static("/assets/images", "./lib/web/assets/images")

	homeController := webcontrollers.HomeController{}
	sessionsController := webcontrollers.SessionsController{}
	usersController := webcontrollers.UsersController{}

	// Routes
	router.GET(constants.WebRoutes["home"]["index"], homeController.Index)
	router.GET(constants.WebRoutes["sessions"]["new"], sessionsController.New)
	router.GET(constants.WebRoutes["session"]["create"], sessionsController.Create)
	router.GET(constants.WebRoutes["users"]["new"], usersController.New)
	router.GET(constants.WebRoutes["users"]["create"], usersController.Create)
}
