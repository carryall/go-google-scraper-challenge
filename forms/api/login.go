package apiforms

import (
	"github.com/beego/beego/v2/core/validation"
)

type LoginForm struct {
	Username     string `form:"username" valid:"Email; Required"`
	Password     string `form:"password" valid:"Required;"`
	ClientId     string `form:"client_id" valid:"Required;"`
	ClientSecret string `form:"client_secret" valid:"Required;"`
	GrantType    string `form:"grant_type" valid:"Required;"`
}

// Save validates login form, returns errors if validation failed.
func (loginForm *LoginForm) Save() (errs []error) {
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
