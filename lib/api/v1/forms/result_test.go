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
				user := FabricateUser(faker.Email(), faker.Password())
				keywords := []string{}
				for i := 0; i < 10; i++ {
					keywords = append(keywords, faker.Word())
				}

				form := forms.ResultForm{
					UserID:   user.ID,
					Keywords: keywords,
				}

				valid, err := form.Validate()

				Expect(valid).To(BeTrue())
				Expect(err).To(BeNil())
			})
		})

		Context("given result form with INVALID params", func() {
			Context("given NO user ID", func() {
				It("returns an error", func() {
					keywords := []string{}
					for i := 0; i < 10; i++ {
						keywords = append(keywords, faker.Word())
					}
					form := forms.ResultForm{
						Keywords: keywords,
					}

					valid, err := form.Validate()

					Expect(valid).To(BeFalse())
					Expect(err.Error()).To(Equal("UserID: cannot be blank."))
				})
			})

			Context("given NO keyword", func() {
				It("returns an error", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					form := forms.ResultForm{
						UserID: user.ID,
					}

					valid, err := form.Validate()

					Expect(valid).To(BeFalse())
					Expect(err.Error()).To(Equal("Keywords: cannot be blank."))
				})
			})

			Context("given a blank keyword list", func() {
				It("returns an error", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					keywords := []string{}
					form := forms.ResultForm{
						UserID:   user.ID,
						Keywords: keywords,
					}

					valid, err := form.Validate()

					Expect(valid).To(BeFalse())
					Expect(err.Error()).To(Equal("Keywords: cannot be blank."))
				})
			})

			Context("given number of keyword is more than 1000", func() {
				It("returns an error", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					keywords := []string{}
					for i := 0; i < 1001; i++ {
						keywords = append(keywords, faker.Word())
					}
					form := forms.ResultForm{
						UserID:   user.ID,
						Keywords: keywords,
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
				user := FabricateUser(faker.Email(), faker.Password())
				keywords := []string{}
				for i := 0; i < 10; i++ {
					keywords = append(keywords, faker.Word())
				}

				form := forms.ResultForm{
					UserID:   user.ID,
					Keywords: keywords,
				}

				_, err := form.Save()

				Expect(err).To(BeNil())
			})

			It("returns list of the result IDs", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				keywords := []string{}
				for i := 0; i < 10; i++ {
					keywords = append(keywords, faker.Word())
				}

				form := forms.ResultForm{
					UserID:   user.ID,
					Keywords: keywords,
				}

				resultIDs, _ := form.Save()

				Expect(resultIDs).To(HaveLen(10))
				for i := range resultIDs {
					Expect(resultIDs[i]).To(BeNumerically(">", 0))
				}
			})
		})

		Context("given result form with INVALID params", func() {
			Context("given NO user ID", func() {
				It("returns an error", func() {
					keywords := []string{}
					for i := 0; i < 10; i++ {
						keywords = append(keywords, faker.Word())
					}
					form := forms.ResultForm{
						Keywords: keywords,
					}

					resultIDs, err := form.Save()

					Expect(resultIDs).To(BeEmpty())
					Expect(err.Error()).To(Equal("UserID: cannot be blank."))
				})
			})

			Context("given NO keyword", func() {
				It("returns an error", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					form := forms.ResultForm{
						UserID: user.ID,
					}

					resultIDs, err := form.Save()

					Expect(resultIDs).To(BeEmpty())
					Expect(err.Error()).To(Equal("Keywords: cannot be blank."))
				})
			})

			Context("given a blank keyword list", func() {
				It("returns an error", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					keywords := []string{}
					form := forms.ResultForm{
						UserID:   user.ID,
						Keywords: keywords,
					}

					resultIDs, err := form.Save()

					Expect(resultIDs).To(BeEmpty())
					Expect(err.Error()).To(Equal("Keywords: cannot be blank."))
				})
			})

			Context("given number of keyword is more than 1000", func() {
				It("returns an error", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					keywords := []string{}
					for i := 0; i < 1001; i++ {
						keywords = append(keywords, faker.Word())
					}
					form := forms.ResultForm{
						UserID:   user.ID,
						Keywords: keywords,
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
