package controllers

import (
	"net/http"

	oauth_services "go-google-scraper-challenge/services/oauth"

	"github.com/beego/beego/v2/server/web"
)

// OAuthClientController operations for OAuth client
type OAuthClientController struct {
	BaseController
}

// URLMapping map OAuth client controller actions to functions
func (c *OAuthClientController) URLMapping() {
	c.Mapping("New", c.New)
	c.Mapping("Create", c.Create)
	c.Mapping("Show", c.Show)
}

// New handle new OAuth client action
// @Title New
// @Description new OAuth client
// @Success 200
// @router / [get]
func (c *OAuthClientController) New() {
	c.Data["Title"] = "New OAuth Client"

	c.Layout = "layouts/default.tpl"
	c.TplName = "oauth_clients/new.tpl"
}

// Create handle create OAuth client action
// @Title Create
// @Description create OAuth client
// @Param	body		body 	forms.Registration	true		"body for Registration form"
// @Success 302 redirect to signup with success message
// @Failure 302 redirect to signup with error message
// @router / [post]
func (c *OAuthClientController) Create() {
	flash := web.NewFlash()
	oauthClient, err := oauth_services.GenerateClient()
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/oauth_client", http.StatusFound)
	} else {
		flash.Success("The Client was successfully created")
		flash.Store(&c.Controller)
		c.Redirect("/oauth_client/"+oauthClient.ClientID, http.StatusFound)
	}
}

// Show handle show OAuth client action
// @Title Show
// @Description show OAuth client
// @Success 200
// @router / [get]
func (c *OAuthClientController) Show() {
	c.Layout = "layouts/default.tpl"
	c.TplName = "oauth_clients/show.tpl"
	c.Data["Title"] = "OAuth Client Detail"

	clientID := c.Ctx.Input.Param(":id")

	oauthClient, err := oauth_services.GetClientStore().GetByID(clientID)
	if err != nil {
		flash := web.NewFlash()
		flash.Error("OAuth client not found")
		flash.Store(&c.Controller)
		c.Redirect("/oauth_client", http.StatusFound)
	} else {
		web.ReadFromRequest(&c.Controller)
		c.Data["ClientID"] = oauthClient.GetID()
		c.Data["ClientSecret"] = oauthClient.GetSecret()
	}
}
