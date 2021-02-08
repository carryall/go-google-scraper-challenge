package controllers

import (
	"go-google-scraper-challenge/helpers"

	"github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	web.Controller
}

func (this *BaseController) Prepare() {
	helpers.SetControllerAttributes(&this.Controller)
}
