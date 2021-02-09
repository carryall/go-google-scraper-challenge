package apicontrollers_test

import (
	"net/http"

	"go-google-scraper-challenge/initializers"
	. "go-google-scraper-challenge/test/helpers"
	. "go-google-scraper-challenge/test/serializers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AuthController", func() {
	Describe("POST /login", func() {
		Context("given valid params", func() {
			It("returns with status ok", func() {
				FabricateUser("dev@nimblehq.co", "password")
				oauth_client := FabricateOAuthClient()
				body := GenerateRequestBody(map[string]string{
					"client_id":     oauth_client.ClientID,
					"client_secret": oauth_client.ClientSecret,
					"username":      "dev@nimblehq.co",
					"password":      "password",
					"grant_type":    "password",
				})
				response := MakeRequest("POST", "/api/v1/login", body)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
			})

			It("returns with a valid token", func() {
				FabricateUser("dev@nimblehq.co", "password")
				oauth_client := FabricateOAuthClient()
				body := GenerateRequestBody(map[string]string{
					"client_id":     oauth_client.ClientID,
					"client_secret": oauth_client.ClientSecret,
					"username":      "dev@nimblehq.co",
					"password":      "password",
					"grant_type":    "password",
				})
				response := MakeRequest("POST", "/api/v1/login", body)
				responseBody := LoginResponse{}
				GetJSONResponseBody(response, &responseBody)

				Expect(responseBody.AccessToken).NotTo(BeNil())
				Expect(responseBody.RefreshToken).NotTo(BeNil())
				Expect(responseBody.ExpiresIn).NotTo(BeNil())
				Expect(responseBody.TokenType).NotTo(BeNil())
			})
		})

		Context("given missing params", func() {
			Context("given NO client id", func() {
				It("returns with a bad request error", func() {
					FabricateUser("dev@nimblehq.co", "password")
					oauth_client := FabricateOAuthClient()
					body := GenerateRequestBody(map[string]string{
						"client_id":     "",
						"client_secret": oauth_client.ClientSecret,
						"username":      "dev@nimblehq.co",
						"password":      "password",
						"grant_type":    "password",
					})
					response := MakeRequest("POST", "/api/v1/login", body)
					responseBody := ErrorResponse{}
					GetJSONResponseBody(response, &responseBody)

					Expect(response.StatusCode).To(Equal(http.StatusBadRequest))
					Expect(responseBody.Error).To(Equal("Bad Request"))
					Expect(responseBody.ErrorDescription).To(Equal("The request is missing a required parameter, includes an invalid parameter value, includes a parameter more than once, or is otherwise malformed"))
				})
			})

			Context("given NO client secret", func() {
				It("returns with a bad request error", func() {
					FabricateUser("dev@nimblehq.co", "password")
					oauth_client := FabricateOAuthClient()
					body := GenerateRequestBody(map[string]string{
						"client_id":     oauth_client.ClientID,
						"client_secret": "",
						"username":      "dev@nimblehq.co",
						"password":      "password",
						"grant_type":    "password",
					})
					response := MakeRequest("POST", "/api/v1/login", body)
					responseBody := ErrorResponse{}
					GetJSONResponseBody(response, &responseBody)

					Expect(response.StatusCode).To(Equal(http.StatusBadRequest))
					Expect(responseBody.Error).To(Equal("Bad Request"))
					Expect(responseBody.ErrorDescription).To(Equal("The request is missing a required parameter, includes an invalid parameter value, includes a parameter more than once, or is otherwise malformed"))
				})
			})

			Context("given NO grant type", func() {
				It("returns with a bad request error", func() {
					FabricateUser("dev@nimblehq.co", "password")
					oauth_client := FabricateOAuthClient()
					body := GenerateRequestBody(map[string]string{
						"client_id":     oauth_client.ClientID,
						"client_secret": oauth_client.ClientSecret,
						"username":      "dev@nimblehq.co",
						"password":      "password",
						"grant_type":    "",
					})
					response := MakeRequest("POST", "/api/v1/login", body)
					responseBody := ErrorResponse{}
					GetJSONResponseBody(response, &responseBody)

					Expect(response.StatusCode).To(Equal(http.StatusBadRequest))
					Expect(responseBody.Error).To(Equal("Bad Request"))
					Expect(responseBody.ErrorDescription).To(Equal("The request is missing a required parameter, includes an invalid parameter value, includes a parameter more than once, or is otherwise malformed"))
				})
			})

			Context("given NO username", func() {
				It("returns with a bad request error", func() {
					FabricateUser("dev@nimblehq.co", "password")
					oauth_client := FabricateOAuthClient()
					body := GenerateRequestBody(map[string]string{
						"client_id":     oauth_client.ClientID,
						"client_secret": oauth_client.ClientSecret,
						"username":      "",
						"password":      "password",
						"grant_type":    "password",
					})
					response := MakeRequest("POST", "/api/v1/login", body)
					responseBody := ErrorResponse{}
					GetJSONResponseBody(response, &responseBody)

					Expect(response.StatusCode).To(Equal(http.StatusBadRequest))
					Expect(responseBody.Error).To(Equal("Bad Request"))
					Expect(responseBody.ErrorDescription).To(Equal("The request is missing a required parameter, includes an invalid parameter value, includes a parameter more than once, or is otherwise malformed"))
				})
			})

			Context("given NO password", func() {
				It("returns with a bad request error", func() {
					FabricateUser("dev@nimblehq.co", "password")
					oauth_client := FabricateOAuthClient()
					body := GenerateRequestBody(map[string]string{
						"client_id":     oauth_client.ClientID,
						"client_secret": oauth_client.ClientSecret,
						"username":      "dev@nimblehq.co",
						"password":      "",
						"grant_type":    "password",
					})
					response := MakeRequest("POST", "/api/v1/login", body)
					responseBody := ErrorResponse{}
					GetJSONResponseBody(response, &responseBody)

					Expect(response.StatusCode).To(Equal(http.StatusBadRequest))
					Expect(responseBody.Error).To(Equal("Bad Request"))
					Expect(responseBody.ErrorDescription).To(Equal("The request is missing a required parameter, includes an invalid parameter value, includes a parameter more than once, or is otherwise malformed"))
				})
			})
		})

		Context("given INVALID params", func() {
			Context("given INVALID client id", func() {
				It("returns with an invalid client error", func() {
					FabricateUser("dev@nimblehq.co", "password")
					oauth_client := FabricateOAuthClient()
					body := GenerateRequestBody(map[string]string{
						"client_id":     "invalid client id",
						"client_secret": oauth_client.ClientSecret,
						"username":      "dev@nimblehq.co",
						"password":      "password",
						"grant_type":    "password",
					})
					response := MakeRequest("POST", "/api/v1/login", body)
					responseBody := ErrorResponse{}
					GetJSONResponseBody(response, &responseBody)

					Expect(response.StatusCode).To(Equal(http.StatusUnauthorized))
					Expect(responseBody.Error).To(Equal("invalid_client"))
					Expect(responseBody.ErrorDescription).To(Equal("Client authentication failed"))
				})
			})

			Context("given INVALID client secret", func() {
				It("returns with an invalid client error", func() {
					FabricateUser("dev@nimblehq.co", "password")
					oauth_client := FabricateOAuthClient()
					body := GenerateRequestBody(map[string]string{
						"client_id":     oauth_client.ClientID,
						"client_secret": "invalid client secret",
						"username":      "dev@nimblehq.co",
						"password":      "password",
						"grant_type":    "password",
					})
					response := MakeRequest("POST", "/api/v1/login", body)
					responseBody := ErrorResponse{}
					GetJSONResponseBody(response, &responseBody)

					Expect(response.StatusCode).To(Equal(http.StatusUnauthorized))
					Expect(responseBody.Error).To(Equal("invalid_client"))
					Expect(responseBody.ErrorDescription).To(Equal("Client authentication failed"))
				})
			})

			Context("given INVALID grant type", func() {
				It("returns with an unsupport grant type error", func() {
					FabricateUser("dev@nimblehq.co", "password")
					oauth_client := FabricateOAuthClient()
					body := GenerateRequestBody(map[string]string{
						"client_id":     oauth_client.ClientID,
						"client_secret": oauth_client.ClientSecret,
						"username":      "dev@nimblehq.co",
						"password":      "password",
						"grant_type":    "invalid grant type",
					})
					response := MakeRequest("POST", "/api/v1/login", body)
					responseBody := ErrorResponse{}
					GetJSONResponseBody(response, &responseBody)

					Expect(response.StatusCode).To(Equal(http.StatusUnauthorized))
					Expect(responseBody.Error).To(Equal("unsupported_grant_type"))
					Expect(responseBody.ErrorDescription).To(Equal("The authorization grant type is not supported by the authorization server"))
				})
			})

			Context("given INVALID username", func() {
				It("returns with an invalid client error", func() {
					FabricateUser("dev@nimblehq.co", "password")
					oauth_client := FabricateOAuthClient()
					body := GenerateRequestBody(map[string]string{
						"client_id":     oauth_client.ClientID,
						"client_secret": oauth_client.ClientSecret,
						"username":      "invalid@email.com",
						"password":      "password",
						"grant_type":    "password",
					})
					response := MakeRequest("POST", "/api/v1/login", body)
					responseBody := ErrorResponse{}
					GetJSONResponseBody(response, &responseBody)

					Expect(response.StatusCode).To(Equal(http.StatusUnauthorized))
					Expect(responseBody.Error).To(Equal("invalid_client"))
					Expect(responseBody.ErrorDescription).To(Equal("Client authentication failed"))
				})
			})

			Context("given INVALID password", func() {
				It("returns with an invalid client error", func() {
					FabricateUser("dev@nimblehq.co", "password")
					oauth_client := FabricateOAuthClient()
					body := GenerateRequestBody(map[string]string{
						"client_id":     oauth_client.ClientID,
						"client_secret": oauth_client.ClientSecret,
						"username":      "dev@nimblehq.co",
						"password":      "invalid password",
						"grant_type":    "password",
					})
					response := MakeRequest("POST", "/api/v1/login", body)
					responseBody := ErrorResponse{}
					GetJSONResponseBody(response, &responseBody)

					Expect(response.StatusCode).To(Equal(http.StatusUnauthorized))
					Expect(responseBody.Error).To(Equal("invalid_client"))
					Expect(responseBody.ErrorDescription).To(Equal("Client authentication failed"))
				})
			})
		})
	})

	AfterEach(func() {
		initializers.CleanupDatabase("user")
		initializers.CleanupDatabase("oauth2_clients")
		initializers.CleanupDatabase("oauth2_tokens")
	})
})
