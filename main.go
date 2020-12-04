package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"

	"github.com/carryall/go-google-scraper-challenge/controller"
	api "github.com/carryall/go-google-scraper-challenge/controller/api"
)

func main() {
	app := iris.New()

	mvc.Configure(app.Party("/"), appHandler)
	mvc.Configure(app.Party("/api"), apiHandler)

	app.Listen(":8080", iris.WithLogLevel("debug"))
}

func apiHandler(app *mvc.Application) {
	app.Handle(new(api.MainController))
}

func appHandler(app *mvc.Application) {
	app.Handle(new(controller.MainController))
}
