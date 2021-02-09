package forms_test

import (
	"go-google-scraper-challenge/forms"
	"go-google-scraper-challenge/initializers"
	. "go-google-scraper-challenge/test/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Forms/SessionForm", func() {
	Describe("#Save", func() {
		Context("given session form with valid params", func() {
			It("returns user with NO error", func() {
				user := FabricateUser("dev@nimblehq.co", "password")
				form := forms.SessionForm{
					Email:    "dev@nimblehq.co",
					Password: "password",
				}

				currentUser, errors := form.Save()

				Expect(len(errors)).To(BeZero())
				Expect(currentUser.Id).To(Equal(user.Id))
			})
		})

		Context("given session form with INVALID params", func() {
			Context("given email is not an valid email", func() {
				It("returns a email is invalid error", func() {
					form := forms.SessionForm{
						Email:    "not an email",
						Password: "password",
					}

					user, errors := form.Save()

					Expect(errors[0].Error()).To(Equal("Email must be a valid email address"))
					Expect(user).To(BeNil())
				})
			})

			Context("given NO email", func() {
				It("returns an invalid email or password error", func() {
					form := forms.SessionForm{
						Email:    "",
						Password: "password",
					}

					user, errors := form.Save()

					Expect(errors[0].Error()).To(Equal("Email must be a valid email address"))
					Expect(user).To(BeNil())
				})
			})

			Context("given NO password", func() {
				It("returns an invalid email or password error", func() {
					form := forms.SessionForm{
						Email:    "dev@nimblehq.co",
						Password: "",
					}

					user, errors := form.Save()

					Expect(errors[0].Error()).To(Equal("Password can not be empty"))
					Expect(user).To(BeNil())
				})
			})
		})
	})

	AfterEach(func() {
		initializers.CleanupDatabase("users")
	})
})
