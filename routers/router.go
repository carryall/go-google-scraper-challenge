package routers

import (
	"go-google-scraper-challenge/controllers"
	apicontrollers "go-google-scraper-challenge/controllers/api"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	//
	web.Router("/", &controllers.SearchResultController{}, "get:List")
	web.Router("/search_results", &controllers.SearchResultController{}, "post:Create")

	// Authentication
	web.Router("/signup", &controllers.UserController{}, "get:New")
	web.Router("/users", &controllers.UserController{}, "post:Create")
	web.Router("/signin", &controllers.SessionController{}, "get:New")
	web.Router("/signout", &controllers.SessionController{}, "get:Delete")
	web.Router("/sessions", &controllers.SessionController{}, "post:Create")

	// OAuth Client
	web.Router("/oauth_client", &controllers.OAuthClientController{}, "get:New;post:Create")
	web.Router("/oauth_client/:id", &controllers.OAuthClientController{}, "get:Show")

	// API
	web.Router("/api/v1/login", &apicontrollers.AuthController{}, "post:Login")
}
