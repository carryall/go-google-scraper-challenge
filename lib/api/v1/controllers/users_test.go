package controllers_test

import (
	"net/http"
	"net/url"

	"go-google-scraper-challenge/constants"
	"go-google-scraper-challenge/lib/api/v1/controllers"
	"go-google-scraper-challenge/lib/api/v1/serializers"
	"go-google-scraper-challenge/test"
	. "go-google-scraper-challenge/test"

	"github.com/bxcodec/faker/v3"
	"github.com/google/jsonapi"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UsersController", func() {
	Describe("POST /register", func() {
		Context("given valid params", func() {
			It("responses with status OK", func() {
				authClient := FabricateAuthClient()
				formData := url.Values{
					"username":      {faker.Email()},
					"password":      {faker.Password()},
					"client_id":     {authClient.ClientID},
					"client_secret": {authClient.ClientSecret},
				}

				ctx, response := MakeFormRequest("POST", "/register", formData)

				usersController := controllers.UsersController{}
				usersController.Register(ctx)

				Expect(response.Code).To(Equal(http.StatusOK))
			})

			It("responses with the user information", func() {
				authClient := FabricateAuthClient()
				formData := url.Values{
					"username":      {faker.Email()},
					"password":      {faker.Password()},
					"client_id":     {authClient.ClientID},
					"client_secret": {authClient.ClientSecret},
				}

				ctx, response := MakeFormRequest("POST", "/register", formData)

				usersController := controllers.UsersController{}
				usersController.Register(ctx)

				jsonResponse := serializers.RegistrationJSONResponse{}
				test.GetJSONResponseBody(response.Result(), &jsonResponse)

				Expect(jsonResponse.Data.ID).To(BeNumerically(">", 0))
				Expect(jsonResponse.Data.Attributes.UserID).To(BeNumerically(">", 0))
				Expect(jsonResponse.Data.Attributes.AccessToken).NotTo(Equal(""))
				Expect(jsonResponse.Data.Attributes.RefreshToken).NotTo(Equal(""))
			})
		})

		Context("given INVALID params", func() {
			Context("client ID", func() {
				Context("given NO client ID", func() {
					It("response with the error status", func() {
						authClient := FabricateAuthClient()
						formData := url.Values{
							"username":      {faker.Email()},
							"password":      {faker.Password()},
							"client_id":     {""},
							"client_secret": {authClient.ClientSecret},
						}

						ctx, response := MakeFormRequest("POST", "/register", formData)

						usersController := controllers.UsersController{}
						usersController.Register(ctx)

						Expect(response.Code).To(Equal(http.StatusBadRequest))
					})

					It("response with the error detail", func() {
						authClient := FabricateAuthClient()
						formData := url.Values{
							"username":      {faker.Email()},
							"password":      {faker.Password()},
							"client_id":     {""},
							"client_secret": {authClient.ClientSecret},
						}

						ctx, response := MakeFormRequest("POST", "/register", formData)

						usersController := controllers.UsersController{}
						usersController.Register(ctx)

						jsonResponse := &jsonapi.ErrorsPayload{}
						test.GetJSONResponseBody(response.Result(), &jsonResponse)

						Expect(jsonResponse.Errors[0].Title).To(Equal(constants.Errors[http.StatusBadRequest]))
						Expect(jsonResponse.Errors[0].Code).To(Equal(constants.ERROR_CODE_INVALID_PARAM))
						Expect(jsonResponse.Errors[0].Detail).To(Equal("ClientID: cannot be blank."))
					})
				})

				Context("given INVALID client ID", func() {
					It("response with the error status", func() {
						authClient := FabricateAuthClient()
						formData := url.Values{
							"username":      {faker.Email()},
							"password":      {faker.Password()},
							"client_id":     {"invalid_id"},
							"client_secret": {authClient.ClientSecret},
						}

						ctx, response := MakeFormRequest("POST", "/register", formData)

						usersController := controllers.UsersController{}
						usersController.Register(ctx)

						Expect(response.Code).To(Equal(http.StatusUnauthorized))
					})

					It("response with the error detail", func() {
						authClient := FabricateAuthClient()
						formData := url.Values{
							"username":      {faker.Email()},
							"password":      {faker.Password()},
							"client_id":     {"invalid_id"},
							"client_secret": {authClient.ClientSecret},
						}

						ctx, response := MakeFormRequest("POST", "/register", formData)

						usersController := controllers.UsersController{}
						usersController.Register(ctx)

						jsonResponse := &jsonapi.ErrorsPayload{}
						test.GetJSONResponseBody(response.Result(), &jsonResponse)

						Expect(jsonResponse.Errors[0].Title).To(Equal(constants.Errors[http.StatusUnauthorized]))
						Expect(jsonResponse.Errors[0].Code).To(Equal(constants.ERROR_CODE_INVALID_CREDENTIALS))
						Expect(jsonResponse.Errors[0].Detail).To(Equal(constants.OAuthClientInvalid))
					})
				})
			})

			Context("client secret", func() {
				Context("given NO client secret", func() {
					It("response with the error status", func() {
						authClient := FabricateAuthClient()
						formData := url.Values{
							"username":      {faker.Email()},
							"password":      {faker.Password()},
							"client_id":     {authClient.ClientID},
							"client_secret": {""},
						}

						ctx, response := MakeFormRequest("POST", "/register", formData)

						usersController := controllers.UsersController{}
						usersController.Register(ctx)

						Expect(response.Code).To(Equal(http.StatusBadRequest))
					})

					It("response with the error detail", func() {
						authClient := FabricateAuthClient()
						formData := url.Values{
							"username":      {faker.Email()},
							"password":      {faker.Password()},
							"client_id":     {authClient.ClientID},
							"client_secret": {""},
						}

						ctx, response := MakeFormRequest("POST", "/register", formData)

						usersController := controllers.UsersController{}
						usersController.Register(ctx)

						jsonResponse := &jsonapi.ErrorsPayload{}
						test.GetJSONResponseBody(response.Result(), &jsonResponse)

						Expect(jsonResponse.Errors[0].Title).To(Equal(constants.Errors[http.StatusBadRequest]))
						Expect(jsonResponse.Errors[0].Code).To(Equal(constants.ERROR_CODE_INVALID_PARAM))
						Expect(jsonResponse.Errors[0].Detail).To(Equal("ClientSecret: cannot be blank."))
					})
				})

				Context("given INVALID client secret", func() {
					It("response with the error status", func() {
						authClient := FabricateAuthClient()
						formData := url.Values{
							"username":      {faker.Email()},
							"password":      {faker.Password()},
							"client_id":     {authClient.ClientID},
							"client_secret": {"invalid secret"},
						}

						ctx, response := MakeFormRequest("POST", "/register", formData)

						usersController := controllers.UsersController{}
						usersController.Register(ctx)

						Expect(response.Code).To(Equal(http.StatusUnauthorized))
					})

					It("response with the error detail", func() {
						authClient := FabricateAuthClient()
						formData := url.Values{
							"username":      {faker.Email()},
							"password":      {faker.Password()},
							"client_id":     {authClient.ClientID},
							"client_secret": {"invalid secret"},
						}

						ctx, response := MakeFormRequest("POST", "/register", formData)

						usersController := controllers.UsersController{}
						usersController.Register(ctx)

						jsonResponse := &jsonapi.ErrorsPayload{}
						test.GetJSONResponseBody(response.Result(), &jsonResponse)

						Expect(jsonResponse.Errors[0].Title).To(Equal(constants.Errors[http.StatusUnauthorized]))
						Expect(jsonResponse.Errors[0].Code).To(Equal(constants.ERROR_CODE_INVALID_CREDENTIALS))
						Expect(jsonResponse.Errors[0].Detail).To(Equal(constants.OAuthClientInvalid))
					})
				})
			})

			Context("email", func() {
				Context("given NO email", func() {
					It("responses with the error status", func() {
						authClient := FabricateAuthClient()
						formData := url.Values{
							"username":      {""},
							"password":      {faker.Password()},
							"client_id":     {authClient.ClientID},
							"client_secret": {authClient.ClientSecret},
						}

						ctx, response := MakeFormRequest("POST", "/register", formData)

						usersController := controllers.UsersController{}
						usersController.Register(ctx)

						Expect(response.Code).To(Equal(http.StatusBadRequest))
					})

					It("responses with the error detail", func() {
						authClient := FabricateAuthClient()
						formData := url.Values{
							"username":      {""},
							"password":      {faker.Password()},
							"client_id":     {authClient.ClientID},
							"client_secret": {authClient.ClientSecret},
						}

						ctx, response := MakeFormRequest("POST", "/register", formData)

						usersController := controllers.UsersController{}
						usersController.Register(ctx)

						jsonResponse := &jsonapi.ErrorsPayload{}
						test.GetJSONResponseBody(response.Result(), &jsonResponse)

						Expect(jsonResponse.Errors[0].Title).To(Equal(constants.Errors[http.StatusBadRequest]))
						Expect(jsonResponse.Errors[0].Code).To(Equal(constants.ERROR_CODE_INVALID_PARAM))
						Expect(jsonResponse.Errors[0].Detail).To(Equal("Email: cannot be blank."))
					})
				})

				Context("given an INVALID email", func() {
					It("responses with the error status", func() {
						authClient := FabricateAuthClient()
						formData := url.Values{
							"username":      {"invalid email"},
							"password":      {faker.Password()},
							"client_id":     {authClient.ClientID},
							"client_secret": {authClient.ClientSecret},
						}

						ctx, response := MakeFormRequest("POST", "/register", formData)

						usersController := controllers.UsersController{}
						usersController.Register(ctx)

						Expect(response.Code).To(Equal(http.StatusBadRequest))
					})

					It("responses with the error detail", func() {
						authClient := FabricateAuthClient()
						formData := url.Values{
							"username":      {"invalid email"},
							"password":      {faker.Password()},
							"client_id":     {authClient.ClientID},
							"client_secret": {authClient.ClientSecret},
						}

						ctx, response := MakeFormRequest("POST", "/register", formData)

						usersController := controllers.UsersController{}
						usersController.Register(ctx)

						jsonResponse := &jsonapi.ErrorsPayload{}
						test.GetJSONResponseBody(response.Result(), &jsonResponse)

						Expect(jsonResponse.Errors[0].Title).To(Equal(constants.Errors[http.StatusBadRequest]))
						Expect(jsonResponse.Errors[0].Code).To(Equal(constants.ERROR_CODE_INVALID_PARAM))
						Expect(jsonResponse.Errors[0].Detail).To(Equal("Email: must be a valid email address."))
					})
				})
			})

			Context("password", func() {
				Context("given NO password", func() {
					It("responses with the error status", func() {
						authClient := FabricateAuthClient()
						formData := url.Values{
							"username":      {faker.Email()},
							"password":      {""},
							"client_id":     {authClient.ClientID},
							"client_secret": {authClient.ClientSecret},
						}

						ctx, response := MakeFormRequest("POST", "/register", formData)

						usersController := controllers.UsersController{}
						usersController.Register(ctx)

						Expect(response.Code).To(Equal(http.StatusBadRequest))
					})

					It("responses with the error detail", func() {
						authClient := FabricateAuthClient()
						formData := url.Values{
							"username":      {faker.Email()},
							"password":      {""},
							"client_id":     {authClient.ClientID},
							"client_secret": {authClient.ClientSecret},
						}

						ctx, response := MakeFormRequest("POST", "/register", formData)

						usersController := controllers.UsersController{}
						usersController.Register(ctx)

						jsonResponse := &jsonapi.ErrorsPayload{}
						test.GetJSONResponseBody(response.Result(), &jsonResponse)

						Expect(jsonResponse.Errors[0].Title).To(Equal(constants.Errors[http.StatusBadRequest]))
						Expect(jsonResponse.Errors[0].Code).To(Equal(constants.ERROR_CODE_INVALID_PARAM))
						Expect(jsonResponse.Errors[0].Detail).To(Equal("Password: cannot be blank."))
					})
				})

				Context("given password is shorter than 6 charactors", func() {
					It("responses with the error status", func() {
						authClient := FabricateAuthClient()
						formData := url.Values{
							"username":      {faker.Email()},
							"password":      {"short"},
							"client_id":     {authClient.ClientID},
							"client_secret": {authClient.ClientSecret},
						}

						ctx, response := MakeFormRequest("POST", "/register", formData)

						usersController := controllers.UsersController{}
						usersController.Register(ctx)

						Expect(response.Code).To(Equal(http.StatusBadRequest))
					})

					It("responses with the error detail", func() {
						authClient := FabricateAuthClient()
						formData := url.Values{
							"username":      {faker.Email()},
							"password":      {"short"},
							"client_id":     {authClient.ClientID},
							"client_secret": {authClient.ClientSecret},
						}

						ctx, response := MakeFormRequest("POST", "/register", formData)

						usersController := controllers.UsersController{}
						usersController.Register(ctx)

						jsonResponse := &jsonapi.ErrorsPayload{}
						test.GetJSONResponseBody(response.Result(), &jsonResponse)

						Expect(jsonResponse.Errors[0].Title).To(Equal(constants.Errors[http.StatusBadRequest]))
						Expect(jsonResponse.Errors[0].Code).To(Equal(constants.ERROR_CODE_INVALID_PARAM))
						Expect(jsonResponse.Errors[0].Detail).To(Equal("Password: the length must be between 6 and 50."))
					})
				})
			})
		})
	})

	AfterEach(func() {
		CleanupDatabase([]string{"users", "oauth2_clients", "oauth2_tokens"})
	})
})
