package routers

import (
	"go-google-scraper-challenge/controllers"
	api_controllers "go-google-scraper-challenge/controllers/api"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/login", &controllers.AuthController{}, "get:New")
	web.Router("/signup", &controllers.UserController{}, "get:New")
	web.Router("/users", &controllers.UserController{}, "post:Create")
	web.Router("/oauth_client", &controllers.OAuthClientController{}, "get:New;post:Create")
	web.Router("/oauth_client/:id", &controllers.OAuthClientController{}, "get:Show")

	web.Router("/api/v1/login", &api_controllers.AuthController{}, "post:Login")
}
