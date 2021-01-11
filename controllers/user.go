package controllers

import (
	"encoding/json"
	"go-google-scraper-challenge/models"
	"log"

	beego "github.com/beego/beego/v2/server/web"
)

//  UserController operations for User
type UserController struct {
	beego.Controller
}

// URLMapping ...
func (c *UserController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// Post ...
// @Title Post
// @Description create User
// @Param	body		body 	models.User	true		"body for User content"
// @Success 201 {int} models.User
// @Failure 403 body is empty
// @router / [post]
func (c *UserController) Post() {
	var v models.User

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		c.Data["json"] = err.Error()
	}

	if _, err := models.AddUser(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = v
	} else {
		c.Data["json"] = err.Error()
	}

	if err := c.ServeJSON(); err != nil {
		log.Fatal(err)
	}
}
