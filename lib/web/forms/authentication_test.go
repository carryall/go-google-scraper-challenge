package webforms_test

import (
	"go-google-scraper-challenge/constants"
	webforms "go-google-scraper-challenge/lib/web/forms"
	. "go-google-scraper-challenge/test"

	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("API Authentication Form", func() {
	Describe("Validate", func() {
		Context("given authentication form with valid params", func() {
			It("returns NO error", func() {
				password := faker.Password()
				user := FabricateUser(faker.Email(), password)
				form := webforms.AuthenticationForm{
					Email:    user.Email,
					Password: password,
				}

				valid, err := form.Validate()

				Expect(valid).To(BeTrue())
				Expect(err).To(BeNil())
			})
		})

		Context("given authentication form with INVALID params", func() {
			Context("email", func() {
				Context("given NO email", func() {
					It("returns an INVALID email error", func() {
						password := faker.Password()
						FabricateUser(faker.Email(), password)
						form := webforms.AuthenticationForm{
							Email:    "",
							Password: password,
						}

						valid, err := form.Validate()

						Expect(valid).To(BeFalse())
						Expect(err.Error()).To(Equal("Email: cannot be blank."))
					})
				})

				Context("given an INVALID email", func() {
					It("returns an INVALID email error", func() {
						password := faker.Password()
						FabricateUser(faker.Email(), password)
						form := webforms.AuthenticationForm{
							Email:    "invalid",
							Password: password,
						}

						valid, err := form.Validate()

						Expect(valid).To(BeFalse())
						Expect(err.Error()).To(Equal("Email: must be a valid email address."))
					})
				})
			})

			Context("password", func() {
				Context("given NO password", func() {
					It("returns an INVALID password error", func() {
						user := FabricateUser(faker.Email(), faker.Password())
						form := webforms.AuthenticationForm{
							Email:    user.Email,
							Password: "",
						}

						valid, err := form.Validate()

						Expect(valid).To(BeFalse())
						Expect(err.Error()).To(Equal("Password: cannot be blank."))
					})
				})

			})
		})
	})

	Describe("Save", func() {
		Context("given valid email and password", func() {
			It("returns the user", func() {
				password := faker.Password()
				user := FabricateUser(faker.Email(), password)
				form := webforms.AuthenticationForm{
					Email:    user.Email,
					Password: password,
				}

				returnUser, err := form.Save()

				Expect(returnUser).NotTo(BeNil())
				Expect(returnUser.ID).To(Equal(user.ID))
				Expect(err).To(BeNil())
			})
		})

		Context("given the user email does NOT exist", func() {
			It("returns the sign in error", func() {
				password := faker.Password()
				FabricateUser(faker.Email(), password)
				form := webforms.AuthenticationForm{
					Email:    "invalid@email.com",
					Password: password,
				}

				user, err := form.Save()

				Expect(user).To(BeNil())
				Expect(err.Error()).To(Equal(constants.SignInFail))
			})
		})

		Context("given a wrong password", func() {
			It("returns the sign in error", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				form := webforms.AuthenticationForm{
					Email:    user.Email,
					Password: "wrong",
				}

				user, err := form.Save()

				Expect(user).To(BeNil())
				Expect(err.Error()).To(Equal(constants.SignInFail))
			})
		})
	})

	AfterEach(func() {
		CleanupDatabase([]string{"users"})
	})
})
