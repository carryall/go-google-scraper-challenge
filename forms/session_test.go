package forms_test

import (
	apiforms "go-google-scraper-challenge/forms/api"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Forms/LoginForm", func() {
	Describe("#Save", func() {
		Context("given login form with valid params", func() {
			It("returns NO error", func() {
				form := apiforms.LoginForm{
					Username: "dev@nimblehq.co",
					Password: "password",
				}

				errors := form.Save()

				Expect(len(errors)).To(BeZero())
			})
		})

		Context("given login form with INVALID params", func() {
			Context("given email is not an valid email", func() {
				It("returns a email is invalid error", func() {
					form := apiforms.LoginForm{
						Username: "not an email",
						Password: "password",
					}

					errors := form.Save()

					Expect(errors[0].Error()).To(Equal("Email must be a valid email address"))
				})
			})

			Context("given NO email", func() {
				It("returns an invalid email or password error", func() {
					form := apiforms.LoginForm{
						Username: "",
						Password: "password",
					}

					errors := form.Save()

					Expect(errors[0].Error()).To(Equal("Incorrect email or password"))
				})
			})

			Context("given NO password", func() {
				It("returns an invalid email or password error", func() {
					form := apiforms.LoginForm{
						Username: "dev@nimblehq.co",
						Password: "",
					}

					errors := form.Save()

					Expect(errors[0].Error()).To(Equal("Incorrect email or password"))
				})
			})
		})
	})
})
