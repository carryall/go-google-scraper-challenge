package controllers

import (
	"fmt"
	"net/http"

	"go-google-scraper-challenge/models"
	oauth_services "go-google-scraper-challenge/services/oauth"

	"github.com/beego/beego/v2/server/web"
)

// OAuthClientController operations for OAuth client
type OAuthClientController struct {
	BaseController
}

type OAuthResponse struct {
	ClientID     *string `json:"client_id"`
	ClientSecret *string `json:"client_secret"`
}

// URLMapping map OAuth client controller actions to functions
func (c *OAuthClientController) URLMapping() {
	c.Mapping("New", c.New)
	c.Mapping("Post", c.Create)
}

// New handle new OAuth client action
// @Title New
// @Description new OAuth client
// @Success 200
// @router / [get]
func (c *OAuthClientController) New() {
	c.Data["Title"] = "New Client"

	c.Layout = "layouts/default.tpl"
	c.TplName = "oauth_clients/new.tpl"

	web.ReadFromRequest(&c.Controller)

	fmt.Println(c.Data)
}

// Post handle create OAuth client action
// @Title Post
// @Description create OAuth client
// @Param	body		body 	forms.Registration	true		"body for Registration form"
// @Success 302 redirect to signup with success message
// @Failure 302 redirect to signup with error message
// @router / [post]
func (c *OAuthClientController) Create() {
	flash := web.NewFlash()
	oauthClient, err := oauth_services.GenerateClient()
	if err != nil {
		fmt.Println("Failed to generate new client", err.Error())
		flash.Error(err.Error())

	} else {
		flash.Success("New Client was successfully created")
	}

	flash.Store(&c.Controller)

	c.Redirect("/oauth_client/"+oauthClient.ClientID, http.StatusFound)
}

// Show handle show OAuth client action
// @Title Show
// @Description show OAuth client
// @Success 200
// @router / [get]
func (c *OAuthClientController) Show() {
	flash := web.NewFlash()
	c.Layout = "layouts/default.tpl"
	c.TplName = "oauth_clients/show.tpl"
	c.Data["Title"] = "OAuth Client Detail"

	clientID := c.Ctx.Input.Param(":id")

	oauthClient, err := models.FindClientByID(clientID)
	if err != nil {
		flash.Error("OAuth client not found")
		flash.Store(&c.Controller)
		web.ReadFromRequest(&c.Controller)
	}

	c.Data["Client"] = oauthClient

	fmt.Println("Client ID", clientID)
	fmt.Println(c.Data)
}
