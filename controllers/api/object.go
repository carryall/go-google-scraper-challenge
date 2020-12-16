// TODO: remove this file when work on real API
package api

import (
	"go-google-scraper-challenge/models"

	"github.com/astaxie/beego"
)

// Operations about object
type ObjectController struct {
	beego.Controller
}

// @Title List
// @Description get all objects
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [get]
func (o *ObjectController) List() {
	obs := models.GetAll()
	o.Data["json"] = obs
	o.ServeJSON()
}
