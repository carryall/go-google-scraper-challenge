package forms

import (
	"go-google-scraper-challenge/constants"
	"go-google-scraper-challenge/helpers"
	"go-google-scraper-challenge/models"

	"github.com/beego/beego/v2/core/validation"
)

type SessionForm struct {
	Email    string `form:"email" valid:"Required; Email"`
	Password string `form:"password" valid:"Required;"`
}

var currentUser *models.User

// Valid adds custom validation to registration form, sets error when the validation failed.
func (sf *SessionForm) Valid(v *validation.Validation) {
	user, err := models.GetUserByEmail(sf.Email)
	if err != nil {
		_ = v.SetError("Email", constants.SignInFail)

		return
	}

	validPassword := helpers.CompareHashWithPassword(user.HashedPassword, sf.Password)
	if !validPassword {
		_ = v.SetError("Password", constants.SignInFail)

		return
	}

	currentUser = user
}

// Save validates login form, returns errors if validation failed.
func (sf *SessionForm) Save() (*models.User, error) {
	validation := validation.Validation{}

	valid, err := validation.Valid(sf)
	if err != nil {
		return nil, err
	}

	if !valid {
		return nil, validation.Errors[0]
	}

	return currentUser, nil
}
