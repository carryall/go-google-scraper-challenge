package routers

import (
	"go-google-scraper-challenge/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
