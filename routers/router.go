package routers

import (
	"go-google-scraper-challenge/controllers"
	apiControllers "go-google-scraper-challenge/controllers/api"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/object", &apiControllers.ObjectController{}, "*:List")
}
