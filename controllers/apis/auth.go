package apicontrollers

import (
	apiforms "go-google-scraper-challenge/forms/api"
	oauth_services "go-google-scraper-challenge/services/oauth"

	"gopkg.in/oauth2.v3/errors"
)

// AuthController operations for User
type AuthController struct {
	BaseController
}

// URLMapping map user controller actions to functions
func (c *AuthController) URLMapping() {
	c.Mapping("Login", c.Login)
}

var (
	ErrInvalidRequestDescription = errors.Descriptions[errors.ErrInvalidRequest]
	ErrInvalidRequestStatus      = errors.StatusCodes[errors.ErrInvalidRequest]
)

// Login provide user login API
// @Title Login
// @Description User login
// @Success 200
// @router / [post]
func (c *AuthController) Login() {
	form := apiforms.LoginForm{}

	err := c.ParseForm(&form)
	if err != nil {
		c.ResponseWithError(ErrInvalidRequestDescription, ErrInvalidRequestStatus)
	}

	errs := form.Save()
	if len(errs) > 0 {
		c.ResponseWithError(ErrInvalidRequestDescription, ErrInvalidRequestStatus)
	} else {
		err = oauth_services.GenerateToken(c.Ctx)
		if err != nil {
			c.ResponseWithError(ErrInvalidRequestDescription, ErrInvalidRequestStatus)
		}
	}
}
