package forms

import (
	"errors"

	"go-google-scraper-challenge/constants"
	"go-google-scraper-challenge/lib/models"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type AuthenticationForm struct {
	ClientID     string `form:"client_id"`
	ClientSecret string `form:"client_secret"`
	Email        string `form:"username"`
	Password     string `form:"password"`
	GrantType    string `form:"grant_type"`
}

func (f AuthenticationForm) Validate() (valid bool, err error) {
	err = validation.ValidateStruct(&f,
		validation.Field(&f.ClientID, validation.Required),
		validation.Field(&f.ClientSecret, validation.Required),
		validation.Field(&f.Email, validation.Required, is.Email),
		validation.Field(&f.Password, validation.Required),
		validation.Field(&f.GrantType, validation.Required),
	)

	if err != nil {
		return false, err
	}

	userExisted := models.UserEmailAlreadyExisted(f.Email)
	if !userExisted {
		return false, errors.New(constants.UserDoesNotExist)
	}

	return true, nil
}

// Save validates authentication form,
// returns errors if validation failed or user with with the given email does not exist.
func (f AuthenticationForm) Save() error {
	_, err := f.Validate()
	if err != nil {
		return err
	}

	_, err = models.GetUserByEmail(f.Email)
	if err != nil {
		return err
	}

	return nil
}
