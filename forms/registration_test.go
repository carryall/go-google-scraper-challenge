package forms_test

import (
	"go-google-scraper-challenge/forms"
	. "go-google-scraper-challenge/helpers/test"
	"go-google-scraper-challenge/initializers"

	"github.com/beego/beego/v2/core/validation"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Forms/RegistrationForm", func() {
	Describe("#Valid", func() {
		Context("given registration form with valid params", func() {
			It("does NOT add error to validation", func() {
				form := forms.RegistrationForm{
					Email:                "dev@nimblehq.co",
					Password:             "password",
					PasswordConfirmation: "password",
				}

				formValidation := validation.Validation{}
				form.Valid(&formValidation)

				Expect(len(formValidation.Errors)).To(BeZero())
			})
		})

		Context("given registration form with INVALID params", func() {
			Context("given email that already registered", func() {
				It("adds duplicate email error to validation", func() {
					FabricateUser("dev@nimblehq.co", "password")

					form := forms.RegistrationForm{
						Email:                "dev@nimblehq.co",
						Password:             "password",
						PasswordConfirmation: "password",
					}

					formValidation := validation.Validation{}
					form.Valid(&formValidation)

					Expect(len(formValidation.Errors)).To(Equal(1))
					Expect(formValidation.Errors[0].Key).To(Equal("Email"))
					Expect(formValidation.Errors[0].Message).To(Equal("User with this email already exist"))
				})
			})

			Context("given password confirmation is NOT match with the password", func() {
				It("adds a mismatch password confirmation error to validation", func() {
					form := forms.RegistrationForm{
						Email:                "dev@nimblehq.co",
						Password:             "password",
						PasswordConfirmation: "does not match the password",
					}

					formValidation := validation.Validation{}
					form.Valid(&formValidation)

					Expect(len(formValidation.Errors)).To(Equal(1))
					Expect(formValidation.Errors[0].Key).To(Equal("PasswordConfirmation"))
					Expect(formValidation.Errors[0].Message).To(Equal("Password confirmation must match the password"))
				})
			})
		})
	})

	Describe("#Save", func() {
		Context("given registration form with valid params", func() {
			It("returns a user ID", func() {
				form := forms.RegistrationForm{
					Email:                "dev@nimblehq.co",
					Password:             "password",
					PasswordConfirmation: "password",
				}

				userID, errors := form.Save()
				if len(errors) > 0 {
					Fail("Failed to save form")
				}

				Expect(userID).NotTo(BeNil())
			})

			It("returns NO error", func() {
				form := forms.RegistrationForm{
					Email:                "dev@nimblehq.co",
					Password:             "password",
					PasswordConfirmation: "password",
				}

				_, errors := form.Save()

				Expect(len(errors)).To(BeZero())
			})
		})

		Context("given registration form with INVALID params", func() {
			Context("given email that already registered", func() {
				It("returns a duplicate email error", func() {
					FabricateUser("dev@nimblehq.co", "password")

					form := forms.RegistrationForm{
						Email:                "dev@nimblehq.co",
						Password:             "password",
						PasswordConfirmation: "password",
					}

					userID, errors := form.Save()

					Expect(userID).To(BeNil())
					Expect(errors[0].Error()).To(Equal("User with this email already exist"))
				})
			})

			Context("given NO email", func() {
				It("returns an invalid email error", func() {
					form := forms.RegistrationForm{
						Email:                "",
						Password:             "password",
						PasswordConfirmation: "password",
					}

					userID, errors := form.Save()

					Expect(userID).To(BeNil())
					Expect(errors[0].Error()).To(Equal("Email must be a valid email address"))
				})
			})

			Context("given an INVALID email", func() {
				It("returns an invalid email error", func() {
					form := forms.RegistrationForm{
						Email:                "invalid",
						Password:             "password",
						PasswordConfirmation: "password",
					}

					userID, errors := form.Save()

					Expect(userID).To(BeNil())
					Expect(errors[0].Error()).To(Equal("Email must be a valid email address"))
				})
			})

			Context("given NO password", func() {
				It("returns an invalid password error", func() {
					form := forms.RegistrationForm{
						Email:                "dev@nimblehq.co",
						Password:             "",
						PasswordConfirmation: "password",
					}

					userID, errors := form.Save()

					Expect(userID).To(BeNil())
					Expect(errors[0].Error()).To(Equal("Password can not be empty"))
				})
			})

			Context("given password length is less than 6", func() {
				It("returns an invalid password error", func() {
					form := forms.RegistrationForm{
						Email:                "dev@nimblehq.co",
						Password:             "1234",
						PasswordConfirmation: "password",
					}

					userID, errors := form.Save()

					Expect(userID).To(BeNil())
					Expect(errors[0].Error()).To(Equal("Password minimum size is 6"))
				})
			})

			Context("given NO password confirmation", func() {
				It("returns an invalid password confirmation error", func() {
					form := forms.RegistrationForm{
						Email:                "dev@nimblehq.co",
						Password:             "password",
						PasswordConfirmation: "",
					}

					userID, errors := form.Save()

					Expect(userID).To(BeNil())
					Expect(errors[0].Error()).To(Equal("PasswordConfirmation can not be empty"))
				})
			})

			Context("given password confirmation is length less than 6", func() {
				It("returns an invalid password error", func() {
					form := forms.RegistrationForm{
						Email:                "dev@nimblehq.co",
						Password:             "password",
						PasswordConfirmation: "1234",
					}

					userID, errors := form.Save()

					Expect(userID).To(BeNil())
					Expect(errors[0].Error()).To(Equal("PasswordConfirmation minimum size is 6"))
				})
			})

			Context("given password confirmation is NOT match with the password", func() {
				It("returns a mismatch password confirmation error", func() {
					form := forms.RegistrationForm{
						Email:                "dev@nimblehq.co",
						Password:             "password",
						PasswordConfirmation: "does not match the password",
					}

					userID, errors := form.Save()

					Expect(userID).To(BeNil())
					Expect(errors[0].Error()).To(Equal("Password confirmation must match the password"))
				})
			})
		})
	})

	AfterEach(func() {
		initializers.CleanupDatabase("users")
	})
})
