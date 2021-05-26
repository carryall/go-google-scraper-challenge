package forms

import (
	"mime/multipart"

	"go-google-scraper-challenge/constants"
	"go-google-scraper-challenge/helpers"
	"go-google-scraper-challenge/models"

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
		_ = v.SetError("File", constants.FileEmpty)

		return
	}

	if !uf.validateFileType() {
		_ = v.SetError("File", constants.FileTypeInvalid)

		return
	}

	if !uf.validateFileReadability() {
		_ = v.SetError("File", constants.FileUnreadable)

		return
	}

	if !uf.validateKeywordsLength() {
		_ = v.SetError("File", "File should contains between 1 to 1000 keywords")
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

func (uf *UploadForm) validateKeywordsLength() bool {
	return len(uf.keywords) > 0 && len(uf.keywords) <= 1000
}
