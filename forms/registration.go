package forms

import (
	"go-google-scraper-challenge/constants"
	"go-google-scraper-challenge/helpers"
	"go-google-scraper-challenge/models"

	"github.com/beego/beego/v2/core/validation"
)

type RegistrationForm struct {
	Email                string `form:"email" valid:"Email; Required"`
	Password             string `form:"password" valid:"Required; MinSize(6)"`
	PasswordConfirmation string `form:"password_confirmation" valid:"Required; MinSize(6)"`
}

// Valid adds custom validation to registration form, sets error when the validation failed.
func (rf *RegistrationForm) Valid(v *validation.Validation) {
	userExist := models.UserEmailAlreadyExist(rf.Email)
	if userExist {
		_ = v.SetError("Email", constants.UserAlreadyExist)
	}

	if rf.Password != rf.PasswordConfirmation {
		_ = v.SetError("PasswordConfirmation", constants.PasswordConfirmNotMatch)
	}
}

// Save validates registration form and adds a new User with email and password from the form,
// returns errors if validation failed or cannot add the user to database.
func (rf *RegistrationForm) Save() (*models.User, []error) {
	validation := validation.Validation{}

	valid, err := validation.Valid(rf)
	if err != nil {
		return nil, []error{err}
	}

	if !valid {
		var errors []error
		for _, err := range validation.Errors {
			errors = append(errors, err)
		}

		return nil, errors
	}

	hashedPassword, err := helpers.HashPassword(rf.Password)
	if err != nil {
		return nil, []error{err}
	}

	user := &models.User{
		Email:          rf.Email,
		HashedPassword: hashedPassword,
	}

	userID, err := models.CreateUser(user)
	if err != nil {
		return nil, []error{err}
	}

	user, err = models.GetUserById(userID)
	if err != nil {
		return nil, []error{err}
	}

	return user, []error{}
}
