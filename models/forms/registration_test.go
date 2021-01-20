package forms_test

import (
	"go-google-scraper-challenge/initializers"
	"go-google-scraper-challenge/models/forms"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Forms/RegistrationForm", func() {
	Describe("#Valid", func() {

	})

	Describe("#Save", func() {
		Context("given registration form with valid params", func() {
			It("returns a user ID", func() {
				form := forms.RegistrationForm{
					Email:                "dev@nimblehq.co",
					Password:             "password",
					PasswordConfirmation: "password",
				}

				userID, _ := form.Save()

				Expect(userID).NotTo(BeNil())
			})

			It("returns NO error", func() {
				form := forms.RegistrationForm{
					Email:                "dev@nimblehq.co",
					Password:             "password",
					PasswordConfirmation: "password",
				}

				_, errors := form.Save()

				Expect(len(errors)).To(Equal(0))
			})
		})

		Context("given registration form with INVALID params", func() {
			Context("given email that already registered", func() {
				It("returns a duplicate email error", func() {
					form1 := forms.RegistrationForm{
						Email:                "dev@nimblehq.co",
						Password:             "password",
						PasswordConfirmation: "password",
					}
					form1.Save()

					form2 := forms.RegistrationForm{
						Email:                "dev@nimblehq.co",
						Password:             "password",
						PasswordConfirmation: "password",
					}

					_, errors := form2.Save()

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

					_, errors := form.Save()

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

					_, errors := form.Save()

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

					_, errors := form.Save()

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

					_, errors := form.Save()

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

					_, errors := form.Save()

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

					_, errors := form.Save()

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

					_, errors := form.Save()

					Expect(errors[0].Error()).To(Equal("Password confirmation must match the password"))
				})
			})
		})
	})

	AfterEach(func() {
		initializers.CleanupDatabase("user")
	})
})
