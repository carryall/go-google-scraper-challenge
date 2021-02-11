package initializers

import (
	"strings"

	"github.com/beego/beego/v2/core/validation"
)

func SetLowercaseValidationErrors() {
	lowerCaseErrorMessage := map[string]string{}
	for key, value := range validation.MessageTmpls {
		lowerCaseErrorMessage[key] = strings.ToLower(value)
	}

	validation.SetDefaultMessage(lowerCaseErrorMessage)
}
