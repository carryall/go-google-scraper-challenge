package forms_test

import (
	"go-google-scraper-challenge/constants"
	"go-google-scraper-challenge/forms"
	"go-google-scraper-challenge/initializers"
	. "go-google-scraper-challenge/tests/helpers"

	"github.com/beego/beego/v2/core/validation"
	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Forms/UploadForm", func() {
	Describe("#Valid", func() {
		Context("given upload form with valid params", func() {
			It("does NOT add error to validation", func() {
				file, fileHeader := GetMultipartFromFile("tests/fixtures/files/valid.csv")
				user := FabricateUser(faker.Email(), faker.Password())
				form := forms.UploadForm{
					File: file,
					FileHeader: fileHeader,
					User: user,
				}

				formValidation := validation.Validation{}
				form.Valid(&formValidation)

				Expect(len(formValidation.Errors)).To(BeZero())
			})
		})

		Context("given upload form with INVALID params", func() {
			Context("given NO file", func() {
				It("adds an error to validation", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					form := forms.UploadForm{
						User: user,
					}

					formValidation := validation.Validation{}
					form.Valid(&formValidation)

					Expect(len(formValidation.Errors)).To(Equal(1))
					Expect(formValidation.Errors[0].Key).To(Equal("File"))
					Expect(formValidation.Errors[0].Message).To(Equal(constants.FileEmpty))
				})
			})

			Context("given wrong file type", func() {
				It("adds an error to validation", func() {
					file, fileHeader := GetMultipartFromFile("tests/fixtures/files/text.txt")
					user := FabricateUser(faker.Email(), faker.Password())
					form := forms.UploadForm{
						File: file,
						FileHeader: fileHeader,
						User: user,
					}

					formValidation := validation.Validation{}
					form.Valid(&formValidation)

					Expect(len(formValidation.Errors)).To(Equal(1))
					Expect(formValidation.Errors[0].Key).To(Equal("File"))
					Expect(formValidation.Errors[0].Message).To(Equal(constants.FileTypeInvalid))
				})
			})

			Context("given an empty CSV file", func() {
				It("adds an error to validation", func() {
					file, fileHeader := GetMultipartFromFile("tests/fixtures/files/empty.csv")
					user := FabricateUser(faker.Email(), faker.Password())
					form := forms.UploadForm{
						File: file,
						FileHeader: fileHeader,
						User: user,
					}

					formValidation := validation.Validation{}
					form.Valid(&formValidation)

					Expect(len(formValidation.Errors)).To(Equal(1))
					Expect(formValidation.Errors[0].Key).To(Equal("File"))
					Expect(formValidation.Errors[0].Message).To(Equal("File should contains between 1 to 1000 keywords"))
				})
			})

			Context("given a CSV file that contains more than 1000 keywords", func() {
				It("adds an error to validation", func() {
					file, fileHeader := GetMultipartFromFile("tests/fixtures/files/invalid.csv")
					user := FabricateUser(faker.Email(), faker.Password())
					form := forms.UploadForm{
						File: file,
						FileHeader: fileHeader,
						User: user,
					}

					formValidation := validation.Validation{}
					form.Valid(&formValidation)

					Expect(len(formValidation.Errors)).To(Equal(1))
					Expect(formValidation.Errors[0].Key).To(Equal("File"))
					Expect(formValidation.Errors[0].Message).To(Equal("File should contains between 1 to 1000 keywords"))
				})
			})
		})
	})

	Describe("#Save", func() {
		Context("given upload form with a valid params", func() {
			It("returns keywords from the given file", func() {
				file, fileHeader := GetMultipartFromFile("tests/fixtures/files/valid.csv")
				user := FabricateUser(faker.Email(), faker.Password())
				form := forms.UploadForm{
					File: file,
					FileHeader: fileHeader,
					User: user,
				}
				expectedKeyword := []string{
					"ergonomic chair",
				}

				keywords, error := form.Save()

				Expect(error).To(BeNil())
				Expect(keywords).To(Equal(expectedKeyword))
			})
		})

		Context("given upload form with an INVALID params", func() {
			Context("given NO user", func() {
				It("returns an invalid user error", func() {
					file, fileHeader := GetMultipartFromFile("tests/fixtures/files/invalid.csv")
					form := forms.UploadForm{
						File: file,
						FileHeader: fileHeader,
					}

					_, error := form.Save()

					Expect(error.Error()).To(Equal("User can not be empty"))
				})
			})
		})
	})

	AfterEach(func() {
		initializers.CleanupDatabase([]string{"users"})
	})
})
