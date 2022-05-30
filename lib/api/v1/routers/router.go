package routers

import (
	"go-google-scraper-challenge/lib/api/v1/controllers"
	oauth_controllers "go-google-scraper-challenge/lib/api/v1/controllers/oauth"

	"github.com/gin-gonic/gin"
)

func ComebineRoutes(engine *gin.Engine) {
	v1 := engine.Group("/api/v1")

	healthController := controllers.HealthController{}
	oauthClientsController := oauth_controllers.OAuthClientsController{}
	registerController := controllers.UsersController{}
	authenticationController := controllers.AuthenticationController{}

	v1.GET("/health", healthController.HealthStatus)
	v1.POST("/oauth/clients", oauthClientsController.Create)
	v1.POST("/register", registerController.Register)
	v1.POST("/login", authenticationController.Login)
}
