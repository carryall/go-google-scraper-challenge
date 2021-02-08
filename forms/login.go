package forms

import (
	"go-google-scraper-challenge/helpers"
	"go-google-scraper-challenge/models"
	"log"

	"github.com/beego/beego/v2/core/validation"
)

type LoginForm struct {
	Email    string `form:"email" valid:"Email; Required"`
	Password string `form:"password" valid:"Required;"`
}

// Valid adds custom validation to registration form, sets error when the validation failed.
func (loginForm *LoginForm) Valid(v *validation.Validation) {
	user, err := models.GetUserByEmail(loginForm.Email)
	if err != nil {
		validationError := v.SetError("Email", "Incorrect email or password")
		if validationError == nil {
			log.Print("Failed to set error on validation")
		}
	} else {
		validPassword := helpers.CompareHashWithPassword(user.HashedPassword, loginForm.Password)
		if !validPassword {
			validationError := v.SetError("Password", "Incorrect email or password")
			if validationError == nil {
				log.Print("Failed to set error on validation")
			}
		}
	}
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
