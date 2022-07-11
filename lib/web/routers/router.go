package routers

import (
	"go-google-scraper-challenge/constants"
	middlewares "go-google-scraper-challenge/lib/middlewares/web"
	webcontrollers "go-google-scraper-challenge/lib/web/controllers"

	"github.com/gin-gonic/gin"
)

func ComebineRoutes(engine *gin.Engine) {
	engine.Use(middlewares.CurrentUser)

	// Assets
	router := engine.Group("/")
	router.Static("/static", "./static")
	router.Static("/assets/images", "./lib/web/assets/images")

	homeController := webcontrollers.HomeController{}
	sessionsController := webcontrollers.SessionsController{}
	usersController := webcontrollers.UsersController{}

	publicRoutes := router.Group("/")
	publicRoutes.Use(middlewares.EnsureGuestUser)
	publicRoutes.GET(constants.WebRoutes["sessions"]["new"], sessionsController.New)
	publicRoutes.POST(constants.WebRoutes["sessions"]["create"], sessionsController.Create)
	publicRoutes.GET(constants.WebRoutes["users"]["new"], usersController.New)
	publicRoutes.POST(constants.WebRoutes["users"]["create"], usersController.Create)

	privateRoutes := router.Group("/")
	privateRoutes.Use(middlewares.EnsureAuthenticatedUser)
	privateRoutes.GET(constants.WebRoutes["home"]["index"], homeController.Index)
}
