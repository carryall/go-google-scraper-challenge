package forms

import (
	"errors"
	"mime/multipart"

	"go-google-scraper-challenge/helpers"
	"go-google-scraper-challenge/lib/models"

	validation "github.com/go-ozzo/ozzo-validation"
)

type UploadForm struct {
	File       multipart.File
	FileHeader *multipart.FileHeader
	User       *models.User
	Keywords   []string
}

func (f *UploadForm) Validate() (valid bool, err error) {
	err = validation.ValidateStruct(f,
		validation.Field(&f.File, validation.Required),
		validation.Field(&f.FileHeader, validation.Required, validation.By(f.validateFileType())),
		validation.Field(&f.User, validation.Required, validation.By(f.validateUser())),
	)

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
			User:    f.User,
		})
	}

	resultIDs, err := models.CreateResults(&results)
	if err != nil {
		return []int64{}, err
	}

	return resultIDs, nil
}

func (f UploadForm) validateUser() validation.RuleFunc {
	return func(value interface{}) error {
		_, err := models.GetUserByID(f.User.ID)

		if err != nil {
			return err
		}
		return nil
	}
}

func (f UploadForm) validateFileType() validation.RuleFunc {
	return func(value interface{}) error {
		fileHeader, _ := value.(*multipart.FileHeader)
		fileType := helpers.GetFileType(fileHeader)

		if fileType != "text/csv" {
			return errors.New("invalid file type")
		}
		return nil
	}
}

func (f UploadForm) validateFileReadability() (keywords []string, err error) {
	keywords, err = helpers.GetFileContent(f.File)

	if err != nil {
		return []string{}, err
	}

	return keywords, nil
}
