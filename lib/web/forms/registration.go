package webforms

import (
	"errors"

	"go-google-scraper-challenge/constants"
	"go-google-scraper-challenge/helpers"
	"go-google-scraper-challenge/lib/models"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type RegistrationForm struct {
	Email                string `form:"email"`
	Password             string `form:"password"`
	PasswordConfirmation string `form:"passwordConfirmation"`
}

func (f RegistrationForm) Validate() (valid bool, err error) {
	err = validation.ValidateStruct(&f,
		validation.Field(&f.Email, validation.Required, is.Email),
		validation.Field(&f.Password, validation.Required),
		validation.Field(&f.PasswordConfirmation, validation.Required, validation.By(f.validatePasswordConfirmation())),
	)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (f RegistrationForm) validatePasswordConfirmation() validation.RuleFunc {
	return func(value interface{}) error {
		if f.Password != f.PasswordConfirmation {
			return errors.New("does not match the password")
		}
		return nil
	}
}

func (f RegistrationForm) Save() (*models.User, error) {
	existingUser, err := models.GetUserByEmail(f.Email)
	if err == nil && existingUser != nil {
		return nil, errors.New(constants.UserAlreadyExist)
	}

	hashedPassword, err := helpers.HashPassword(f.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email:          f.Email,
		HashedPassword: hashedPassword,
	}

	userID, err := models.CreateUser(user)
	if err != nil {
		return nil, err
	}

	user, err = models.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
