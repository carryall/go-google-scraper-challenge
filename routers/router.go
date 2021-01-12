package routers

import (
	"go-google-scraper-challenge/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/signup", &controllers.UserController{}, "get:New")
	beego.Router("/users", &controllers.UserController{}, "post:Post")
}
