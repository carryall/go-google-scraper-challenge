package forms_test

import (
	"go-google-scraper-challenge/constants"
	"go-google-scraper-challenge/lib/api/v1/forms"
	. "go-google-scraper-challenge/test"

	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("API Authentication Form", func() {
	Describe("#Save", func() {
		Context("given authentication form with valid params", func() {
			It("returns NO error", func() {
				authClient := FabricateAuthClient()
				password := faker.Password()
				user := FabricateUser(faker.Email(), password)
				form := forms.AuthenticationForm{
					ClientID:     authClient.ClientID,
					ClientSecret: authClient.ClientSecret,
					Email:        user.Email,
					Password:     password,
					GrantType:    "password",
				}

				err := form.Save()

				Expect(err).To(BeNil())
			})
		})

		Context("given registration form with INVALID params", func() {
			Context("authentication client", func() {
				Context("given NO client id", func() {
					It("returns an INVALID client ID error", func() {
						authClient := FabricateAuthClient()
						password := faker.Password()
						user := FabricateUser(faker.Email(), password)
						form := forms.AuthenticationForm{
							ClientID:     "",
							ClientSecret: authClient.ClientSecret,
							Email:        user.Email,
							Password:     password,
							GrantType:    "password",
						}

						err := form.Save()

						Expect(err.Error()).To(Equal("ClientID: cannot be blank."))
					})
				})

				Context("given NO client secret", func() {
					It("returns an INVALID client ID error", func() {
						authClient := FabricateAuthClient()
						password := faker.Password()
						user := FabricateUser(faker.Email(), password)
						form := forms.AuthenticationForm{
							ClientID:     authClient.ClientID,
							ClientSecret: "",
							Email:        user.Email,
							Password:     password,
							GrantType:    "password",
						}

						err := form.Save()

						Expect(err.Error()).To(Equal("ClientSecret: cannot be blank."))
					})
				})

				Context("given NO grant type", func() {
					It("returns an INVALID grant type error", func() {
						authClient := FabricateAuthClient()
						password := faker.Password()
						user := FabricateUser(faker.Email(), password)
						form := forms.AuthenticationForm{
							ClientID:     authClient.ClientID,
							ClientSecret: authClient.ClientSecret,
							Email:        user.Email,
							Password:     password,
						}

						err := form.Save()

						Expect(err.Error()).To(Equal("GrantType: cannot be blank."))
					})
				})
			})

			Context("email", func() {
				Context("given email that does NOT belongs to any user", func() {
					It("returns an record not found error", func() {
						authClient := FabricateAuthClient()
						form := forms.AuthenticationForm{
							ClientID:     authClient.ClientID,
							ClientSecret: authClient.ClientSecret,
							Email:        faker.Email(),
							Password:     faker.Password(),
							GrantType:    "password",
						}

						err := form.Save()

						Expect(err.Error()).To(Equal(constants.UserDoesNotExist))
					})
				})

				Context("given NO email", func() {
					It("returns an INVALID email error", func() {
						authClient := FabricateAuthClient()
						password := faker.Password()
						FabricateUser(faker.Email(), password)
						form := forms.AuthenticationForm{
							ClientID:     authClient.ClientID,
							ClientSecret: authClient.ClientSecret,
							Email:        "",
							Password:     password,
							GrantType:    "password",
						}

						err := form.Save()

						Expect(err.Error()).To(Equal("Email: cannot be blank."))
					})
				})

				Context("given an INVALID email", func() {
					It("returns an INVALID email error", func() {
						authClient := FabricateAuthClient()
						password := faker.Password()
						FabricateUser(faker.Email(), password)
						form := forms.AuthenticationForm{
							ClientID:     authClient.ClientID,
							ClientSecret: authClient.ClientSecret,
							Email:        "invalid",
							Password:     password,
							GrantType:    "password",
						}

						err := form.Save()

						Expect(err.Error()).To(Equal("Email: must be a valid email address."))
					})
				})
			})

			Context("password", func() {
				Context("given NO password", func() {
					It("returns an INVALID password error", func() {
						authClient := FabricateAuthClient()
						user := FabricateUser(faker.Email(), faker.Password())
						form := forms.AuthenticationForm{
							ClientID:     authClient.ClientID,
							ClientSecret: authClient.ClientSecret,
							Email:        user.Email,
							Password:     "",
							GrantType:    "password",
						}

						err := form.Save()

						Expect(err.Error()).To(Equal("Password: cannot be blank."))
					})
				})
			})
		})
	})

	AfterEach(func() {
		CleanupDatabase([]string{"users", "oauth2_clients", "oauth2_tokens"})
	})
})
