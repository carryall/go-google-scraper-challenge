package forms_test

import (
	"go-google-scraper-challenge/lib/api/v1/forms"
	. "go-google-scraper-challenge/test"

	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("API Result Form", func() {
	Describe("#Validate", func() {
		Context("given result form with valid params", func() {
			It("returns NO error", func() {
				file, fileHeader := GetMultipartFromFile("test/fixtures/files/valid.csv")
				user := FabricateUser(faker.Email(), faker.Password())

				form := forms.UploadForm{
					File:       file,
					FileHeader: fileHeader,
					UserID:     user.ID,
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
					UserID:     user.ID,
				}

				form.Validate()

				Expect(form.Keywords).To(HaveLen(1))
				Expect(form.Keywords[0]).To(Equal("ergonomic chair"))
			})
		})

		Context("given result form with INVALID params", func() {
			Context("given NO user ID", func() {
				It("returns an error", func() {
					file, fileHeader := GetMultipartFromFile("test/fixtures/files/valid.csv")
					form := forms.UploadForm{
						File:       file,
						FileHeader: fileHeader,
					}

					valid, err := form.Validate()

					Expect(valid).To(BeFalse())
					Expect(err.Error()).To(Equal("UserID: cannot be blank."))
				})
			})

			Context("given NO file", func() {
				It("returns an error", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					form := forms.UploadForm{
						UserID: user.ID,
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
						UserID:     user.ID,
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
						UserID:     user.ID,
					}

					valid, err := form.Validate()

					Expect(valid).To(BeFalse())
					Expect(err.Error()).To(Equal("File: wrong file type"))
				})
			})

			Context("given number of keyword is more than 1000", func() {
				It("returns an error", func() {
					file, fileHeader := GetMultipartFromFile("test/fixtures/files/invalid.csv")
					user := FabricateUser(faker.Email(), faker.Password())

					form := forms.UploadForm{
						File:       file,
						FileHeader: fileHeader,
						UserID:     user.ID,
					}

					valid, err := form.Validate()

					Expect(valid).To(BeFalse())
					Expect(err.Error()).To(Equal("Keywords: the length must be between 1 and 1000."))
				})
			})
		})
	})

	Describe("#Save", func() {
		Context("given result form with valid params", func() {
			It("returns NO error", func() {
				file, fileHeader := GetMultipartFromFile("test/fixtures/files/valid.csv")
				user := FabricateUser(faker.Email(), faker.Password())

				form := forms.UploadForm{
					File:       file,
					FileHeader: fileHeader,
					UserID:     user.ID,
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
					UserID:     user.ID,
				}

				resultIDs, err := form.Save()

				Expect(err).To(BeNil())
				Expect(resultIDs).To(HaveLen(1))
				for i := range resultIDs {
					Expect(resultIDs[i]).To(BeNumerically(">", 0))
				}
			})
		})

		Context("given result form with INVALID params", func() {
			Context("given NO user ID", func() {
				It("returns an error", func() {
					file, fileHeader := GetMultipartFromFile("test/fixtures/files/valid.csv")
					form := forms.UploadForm{
						File:       file,
						FileHeader: fileHeader,
					}

					resultIDs, err := form.Save()

					Expect(resultIDs).To(BeEmpty())
					Expect(err.Error()).To(Equal("UserID: cannot be blank."))
				})
			})

			Context("given NO file", func() {
				It("returns an error", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					form := forms.UploadForm{
						UserID: user.ID,
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
						UserID:     user.ID,
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
						UserID:     user.ID,
					}

					resultIDs, err := form.Save()

					Expect(resultIDs).To(BeEmpty())
					Expect(err.Error()).To(Equal("File: wrong file type"))
				})
			})

			Context("given number of keyword is more than 1000", func() {
				It("returns an error", func() {
					file, fileHeader := GetMultipartFromFile("test/fixtures/files/invalid.csv")
					user := FabricateUser(faker.Email(), faker.Password())
					form := forms.UploadForm{
						File:       file,
						FileHeader: fileHeader,
						UserID:     user.ID,
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