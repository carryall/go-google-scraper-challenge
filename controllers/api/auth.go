package api_controllers

import (
	"net/http"
	"strings"

	"go-google-scraper-challenge/forms"
	oauth_services "go-google-scraper-challenge/services/oauth"
)

// AuthController operations for User
type AuthController struct {
	BaseController
}

// URLMapping map user controller actions to functions
func (c *AuthController) URLMapping() {
	c.Mapping("Login", c.Login)
}

// Login provide user login API
// @Title Login
// @Description User login
// @Success 200
// @router / [post]
func (c *AuthController) Login() {
	form := forms.LoginForm{}

	err := c.ParseForm(&form)
	if err != nil {
		c.ResponseWithError(err.Error(), http.StatusBadRequest)
	}

	errs := form.Save()
	if len(errs) > 0 {
		errorMessages := []string{}
		for _, err := range errs {
			errorMessages = append(errorMessages, err.Error())
		}

		c.ResponseWithError(strings.Join(errorMessages[:], ", "), http.StatusBadRequest)
	} else {

		err = oauth_services.GenerateToken(c.Ctx)
		if err != nil {
			c.ResponseWithError(err.Error(), http.StatusBadRequest)
		}
	}
}
