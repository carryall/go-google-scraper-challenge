package routers

import (
	"go-google-scraper-challenge/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/signup", &controllers.UserController{}, "get:New")
	web.Router("/users", &controllers.UserController{}, "post:Post")
}
