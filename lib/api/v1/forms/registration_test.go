package forms_test

import (
	"go-google-scraper-challenge/constants"
	"go-google-scraper-challenge/lib/api/v1/forms"
	"go-google-scraper-challenge/lib/models"
	. "go-google-scraper-challenge/test"

	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("API Registration Form", func() {
	Describe("#Save", func() {
		Context("given registration form with valid params", func() {
			It("returns a user ID", func() {
				authClient := FabricateAuthClient()
				form := forms.RegistrationForm{
					ClientID:     authClient.ClientID,
					ClientSecret: authClient.ClientSecret,
					Email:        faker.Email(),
					Password:     faker.Password(),
				}

				userID, err := form.Save()
				if err != nil {
					Fail("Failed to save form")
				}

				Expect(*userID).To(BeNumerically(">", 0))

				user, err := models.GetUserByID(*userID)
				if err != nil {
					Fail("Failed to find the user")
				}

				Expect(user.Email).To(Equal(form.Email))
			})

			It("returns NO error", func() {
				authClient := FabricateAuthClient()
				form := forms.RegistrationForm{
					ClientID:     authClient.ClientID,
					ClientSecret: authClient.ClientSecret,
					Email:        faker.Email(),
					Password:     faker.Password(),
				}

				_, err := form.Save()

				Expect(err).To(BeNil())
			})
		})

		Context("given registration form with INVALID params", func() {
			Context("authentication client", func() {
				Context("given NO client id", func() {
					It("returns an INVALID client ID error", func() {
						authClient := FabricateAuthClient()
						form := forms.RegistrationForm{
							ClientID:     "",
							ClientSecret: authClient.ClientSecret,
							Email:        faker.Email(),
							Password:     faker.Password(),
						}

						userID, err := form.Save()

						Expect(userID).To(BeNil())
						Expect(err.Error()).To(Equal("ClientID: cannot be blank."))
					})
				})

				Context("given NO client secret", func() {
					It("returns an INVALID client ID error", func() {
						authClient := FabricateAuthClient()
						form := forms.RegistrationForm{
							ClientID:     authClient.ClientID,
							ClientSecret: "",
							Email:        faker.Email(),
							Password:     faker.Password(),
						}

						userID, err := form.Save()

						Expect(userID).To(BeNil())
						Expect(err.Error()).To(Equal("ClientSecret: cannot be blank."))
					})
				})
			})

			Context("email", func() {
				Context("given email that already registered", func() {
					It("returns a duplicate email error", func() {
						user := FabricateUser(faker.Email(), faker.Password())
						authClient := FabricateAuthClient()
						form := forms.RegistrationForm{
							ClientID:     authClient.ClientID,
							ClientSecret: authClient.ClientSecret,
							Email:        user.Email,
							Password:     faker.Password(),
						}

						userID, err := form.Save()

						Expect(userID).To(BeNil())
						Expect(err.Error()).To(Equal(constants.UserAlreadyExist))
					})
				})

				Context("given NO email", func() {
					It("returns an INVALID email error", func() {
						authClient := FabricateAuthClient()
						form := forms.RegistrationForm{
							ClientID:     authClient.ClientID,
							ClientSecret: authClient.ClientSecret,
							Email:        "",
							Password:     faker.Password(),
						}

						userID, err := form.Save()

						Expect(userID).To(BeNil())
						Expect(err.Error()).To(Equal("Email: cannot be blank."))
					})
				})

				Context("given an INVALID email", func() {
					It("returns an INVALID email error", func() {
						authClient := FabricateAuthClient()
						form := forms.RegistrationForm{
							ClientID:     authClient.ClientID,
							ClientSecret: authClient.ClientSecret,
							Email:        "invalid",
							Password:     faker.Password(),
						}

						userID, err := form.Save()

						Expect(userID).To(BeNil())
						Expect(err.Error()).To(Equal("Email: must be a valid email address."))
					})
				})
			})

			Context("password", func() {
				Context("given NO password", func() {
					It("returns an INVALID password error", func() {
						authClient := FabricateAuthClient()
						form := forms.RegistrationForm{
							ClientID:     authClient.ClientID,
							ClientSecret: authClient.ClientSecret,
							Email:        faker.Email(),
							Password:     "",
						}

						userID, err := form.Save()

						Expect(userID).To(BeNil())
						Expect(err.Error()).To(Equal("Password: cannot be blank."))
					})
				})

				Context("given password length is less than 6", func() {
					It("returns an INVALID password error", func() {
						authClient := FabricateAuthClient()
						form := forms.RegistrationForm{
							ClientID:     authClient.ClientID,
							ClientSecret: authClient.ClientSecret,
							Email:        faker.Email(),
							Password:     "1234",
						}

						userID, err := form.Save()

						Expect(userID).To(BeNil())
						Expect(err.Error()).To(Equal("Password: the length must be between 6 and 50."))
					})
				})
			})
		})
	})

	AfterEach(func() {
		CleanupDatabase([]string{"users", "oauth2_clients", "oauth2_tokens"})
	})
})
