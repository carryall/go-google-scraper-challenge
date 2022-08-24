package routers

import (
	"go-google-scraper-challenge/constants"
	. "go-google-scraper-challenge/lib/middlewares/web"
	webcontrollers "go-google-scraper-challenge/lib/web/controllers"

	"github.com/gin-gonic/gin"
)

func ComebineRoutes(engine *gin.Engine) {
	engine.Use(CurrentUser)

	// Assets
	router := engine.Group("/")
	router.Static("/static", "./static")
	router.Static("/assets/images", "./lib/web/assets/images")
	router.Static("/files", "./files")

	resultsController := webcontrollers.ResultsController{}
	sessionsController := webcontrollers.SessionsController{}
	usersController := webcontrollers.UsersController{}

	publicRoutes := router.Group("/")
	publicRoutes.Use(EnsureGuestUser)
	publicRoutes.GET(constants.WebRoutes["sessions"]["new"], sessionsController.New)
	publicRoutes.POST(constants.WebRoutes["sessions"]["create"], sessionsController.Create)
	publicRoutes.GET(constants.WebRoutes["users"]["new"], usersController.New)
	publicRoutes.POST(constants.WebRoutes["users"]["create"], usersController.Create)

	privateRoutes := router.Group("/")
	privateRoutes.Use(EnsureAuthenticatedUser)
	privateRoutes.GET(constants.WebRoutes["results"]["index"], resultsController.Index)
	privateRoutes.GET(constants.WebRoutes["results"]["show"], resultsController.Show)
	privateRoutes.POST(constants.WebRoutes["results"]["create"], resultsController.Create)
	privateRoutes.GET(constants.WebRoutes["results"]["cache"], resultsController.Cache)
	privateRoutes.POST(constants.WebRoutes["sessions"]["delete"], sessionsController.Delete)
}
