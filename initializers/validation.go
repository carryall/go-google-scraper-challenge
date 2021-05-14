package initializers

import (
	"strings"

	"github.com/beego/beego/v2/core/validation"
)

func SetLowercaseValidationErrors() {
	lowerCaseErrorMessage := map[string]string{}
	for k, v := range validation.MessageTmpls {
		lowerCaseErrorMessage[k] = strings.ToLower(v)
	}

	validation.SetDefaultMessage(lowerCaseErrorMessage)
}
