package forms

import (
	"log"

	"github.com/beego/beego/v2/core/validation"
)

type Registration struct {
	Email                string `valid:"Email; Required"`
	Password             string `valid:"Required"`
	PasswordConfirmation string `valid:"Required"`
}

func init() {
}

func (formParams *Registration) Valid(v *validation.Validation) {
	if formParams.Password != formParams.PasswordConfirmation {
		err := v.SetError("Password Mismatch", "Does not match the password confirmation")
		if err == nil {
			log.Fatal("Failed to set error on validation")
		}
	}
}
