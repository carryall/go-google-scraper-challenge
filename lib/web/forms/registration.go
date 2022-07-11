package webforms

import (
	"errors"

	"go-google-scraper-challenge/constants"
	"go-google-scraper-challenge/lib/models"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type RegistrationForm struct {
	Email                string `form:"email"`
	Password             string `form:"password"`
	PasswordConfirmation string `form:"password_confirmation"`
}

func (f RegistrationForm) Validate() (valid bool, err error) {
	err = validation.ValidateStruct(&f,
		validation.Field(&f.Email, validation.Required, is.EmailFormat),
		validation.Field(&f.Password, validation.Required, validation.Length(8, 0)),
		validation.Field(&f.PasswordConfirmation, validation.Required, validation.By(f.validatePasswordConfirmation(f.Password))),
	)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (f RegistrationForm) validatePasswordConfirmation(password string) validation.RuleFunc {
	return func(value interface{}) error {
		if value.(string) != password && len(password) > 0 {
			return errors.New("does not match the password")
		}

		return nil
	}
}

func (f RegistrationForm) Save() (*int64, error) {
	existingUser, err := models.GetUserByEmail(f.Email)
	if err == nil && existingUser != nil {
		return nil, errors.New(constants.UserAlreadyExist)
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
