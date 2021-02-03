package forms

import (
	"github.com/beego/beego/v2/core/validation"
	"github.com/beego/beego/v2/server/web/context"
)

type LoginForm struct {
	Username     string `form:"username" valid:"Email; Required"`
	Password     string `form:"password" valid:"Required;"`
	ClientId     string `form:"client_id" valid:"Required;"`
	CLientSecret string `form:"client_secret" valid:"Required;"`
	GrantType    string `form:"grant_type" valid:"Required;"`
}

// Save validates registration form and adds a new User with email and password from the form,
// returns errors if validation failed or cannot add the user to database.
func (loginForm *LoginForm) Save(c *context.Context) (errs []error) {
	validation := validation.Validation{}

	valid, err := validation.Valid(loginForm)
	if err != nil {
		return []error{err}
	}

	if !valid {
		errs := []error{}
		for _, err := range validation.Errors {
			errs = append(errs, err)
		}

		return errs
	}
	return nil
}
