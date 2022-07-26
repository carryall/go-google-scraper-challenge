package apiforms

import (
	"errors"

	"go-google-scraper-challenge/constants"
	"go-google-scraper-challenge/lib/models"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type RegistrationForm struct {
	ClientID     string `form:"client_id"`
	ClientSecret string `form:"client_secret"`
	Email        string `form:"username"`
	Password     string `form:"password"`
}

func (f RegistrationForm) Validate() (valid bool, err error) {
	err = validation.ValidateStruct(&f,
		validation.Field(&f.ClientID, validation.Required),
		validation.Field(&f.ClientSecret, validation.Required),
		validation.Field(&f.Email, validation.Required, is.EmailFormat),
		validation.Field(&f.Password, validation.Required, validation.Length(6, 50)),
	)

	if err != nil {
		return false, err
	}

	emailAlreadyExisted := models.UserEmailAlreadyExisted(f.Email)
	if emailAlreadyExisted {
		return false, errors.New(constants.UserAlreadyExist)
	}

	return true, nil
}

// Save validates registration form and adds a new User with email and password from the form,
// returns errors if validation failed or cannot add the user to database.
func (f RegistrationForm) Save() (*int64, error) {
	_, err := f.Validate()
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email:    f.Email,
		Password: f.Password,
	}

	userID, err := models.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return &userID, nil
}
