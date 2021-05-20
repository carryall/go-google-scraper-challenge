package apiforms_test

import (
	apiforms "go-google-scraper-challenge/forms/api"
	"go-google-scraper-challenge/initializers"
	. "go-google-scraper-challenge/tests/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Forms/API/LoginForm", func() {
	Describe("#Save", func() {
		Context("given login form with valid params", func() {
			It("returns NO error", func() {
				FabricateUser("dev@nimblehq.co", "password")
				oauthClient := FabricateOAuthClient()
				form := apiforms.LoginForm{
					Username:     "dev@nimblehq.co",
					Password:     "password",
					ClientId:     oauthClient.ClientID,
					ClientSecret: oauthClient.ClientSecret,
					GrantType:    "password",
				}

				errors := form.Save()

				Expect(len(errors)).To(BeZero())
			})
		})

		Context("given login form with INVALID params", func() {
			Context("given username is not an email", func() {
				It("returns a username not email error", func() {
					form := apiforms.LoginForm{
						Username:     "non email username",
						Password:     "password",
						ClientId:     "client_id",
						ClientSecret: "client_secret",
						GrantType:    "password",
					}

					errors := form.Save()

					Expect(errors[0].Error()).To(Equal("Username must be a valid email address"))
				})
			})

			Context("given NO username", func() {
				It("returns an invalid username error", func() {
					form := apiforms.LoginForm{
						Username:     "",
						Password:     "password",
						ClientId:     "client_id",
						ClientSecret: "client_secret",
						GrantType:    "password",
					}

					errors := form.Save()

					Expect(errors[0].Error()).To(Equal("Username must be a valid email address"))
				})
			})

			Context("given NO password", func() {
				It("returns an invalid password error", func() {
					form := apiforms.LoginForm{
						Username:     "dev@nimblehq.co",
						Password:     "",
						ClientId:     "client_id",
						ClientSecret: "client_secret",
						GrantType:    "password",
					}

					errors := form.Save()

					Expect(errors[0].Error()).To(Equal("Password can not be empty"))
				})
			})

			Context("given NO client id", func() {
				It("returns an invalid password error", func() {
					form := apiforms.LoginForm{
						Username:     "dev@nimblehq.co",
						Password:     "password",
						ClientId:     "",
						ClientSecret: "client_secret",
						GrantType:    "password",
					}

					errors := form.Save()

					Expect(errors[0].Error()).To(Equal("ClientId can not be empty"))
				})
			})

			Context("given NO client secret", func() {
				It("returns an invalid client secret error", func() {
					form := apiforms.LoginForm{
						Username:     "dev@nimblehq.co",
						Password:     "password",
						ClientId:     "client_id",
						ClientSecret: "",
						GrantType:    "password",
					}

					errors := form.Save()

					Expect(errors[0].Error()).To(Equal("ClientSecret can not be empty"))
				})
			})

			Context("given NO grant type", func() {
				It("returns an invalid grant type error", func() {
					form := apiforms.LoginForm{
						Username:     "dev@nimblehq.co",
						Password:     "password",
						ClientId:     "client_id",
						ClientSecret: "client_secret",
						GrantType:    "",
					}

					errors := form.Save()

					Expect(errors[0].Error()).To(Equal("GrantType can not be empty"))
				})
			})
		})
	})

	AfterEach(func() {
		initializers.CleanupDatabase([]string{"users", "oauth2_clients", "oauth2_tokens"})
	})
})
