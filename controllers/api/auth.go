package api_controllers

import (
	"fmt"
	"go-google-scraper-challenge/models/forms"
	"net/http"
)

// AuthController operations for User
type AuthController struct {
	BaseController
}

type Response struct {
	AccessToken *string `json:"access_token"`
}

type ErrorResponse struct {
	ErrorMessages []string `json:"error_messages"`
	ErrorStatus   int      `json:"error_status"`
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
		fmt.Println(err.Error())
	}

	accessToken, errors := form.Save()
	if len(errors) > 0 {
		errorMessages := []string{}
		for _, err := range errors {
			errorMessages = append(errorMessages, err.Error())
		}

		c.Data["jsonp"] = &ErrorResponse{
			ErrorMessages: errorMessages,
			ErrorStatus:   http.StatusBadRequest,
		}
	} else {
		fmt.Println("The account was successfully created")

		c.Data["jsonp"] = &Response{
			AccessToken: accessToken,
		}
	}

	err = c.ServeJSONP()
	if err != nil {
		fmt.Println("Failed to serve JSON response")
	}
}
