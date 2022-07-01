package webforms_test

import (
	"go-google-scraper-challenge/constants"
	webforms "go-google-scraper-challenge/lib/web/forms"
	. "go-google-scraper-challenge/test"

	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Web Registration Form", func() {
	Describe("Validate", func() {
		Context("given registration form with valid params", func() {
			It("returns NO error", func() {
				password := faker.Password()
				form := webforms.RegistrationForm{
					Email:                faker.Email(),
					Password:             password,
					PasswordConfirmation: password,
				}

				valid, err := form.Validate()

				Expect(valid).To(BeTrue())
				Expect(err).To(BeNil())
			})
		})

		Context("given registration form with INVALID params", func() {
			Context("email", func() {
				Context("given NO email", func() {
					It("returns an INVALID email error", func() {
						password := faker.Password()
						form := webforms.RegistrationForm{
							Email:                "",
							Password:             password,
							PasswordConfirmation: password,
						}

						valid, err := form.Validate()

						Expect(valid).To(BeFalse())
						Expect(err.Error()).To(Equal("Email: cannot be blank."))
					})
				})

				Context("given an INVALID email", func() {
					It("returns an INVALID email error", func() {
						password := faker.Password()
						form := webforms.RegistrationForm{
							Email:                "invalid",
							Password:             password,
							PasswordConfirmation: password,
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
						form := webforms.RegistrationForm{
							Email:                faker.Email(),
							Password:             "",
							PasswordConfirmation: "",
						}

						valid, err := form.Validate()

						Expect(valid).To(BeFalse())
						Expect(err.Error()).To(ContainSubstring("Password: cannot be blank"))
					})
				})
			})

			Context("password confirmation", func() {
				Context("given NO password confirmation", func() {
					It("returns an INVALID password confirmation error", func() {
						password := faker.Password()
						form := webforms.RegistrationForm{
							Email:                faker.Email(),
							Password:             password,
							PasswordConfirmation: "",
						}

						valid, err := form.Validate()

						Expect(valid).To(BeFalse())
						Expect(err.Error()).To(Equal("PasswordConfirmation: cannot be blank."))
					})
				})

				Context("given unmatch password confirmation", func() {
					It("returns an INVALID password confirmation error", func() {
						password := faker.Password()
						form := webforms.RegistrationForm{
							Email:                faker.Email(),
							Password:             password,
							PasswordConfirmation: faker.Password(),
						}

						valid, err := form.Validate()

						Expect(valid).To(BeFalse())
						Expect(err.Error()).To(Equal("PasswordConfirmation: does not match the password."))
					})
				})
			})
		})
	})

	Describe("Save", func() {
		Context("given valid email, password adn password confirmation", func() {
			It("returns the user", func() {
				email := faker.Email()
				password := faker.Password()
				form := webforms.RegistrationForm{
					Email:                email,
					Password:             password,
					PasswordConfirmation: password,
				}

				returnUser, err := form.Save()

				Expect(returnUser).NotTo(BeNil())
				Expect(returnUser.ID).To(BeNumerically(">", 0))
				Expect(returnUser.Email).To(Equal(email))
				Expect(err).To(BeNil())
			})
		})

		Context("given the user email already exist", func() {
			It("returns the user already exist error", func() {
				password := faker.Password()
				user := FabricateUser(faker.Email(), password)
				form := webforms.RegistrationForm{
					Email:                user.Email,
					Password:             password,
					PasswordConfirmation: password,
				}

				user, err := form.Save()

				Expect(user).To(BeNil())
				Expect(err.Error()).To(Equal(constants.UserAlreadyExist))
			})
		})
	})

	AfterEach(func() {
		CleanupDatabase([]string{"users"})
	})
})
