package forms

import (
	"log"
	"strings"

	"go-google-scraper-challenge/helpers"
	"go-google-scraper-challenge/models"

	"github.com/beego/beego/v2/core/validation"
)

type RegistrationForm struct {
	Email                string `form:"email,email" valid:"Email; Required"`
	Password             string `form:"password,password" valid:"Required; MinSize(6)"`
	PasswordConfirmation string `form:"password_confirmation,password" valid:"Required; MinSize(6)"`
}

func init() {
	lowerCaseErrorMessage := map[string]string{}
	for key, value := range validation.MessageTmpls {
		lowerCaseErrorMessage[key] = strings.ToLower(value)
	}

	validation.SetDefaultMessage(lowerCaseErrorMessage)
}

// Valid adds custom validation to registration form, sets error when the validation failed.
func (registrationForm *RegistrationForm) Valid(v *validation.Validation) {
	userExist := models.UserEmailAlreadyExist(registrationForm.Email)
	if userExist {
		validationError := v.SetError("Email", "User with this email already exist")
		if validationError == nil {
			log.Print("Failed to set error on validation")
		}
	}

	if registrationForm.Password != registrationForm.PasswordConfirmation {
		validationError := v.SetError("PasswordConfirmation", "Password confirmation must match the password")
		if validationError == nil {
			log.Print("Failed to set error on validation")
		}
	}
}

// Save validates registration form and adds a new User with email and password from the form,
// returns errors if validation failed or cannot add the user to database.
func (registrationForm RegistrationForm) Save() (user *models.User, errors []error) {
	validation := validation.Validation{}

	valid, err := validation.Valid(&registrationForm)
	if err != nil {
		return nil, []error{err}
	}

	if !valid {
		for _, err := range validation.Errors {
			errors = append(errors, err)
		}

		return nil, errors
	}

	hashedPassword, err := helpers.HashPassword(registrationForm.Password)
	if err != nil {
		return nil, []error{err}
	}

	user = &models.User{
		Email:             registrationForm.Email,
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
