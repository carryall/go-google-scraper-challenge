package forms

import (
	"log"

	"go-google-scraper-challenge/helpers"
	"go-google-scraper-challenge/models"

	"github.com/beego/beego/v2/core/validation"
)

type RegistrationForm struct {
	Email                string `valid:"Email; Required"`
	Password             string `valid:"Required; MinSize(6)"`
	PasswordConfirmation string `valid:"Required; MinSize(6)"`
}

func init() {
}

// Valid adds custom validation to registration form, sets error when the validation failed.
func (registrationForm *RegistrationForm) Valid(v *validation.Validation) {
	// This will raise an error if user with the given email does not exist.
	_, err := models.GetUserByEmail(registrationForm.Email)
	if err == nil {
		validationError := v.SetError("Email", "User with this email already exist")
		if validationError == nil {
			log.Fatal("Failed to set error on validation")
		}
	}

	if registrationForm.Password != registrationForm.PasswordConfirmation {
		validationError := v.SetError("PasswordConfirmation", "Password confirmation must match the password")
		if validationError == nil {
			log.Fatal("Failed to set error on validation")
		}
	}
}

// Save validates registration form and adds a new User with email and password from the form,
// returns errors if validation failed or cannot add the user to database.
func (registrationForm RegistrationForm) Save() (id *int64, errors []error) {
	validation := validation.Validation{}

	valid, err := validation.Valid(&registrationForm)
	if err != nil {
		return nil, []error{err}
	}

	if !valid {
		errors := []error{}
		for _, err := range validation.Errors {
			errors = append(errors, err)
		}

		return nil, errors
	}

	user := models.User{
		Email:             registrationForm.Email,
		EncryptedPassword: helpers.EncryptPassword(registrationForm.Password),
	}

	userID, err := models.AddUser(&user)
	if err != nil {
		return nil, []error{err}
	}

	return &userID, []error{}
}
