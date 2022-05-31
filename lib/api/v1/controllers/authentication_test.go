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

var _ = Describe("AuthenticationController", func() {
	Describe("GET /login", func() {
		Context("given a valid credentials", func() {
			It("responses with status OK", func() {
				authClient := FabricateAuthClient()
				password := faker.Password()
				user := FabricateUser(faker.Email(), password)
				formData := url.Values{
					"username":      {user.Email},
					"password":      {password},
					"client_id":     {authClient.ClientID},
					"client_secret": {authClient.ClientSecret},
					"grant_type":    {"password"},
				}

				ctx, response := MakeFormRequest("POST", "/login", formData)

				authenticationController := controllers.AuthenticationController{}
				authenticationController.Login(ctx)

				Expect(response.Code).To(Equal(http.StatusOK))
			})

			It("responses with authentication tokens", func() {
				authClient := FabricateAuthClient()
				password := faker.Password()
				user := FabricateUser(faker.Email(), password)
				formData := url.Values{
					"username":      {user.Email},
					"password":      {password},
					"client_id":     {authClient.ClientID},
					"client_secret": {authClient.ClientSecret},
					"grant_type":    {"password"},
				}

				ctx, response := MakeFormRequest("POST", "/login", formData)

				authenticationController := controllers.AuthenticationController{}
				authenticationController.Login(ctx)

				jsonResponse := serializers.AuthenticationJSONResponse{}
				test.GetJSONResponseBody(response.Result(), &jsonResponse)

				Expect(jsonResponse.Data.ID).NotTo(Equal(""))
				Expect(jsonResponse.Data.Attributes.AccessToken).NotTo(Equal(""))
				Expect(jsonResponse.Data.Attributes.RefreshToken).NotTo(Equal(""))
				Expect(jsonResponse.Data.Attributes.ExpiresIn).To(BeNumerically(">", 0))
				Expect(jsonResponse.Data.Attributes.TokenType).NotTo(Equal(""))
			})
		})

		Context("given missing params", func() {
			Context("given NO grant type", func() {
				It("responses with a bad request error", func() {
					authClient := FabricateAuthClient()
					password := faker.Password()
					user := FabricateUser(faker.Email(), password)
					formData := url.Values{
						"username":      {user.Email},
						"password":      {"wrong password"},
						"client_id":     {authClient.ClientID},
						"client_secret": {authClient.ClientSecret},
					}

					ctx, response := MakeFormRequest("POST", "/login", formData)

					authenticationController := controllers.AuthenticationController{}
					authenticationController.Login(ctx)

					jsonResponse := &jsonapi.ErrorsPayload{}
					test.GetJSONResponseBody(response.Result(), &jsonResponse)

					Expect(jsonResponse.Errors[0].Title).To(Equal(constants.Errors[http.StatusBadRequest]))
					Expect(jsonResponse.Errors[0].Code).To(Equal(constants.ERROR_CODE_INVALID_PARAM))
					Expect(jsonResponse.Errors[0].Detail).To(Equal("GrantType: cannot be blank."))
				})
			})

			Context("given NO client ID", func() {
				It("responses with a bad request error", func() {
					authClient := FabricateAuthClient()
					password := faker.Password()
					user := FabricateUser(faker.Email(), password)
					formData := url.Values{
						"username":      {user.Email},
						"password":      {"wrong password"},
						"grant_type":    {"password"},
						"client_secret": {authClient.ClientSecret},
					}

					ctx, response := MakeFormRequest("POST", "/login", formData)

					authenticationController := controllers.AuthenticationController{}
					authenticationController.Login(ctx)

					jsonResponse := &jsonapi.ErrorsPayload{}
					test.GetJSONResponseBody(response.Result(), &jsonResponse)

					Expect(jsonResponse.Errors[0].Title).To(Equal(constants.Errors[http.StatusBadRequest]))
					Expect(jsonResponse.Errors[0].Code).To(Equal(constants.ERROR_CODE_INVALID_PARAM))
					Expect(jsonResponse.Errors[0].Detail).To(Equal("ClientID: cannot be blank."))
				})
			})

			Context("given NO client secret", func() {
				It("responses with a bad request error", func() {
					authClient := FabricateAuthClient()
					password := faker.Password()
					user := FabricateUser(faker.Email(), password)
					formData := url.Values{
						"username":   {user.Email},
						"password":   {"wrong password"},
						"grant_type": {"password"},
						"client_id":  {authClient.ClientID},
					}

					ctx, response := MakeFormRequest("POST", "/login", formData)

					authenticationController := controllers.AuthenticationController{}
					authenticationController.Login(ctx)

					jsonResponse := &jsonapi.ErrorsPayload{}
					test.GetJSONResponseBody(response.Result(), &jsonResponse)

					Expect(jsonResponse.Errors[0].Title).To(Equal(constants.Errors[http.StatusBadRequest]))
					Expect(jsonResponse.Errors[0].Code).To(Equal(constants.ERROR_CODE_INVALID_PARAM))
					Expect(jsonResponse.Errors[0].Detail).To(Equal("ClientSecret: cannot be blank."))
				})
			})

			Context("given NO username", func() {
				It("responses with a bad request error", func() {
					authClient := FabricateAuthClient()
					password := faker.Password()
					FabricateUser(faker.Email(), password)
					formData := url.Values{
						"password":      {password},
						"client_id":     {authClient.ClientID},
						"client_secret": {authClient.ClientSecret},
						"grant_type":    {"password"},
					}

					ctx, response := MakeFormRequest("POST", "/login", formData)

					authenticationController := controllers.AuthenticationController{}
					authenticationController.Login(ctx)

					jsonResponse := &jsonapi.ErrorsPayload{}
					test.GetJSONResponseBody(response.Result(), &jsonResponse)

					Expect(jsonResponse.Errors[0].Title).To(Equal(constants.Errors[http.StatusBadRequest]))
					Expect(jsonResponse.Errors[0].Code).To(Equal(constants.ERROR_CODE_INVALID_PARAM))
					Expect(jsonResponse.Errors[0].Detail).To(Equal("Email: cannot be blank."))
				})
			})

			Context("given NO password", func() {
				It("responses with a bad request error", func() {
					authClient := FabricateAuthClient()
					password := faker.Password()
					user := FabricateUser(faker.Email(), password)
					formData := url.Values{
						"username":      {user.Email},
						"client_id":     {authClient.ClientID},
						"client_secret": {authClient.ClientSecret},
						"grant_type":    {"password"},
					}

					ctx, response := MakeFormRequest("POST", "/login", formData)

					authenticationController := controllers.AuthenticationController{}
					authenticationController.Login(ctx)

					jsonResponse := &jsonapi.ErrorsPayload{}
					test.GetJSONResponseBody(response.Result(), &jsonResponse)

					Expect(jsonResponse.Errors[0].Title).To(Equal(constants.Errors[http.StatusBadRequest]))
					Expect(jsonResponse.Errors[0].Code).To(Equal(constants.ERROR_CODE_INVALID_PARAM))
					Expect(jsonResponse.Errors[0].Detail).To(Equal("Password: cannot be blank."))
				})
			})
		})

		Context("given an invalid credentials", func() {
			Context("given an INVALID client ID", func() {
				It("responses with an unauthorized error", func() {
					authClient := FabricateAuthClient()
					password := faker.Password()
					user := FabricateUser(faker.Email(), password)
					formData := url.Values{
						"username":      {user.Email},
						"password":      {password},
						"client_id":     {"invalid client ID"},
						"client_secret": {authClient.ClientSecret},
						"grant_type":    {"password"},
					}

					ctx, response := MakeFormRequest("POST", "/login", formData)

					authenticationController := controllers.AuthenticationController{}
					authenticationController.Login(ctx)

					responseBody := jsonapi.ErrorsPayload{}
					GetJSONResponseBody(response.Result(), &responseBody)

					Expect(response.Code).To(Equal(http.StatusUnauthorized))
					Expect(responseBody.Errors[0].Title).To(Equal(constants.Errors[http.StatusUnauthorized]))
					Expect(responseBody.Errors[0].Code).To(Equal(constants.ERROR_CODE_INVALID_CREDENTIALS))
					Expect(responseBody.Errors[0].Detail).To(Equal(constants.OAuthClientInvalid))
				})
			})

			Context("given an INVALID client secret", func() {
				It("responses with an unauthorized error", func() {
					authClient := FabricateAuthClient()
					password := faker.Password()
					user := FabricateUser(faker.Email(), password)
					formData := url.Values{
						"username":      {user.Email},
						"password":      {password},
						"client_id":     {authClient.ClientID},
						"client_secret": {"invalid client secret"},
						"grant_type":    {"password"},
					}

					ctx, response := MakeFormRequest("POST", "/login", formData)

					authenticationController := controllers.AuthenticationController{}
					authenticationController.Login(ctx)

					responseBody := jsonapi.ErrorsPayload{}
					GetJSONResponseBody(response.Result(), &responseBody)

					Expect(response.Code).To(Equal(http.StatusUnauthorized))
					Expect(responseBody.Errors[0].Title).To(Equal(constants.Errors[http.StatusUnauthorized]))
					Expect(responseBody.Errors[0].Code).To(Equal(constants.ERROR_CODE_INVALID_CREDENTIALS))
					Expect(responseBody.Errors[0].Detail).To(Equal(constants.OAuthClientInvalid))
				})
			})

			Context("given an email that does NOT belongs to any user", func() {
				It("responses with an unauthorized error", func() {
					authClient := FabricateAuthClient()
					password := faker.Password()
					FabricateUser(faker.Email(), password)
					formData := url.Values{
						"username":      {faker.Email()},
						"password":      {password},
						"client_id":     {authClient.ClientID},
						"client_secret": {authClient.ClientSecret},
						"grant_type":    {"password"},
					}

					ctx, response := MakeFormRequest("POST", "/login", formData)

					authenticationController := controllers.AuthenticationController{}
					authenticationController.Login(ctx)

					responseBody := jsonapi.ErrorsPayload{}
					GetJSONResponseBody(response.Result(), &responseBody)

					Expect(response.Code).To(Equal(http.StatusUnauthorized))
					Expect(responseBody.Errors[0].Title).To(Equal(constants.Errors[http.StatusUnauthorized]))
					Expect(responseBody.Errors[0].Code).To(Equal(constants.ERROR_CODE_INVALID_CREDENTIALS))
					Expect(responseBody.Errors[0].Detail).To(Equal(constants.UserDoesNotExist))
				})
			})

			Context("given a wrong password", func() {
				It("responses with an unauthorized error", func() {
					authClient := FabricateAuthClient()
					user := FabricateUser(faker.Email(), faker.Password())
					formData := url.Values{
						"username":      {user.Email},
						"password":      {"wrong password"},
						"client_id":     {authClient.ClientID},
						"client_secret": {authClient.ClientSecret},
						"grant_type":    {"password"},
					}

					ctx, response := MakeFormRequest("POST", "/login", formData)

					authenticationController := controllers.AuthenticationController{}
					authenticationController.Login(ctx)

					responseBody := jsonapi.ErrorsPayload{}
					GetJSONResponseBody(response.Result(), &responseBody)

					Expect(response.Code).To(Equal(http.StatusUnauthorized))
					Expect(responseBody.Errors[0].Title).To(Equal(constants.Errors[http.StatusUnauthorized]))
					Expect(responseBody.Errors[0].Code).To(Equal(constants.ERROR_CODE_INVALID_CREDENTIALS))
					Expect(responseBody.Errors[0].Detail).To(Equal(constants.OAuthClientInvalid))
				})
			})

			Context("given an INVALID grant type", func() {
				It("responses with an unauthorized error", func() {
					authClient := FabricateAuthClient()
					password := faker.Password()
					user := FabricateUser(faker.Email(), password)
					formData := url.Values{
						"username":      {user.Email},
						"password":      {"wrong password"},
						"client_id":     {authClient.ClientID},
						"client_secret": {authClient.ClientSecret},
						"grant_type":    {"invalid grant type"},
					}

					ctx, response := MakeFormRequest("POST", "/login", formData)

					authenticationController := controllers.AuthenticationController{}
					authenticationController.Login(ctx)

					responseBody := jsonapi.ErrorsPayload{}
					GetJSONResponseBody(response.Result(), &responseBody)

					Expect(response.Code).To(Equal(http.StatusUnauthorized))
					Expect(responseBody.Errors[0].Title).To(Equal(constants.Errors[http.StatusUnauthorized]))
					Expect(responseBody.Errors[0].Code).To(Equal(constants.ERROR_CODE_INVALID_CREDENTIALS))
					Expect(responseBody.Errors[0].Detail).To(Equal(constants.OAuthClientInvalid))
				})
			})
		})
	})

	AfterEach(func() {
		CleanupDatabase([]string{"users", "oauth2_clients", "oauth2_tokens"})
	})
})
