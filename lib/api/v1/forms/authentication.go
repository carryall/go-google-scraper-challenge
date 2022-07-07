package forms

import (
	"errors"
	"go-google-scraper-challenge/constants"
	"go-google-scraper-challenge/lib/models"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
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
		validation.Field(&f.Email, validation.Required, is.EmailFormat),
		validation.Field(&f.Password, validation.Required),
		validation.Field(&f.GrantType, validation.Required),
	)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (f AuthenticationForm) ValidateUser() error {
	_, err := models.GetUserByEmail(f.Email)
	if err != nil {
		return errors.New(constants.UserDoesNotExist)
	}

	return nil
}
