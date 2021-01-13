package forms

import (
	"log"

	"go-google-scraper-challenge/helpers"
	"go-google-scraper-challenge/models"

	"github.com/beego/beego/v2/core/validation"
)

type RegistrationForm struct {
	Email                string `valid:"Email; Required"`
	Password             string `valid:"Required"`
	PasswordConfirmation string `valid:"Required"`
}

func init() {
}

func (registrationForm *RegistrationForm) Valid(v *validation.Validation) {
	_, err := models.GetUserByEmail(registrationForm.Email)
	if err == nil {
		err := v.SetError("Email Already Exist", "User with this email already exist")
		if err == nil {
			log.Fatal("Failed to set error on validation")
		}
	}

	if registrationForm.Password != registrationForm.PasswordConfirmation {
		err := v.SetError("Password Mismatch", "Does not match the password confirmation")
		if err == nil {
			log.Fatal("Failed to set error on validation")
		}
	}
}

func (registrationForm *RegistrationForm) Save() (id int64, errors []error) {
	validation := validation.Validation{}

	valid, err := validation.Valid(&registrationForm)
	if err != nil {
		return -1, []error{err}
	}

	if !valid {
		errors := []error{}
		for _, err := range validation.Errors {
			errors = append(errors, err)
		}

		return -1, errors
	}

	user := models.User{
		Email:             registrationForm.Email,
		EncryptedPassword: helpers.EncryptPassword(registrationForm.Password),
	}

	userID, err := models.AddUser(&user)
	return userID, []error{err}
}
