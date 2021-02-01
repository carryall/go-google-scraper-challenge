package forms

import (
	oauth_services "go-google-scraper-challenge/services/oauth"

	"github.com/beego/beego/v2/core/validation"
	"github.com/beego/beego/v2/server/web/context"
	"gopkg.in/oauth2.v3/errors"
)

type LoginForm struct {
	Email        string `valid:"Email; Required"`
	Password     string `valid:"Required;"`
	ClientId     string `valid:"Required;"`
	CLientSecret string `valid:"Required;"`
}

// Save validates registration form and adds a new User with email and password from the form,
// returns errors if validation failed or cannot add the user to database.
func (loginForm *LoginForm) Save(c *context.Context) (accessToken *string, errs []error) {
	validation := validation.Validation{}

	valid, err := validation.Valid(loginForm)
	if err != nil {
		return nil, []error{err}
	}

	if !valid {
		errs := []error{}
		for _, err := range validation.Errors {
			errs = append(errs, err)
		}

		return nil, errs
	}

	err = oauth_services.GenerateToken(c)
	if err != nil {
		return nil, []error{errors.ErrInvalidRequest}
	}

	return nil, nil
}
