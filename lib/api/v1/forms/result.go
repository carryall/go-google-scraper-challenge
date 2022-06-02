package forms

import (
	"go-google-scraper-challenge/lib/models"

	validation "github.com/go-ozzo/ozzo-validation"
)

type ResultForm struct {
	Keywords []string
	UserID   int64
}

func (f ResultForm) Validate() (valid bool, err error) {
	err = validation.ValidateStruct(&f,
		validation.Field(&f.Keywords, validation.Required, validation.Length(1, 1000)),
		validation.Field(&f.UserID, validation.Required),
	)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (f ResultForm) Save() ([]int64, error) {
	_, err := f.Validate()
	if err != nil {
		return []int64{}, err
	}

	results := []models.Result{}
	for i := range f.Keywords {
		result := &models.Result{Keyword: f.Keywords[i], UserID: f.UserID}
		results = append(results, *result)
	}

	resultIDs, err := models.CreateResults(&results)
	if err != nil {
		return []int64{}, err
	}

	return resultIDs, nil
}
