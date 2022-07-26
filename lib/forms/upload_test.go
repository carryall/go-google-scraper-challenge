package forms_test

import (
	"go-google-scraper-challenge/lib/forms"
	"go-google-scraper-challenge/lib/models"
	. "go-google-scraper-challenge/test"

	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("API Upload Form", func() {
	Describe("#Validate", func() {
		Context("given upload form with valid params", func() {
			It("returns NO error", func() {
				file, fileHeader := GetMultipartFromFile("test/fixtures/files/valid.csv")
				user := FabricateUser(faker.Email(), faker.Password())

				form := forms.UploadForm{
					File:       file,
					FileHeader: fileHeader,
					User:       user,
				}

				valid, err := form.Validate()

				Expect(valid).To(BeTrue())
				Expect(err).To(BeNil())
			})

			It("assigns the keywords", func() {
				file, fileHeader := GetMultipartFromFile("test/fixtures/files/valid.csv")
				user := FabricateUser(faker.Email(), faker.Password())

				form := forms.UploadForm{
					File:       file,
					FileHeader: fileHeader,
					User:       user,
				}

				_, err := form.Validate()

				Expect(err).To(BeNil())
				Expect(form.Keywords).To(HaveLen(2))
				Expect(form.Keywords[0]).To(Equal("ergonomic chair"))
				Expect(form.Keywords[1]).To(Equal("mechanical keyboard"))
			})
		})

		Context("given upload form with INVALID params", func() {
			Context("given NO user", func() {
				It("returns an error", func() {
					file, fileHeader := GetMultipartFromFile("test/fixtures/files/valid.csv")
					form := forms.UploadForm{
						File:       file,
						FileHeader: fileHeader,
					}

					valid, err := form.Validate()

					Expect(valid).To(BeFalse())
					Expect(err.Error()).To(Equal("User: cannot be blank."))
				})
			})

			Context("given NO file", func() {
				It("returns an error", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					form := forms.UploadForm{
						User: user,
					}

					valid, err := form.Validate()

					Expect(valid).To(BeFalse())
					Expect(err.Error()).To(Equal("File: cannot be blank; FileHeader: cannot be blank."))
				})
			})

			Context("given a blank file", func() {
				It("returns an error", func() {
					file, fileHeader := GetMultipartFromFile("test/fixtures/files/empty.csv")
					user := FabricateUser(faker.Email(), faker.Password())

					form := forms.UploadForm{
						File:       file,
						FileHeader: fileHeader,
						User:       user,
					}

					valid, err := form.Validate()

					Expect(valid).To(BeFalse())
					Expect(err.Error()).To(Equal("Keywords: cannot be blank."))
				})
			})

			Context("given a file with invalid type", func() {
				It("returns an error", func() {
					file, fileHeader := GetMultipartFromFile("test/fixtures/files/text.txt")
					user := FabricateUser(faker.Email(), faker.Password())

					form := forms.UploadForm{
						File:       file,
						FileHeader: fileHeader,
						User:       user,
					}

					valid, err := form.Validate()

					Expect(valid).To(BeFalse())
					Expect(err.Error()).To(Equal("FileHeader: invalid file type."))
				})
			})

			Context("given number of keyword is more than 1000", func() {
				It("returns an error", func() {
					file, fileHeader := GetMultipartFromFile("test/fixtures/files/invalid.csv")
					user := FabricateUser(faker.Email(), faker.Password())

					form := forms.UploadForm{
						File:       file,
						FileHeader: fileHeader,
						User:       user,
					}

					valid, err := form.Validate()

					Expect(valid).To(BeFalse())
					Expect(err.Error()).To(Equal("Keywords: the length must be between 1 and 1000."))
				})
			})
		})
	})

	Describe("#Save", func() {
		Context("given upload form with valid params", func() {
			It("returns NO error", func() {
				file, fileHeader := GetMultipartFromFile("test/fixtures/files/valid.csv")
				user := FabricateUser(faker.Email(), faker.Password())

				form := forms.UploadForm{
					File:       file,
					FileHeader: fileHeader,
					User:       user,
				}

				_, err := form.Save()

				Expect(err).To(BeNil())
			})

			It("returns list of the result IDs", func() {
				file, fileHeader := GetMultipartFromFile("test/fixtures/files/valid.csv")
				user := FabricateUser(faker.Email(), faker.Password())

				form := &forms.UploadForm{
					File:       file,
					FileHeader: fileHeader,
					User:       user,
				}

				resultIDs, err := form.Save()

				Expect(err).To(BeNil())
				Expect(resultIDs).To(HaveLen(2))
				for i := range resultIDs {
					Expect(resultIDs[i]).To(BeNumerically(">", 0))
				}
			})
		})

		Context("given upload form with INVALID params", func() {
			Context("given NO user", func() {
				It("returns an error", func() {
					file, fileHeader := GetMultipartFromFile("test/fixtures/files/valid.csv")
					form := forms.UploadForm{
						File:       file,
						FileHeader: fileHeader,
					}

					resultIDs, err := form.Save()

					Expect(resultIDs).To(BeEmpty())
					Expect(err.Error()).To(Equal("User: cannot be blank."))
				})
			})

			Context("given an INVALID user", func() {
				It("returns an error", func() {
					file, fileHeader := GetMultipartFromFile("test/fixtures/files/valid.csv")
					invalidUser := models.User{}
					form := forms.UploadForm{
						File:       file,
						FileHeader: fileHeader,
						User:       &invalidUser,
					}

					resultIDs, err := form.Save()

					Expect(resultIDs).To(BeEmpty())
					Expect(err.Error()).To(Equal("User: record not found."))
				})
			})

			Context("given NO file", func() {
				It("returns an error", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					form := forms.UploadForm{
						User: user,
					}

					resultIDs, err := form.Save()

					Expect(resultIDs).To(BeEmpty())
					Expect(err.Error()).To(Equal("File: cannot be blank; FileHeader: cannot be blank."))
				})
			})

			Context("given a blank file", func() {
				It("returns an error", func() {
					file, fileHeader := GetMultipartFromFile("test/fixtures/files/empty.csv")
					user := FabricateUser(faker.Email(), faker.Password())
					form := forms.UploadForm{
						File:       file,
						FileHeader: fileHeader,
						User:       user,
					}

					resultIDs, err := form.Save()

					Expect(resultIDs).To(BeEmpty())
					Expect(err.Error()).To(Equal("Keywords: cannot be blank."))
				})
			})

			Context("given a file with wrong file type", func() {
				It("returns an error", func() {
					file, fileHeader := GetMultipartFromFile("test/fixtures/files/text.txt")
					user := FabricateUser(faker.Email(), faker.Password())
					form := forms.UploadForm{
						File:       file,
						FileHeader: fileHeader,
						User:       user,
					}

					resultIDs, err := form.Save()

					Expect(resultIDs).To(BeEmpty())
					Expect(err.Error()).To(Equal("FileHeader: invalid file type."))
				})
			})

			Context("given number of keyword is more than 1000", func() {
				It("returns an error", func() {
					file, fileHeader := GetMultipartFromFile("test/fixtures/files/invalid.csv")
					user := FabricateUser(faker.Email(), faker.Password())
					form := forms.UploadForm{
						File:       file,
						FileHeader: fileHeader,
						User:       user,
					}

					resultIDs, err := form.Save()

					Expect(resultIDs).To(BeEmpty())
					Expect(err.Error()).To(Equal("Keywords: the length must be between 1 and 1000."))
				})
			})
		})
	})

	AfterEach(func() {
		CleanupDatabase([]string{"users", "results"})
	})
})
