package forms

import (
	"go-google-scraper-challenge/constants"
	"go-google-scraper-challenge/helpers"
	"go-google-scraper-challenge/models"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
)

type SessionForm struct {
	Email    string `form:"email" valid:"Email; Required"`
	Password string `form:"password" valid:"Required;"`
}

var currentUser *models.User

// Valid adds custom validation to registration form, sets error when the validation failed.
func (sf *SessionForm) Valid(v *validation.Validation) {
	user, err := models.GetUserByEmail(sf.Email)
	if err != nil {
		err := v.SetError("Email", constants.SignInFail)
		if err == nil {
			logs.Info("Failed to set error on validation")
		}
	} else {
		validPassword := helpers.CompareHashWithPassword(user.HashedPassword, sf.Password)
		if !validPassword {
			err := v.SetError("Password", constants.SignInFail)
			if err == nil {
				logs.Info("Failed to set error on validation")
			}
		} else {
			currentUser = user
		}
	}
}

// Save validates login form, returns errors if validation failed.
func (sf *SessionForm) Save() (*models.User, []error) {
	validation := validation.Validation{}

	valid, err := validation.Valid(sf)
	if err != nil {
		return nil, []error{err}
	}

	if !valid {
		var errs []error
		for _, err := range validation.Errors {
			errs = append(errs, err)
		}

		return nil, errs
	}

	return currentUser, nil
}
