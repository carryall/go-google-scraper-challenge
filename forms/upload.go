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
	keywords   []string
}

// Valid adds custom validation to upload form, sets error when the validation failed.
func (uf *UploadForm) Valid(v *validation.Validation) {
	if !uf.validateFilePresence() {
		err := v.SetError("File", "File cannot be empty")
		if err == nil {
			logs.Info("Failed to set error on validation")
		}

		return
	}

	if !uf.validateFileType() {
		err := v.SetError("File", "Incorrect file type")
		if err == nil {
			logs.Info("Failed to set error on validation")
		}

		return
	}

	if !uf.validateFileReadability() {
		err := v.SetError("File", "Unreadable file")
		if err == nil {
			logs.Info("Failed to set error on validation")
		}

		return
	}

	if !uf.validateKeywordsNotEmpty() {
		err := v.SetError("File", "File should contains at least one keyword")
		if err == nil {
			logs.Info("Failed to set error on validation")
		}

		return
	}

	if !uf.validateKeywordsCountNotExceed() {
		err := v.SetError("File", "File contains too many keywords")
		if err == nil {
			logs.Info("Failed to set error on validation")
		}
	}
}

// Save validates upload form, returns errors if validation failed.
func (uf *UploadForm) Save() ([]string, error) {
	validation := validation.Validation{}

	valid, err := validation.Valid(uf)
	if err != nil {
		return nil, err
	}

	if !valid {
		return nil, validation.Errors[0]
	} else {
		return uf.keywords, nil
	}
}

func (uf *UploadForm) validateFilePresence() bool {
	return uf.File != nil && uf.FileHeader != nil
}

func (uf *UploadForm) validateFileType() bool {
	fileType := helpers.GetFileType(uf.FileHeader)

	return fileType == "text/csv"
}

func (uf *UploadForm) validateFileReadability() bool {
	keywords, err := helpers.GetFileContent(uf.File)
	uf.keywords = keywords

	return err == nil
}

func (uf *UploadForm) validateKeywordsNotEmpty() bool {
	return len(uf.keywords) > 0
}

func (uf *UploadForm) validateKeywordsCountNotExceed() bool {
	return len(uf.keywords) <= 1000
}
