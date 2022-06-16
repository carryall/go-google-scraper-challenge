package forms

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type ResultSearchForm struct {
	Keyword string `json:"keyword"`
}

func (f ResultSearchForm) Validate() error {
	return validation.ValidateStruct(&f,
		validation.Field(&f.Keyword),
	)
}
