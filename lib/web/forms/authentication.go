package webforms

import (
	"errors"

	"go-google-scraper-challenge/constants"
	"go-google-scraper-challenge/helpers"
	"go-google-scraper-challenge/lib/models"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type AuthenticationForm struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

func (f AuthenticationForm) Validate() (valid bool, err error) {
	err = validation.ValidateStruct(&f,
		validation.Field(&f.Email, validation.Required, is.Email),
		validation.Field(&f.Password, validation.Required),
	)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (f AuthenticationForm) Save() (*models.User, error) {
	user, err := models.GetUserByEmail(f.Email)
	if err != nil {
		return nil, errors.New(constants.SignInFail)
	}

	validPassword := helpers.CompareHashWithPassword(user.HashedPassword, f.Password)
	if !validPassword {
		return nil, errors.New(constants.SignInFail)
	}

	return user, nil
}
