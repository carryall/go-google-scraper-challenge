package forms

import (
	"mime/multipart"

	"go-google-scraper-challenge/helpers"
	"go-google-scraper-challenge/models"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
)

type UploadForm struct {
	File       multipart.File
	FileHeader *multipart.FileHeader
	User       *models.User          `valid:"Required"`
	Keywords   []string
}

// Valid adds custom validation to upload form, sets error when the validation failed.
func (uf *UploadForm) Valid(v *validation.Validation) {
	if uf.File == nil || uf.FileHeader == nil {
		err := v.SetError("File", "File cannot be empty")
		if err == nil {
			logs.Info("Failed to set error on validation")
		}
	} else {
		fileType := helpers.GetFileType(uf.FileHeader)
		if fileType != "text/csv" {
			err := v.SetError("File", "Incorrect file type")
			if err == nil {
				logs.Info("Failed to set error on validation")
			}
		}

		keywords, err := helpers.GetFileContent(uf.File)
		if err != nil {
			err := v.SetError("File", "Unreadable files")
			if err == nil {
				logs.Info("Failed to set error on validation")
			}
		}

		if len(keywords) > 1000 {
			err := v.SetError("File", "File contains too many keywords")
			if err == nil {
				logs.Info("Failed to set error on validation")
			}
		} else {
			uf.Keywords = keywords
		}
	}
}

// Save validates upload form, returns errors if validation failed.
func (uf *UploadForm) Save() ([]string, []error) {
	validation := validation.Validation{}

	valid, err := validation.Valid(uf)
	if err != nil {
		return nil, []error{err}
	}

	if !valid {
		var errs []error
		for _, err := range validation.Errors {
			errs = append(errs, err)
		}

		return nil, errs
	} else {
		return uf.Keywords, nil
	}
}
