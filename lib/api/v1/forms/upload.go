package forms

import (
	"errors"
	"go-google-scraper-challenge/helpers"
	"go-google-scraper-challenge/lib/models"
	"mime/multipart"

	validation "github.com/go-ozzo/ozzo-validation"
)

type UploadForm struct {
	File       multipart.File
	FileHeader *multipart.FileHeader
	UserID     int64
	Keywords   []string
}

func (f *UploadForm) Validate() (valid bool, err error) {
	err = validation.ValidateStruct(f,
		validation.Field(&f.File, validation.Required),
		validation.Field(&f.FileHeader, validation.Required),
		validation.Field(&f.UserID, validation.Required),
	)

	if err != nil {
		return false, err
	}

	_, err = f.validateFileType()
	if err != nil {
		return false, err
	}

	f.Keywords, err = f.validateFileReadability()
	if err != nil {
		return false, err
	}

	err = validation.ValidateStruct(f, validation.Field(&f.Keywords, validation.Required, validation.Length(1, 1000)))
	if err != nil {
		return false, err
	}

	return true, nil
}

func (f *UploadForm) Save() ([]int64, error) {
	_, err := f.Validate()
	if err != nil {
		return []int64{}, err
	}

	keywords := f.Keywords

	results := []models.Result{}
	for i := range keywords {
		results = append(results, models.Result{
			Keyword: keywords[i],
			UserID:  f.UserID,
		})
	}

	resultIDs, err := models.CreateResults(&results)
	if err != nil {
		return []int64{}, err
	}

	return resultIDs, nil
}

func (f UploadForm) validateFileType() (ok bool, err error) {
	fileType := helpers.GetFileType(f.FileHeader)

	if fileType != "text/csv" {
		// TODO: Update this once the PR about improving the error merged
		return false, errors.New("File: wrong file type")
	}

	return true, nil
}

func (f UploadForm) validateFileReadability() (keywords []string, err error) {
	keywords, err = helpers.GetFileContent(f.File)

	if err != nil {
		// TODO: Update this once the PR about improving the error merged
		return []string{}, errors.New("File: is unreadable")
	}

	return keywords, nil
}
