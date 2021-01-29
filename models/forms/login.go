package forms

import (
	"github.com/beego/beego/v2/core/validation"
)

type LoginForm struct {
	Email        string `valid:"Email; Required"`
	Password     string `valid:"Required;"`
	ClientId     string `valid:"Required;"`
	CLientSecret string `valid:"Required;"`
}

// Valid adds custom validation to registration form, sets error when the validation failed.
func (loginForm *LoginForm) Valid(v *validation.Validation) {
	// TODO: validate client id and secret here

}

// Save validates registration form and adds a new User with email and password from the form,
// returns errors if validation failed or cannot add the user to database.
func (loginForm *LoginForm) Save() (accessToken *string, errors []error) {
	validation := validation.Validation{}

	valid, err := validation.Valid(loginForm)
	if err != nil {
		return nil, []error{err}
	}

	if !valid {
		errors := []error{}
		for _, err := range validation.Errors {
			errors = append(errors, err)
		}

		return nil, errors
	}

	return nil, nil
}
