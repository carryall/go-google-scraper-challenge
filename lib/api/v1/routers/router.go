package routers

import (
	"go-google-scraper-challenge/lib/api/v1/controllers"
	oauth_controllers "go-google-scraper-challenge/lib/api/v1/controllers/oauth"
	. "go-google-scraper-challenge/lib/middlewares/api"

	"github.com/gin-gonic/gin"
)

func ComebineRoutes(engine *gin.Engine) {
	v1 := engine.Group("/api/v1")
	v1.Use(CurrentUser)

	healthController := controllers.HealthController{}
	oauthClientsController := oauth_controllers.OAuthClientsController{}
	registerController := controllers.UsersController{}
	authenticationController := controllers.AuthenticationController{}
	resultsController := controllers.ResultsController{}

	publicRoutes := v1.Group("/")
	publicRoutes.GET("/health", healthController.HealthStatus)
	publicRoutes.POST("/oauth/clients", oauthClientsController.Create)
	publicRoutes.POST("/register", registerController.Register)
	publicRoutes.POST("/login", authenticationController.Login)

	privateRoutes := v1.Group("/")
	privateRoutes.Use(EnsureAuthenticatedUser)
	privateRoutes.POST("/results", resultsController.Create)
	privateRoutes.GET("/results", resultsController.List)
	privateRoutes.GET("/results/:id", resultsController.Show)
}
