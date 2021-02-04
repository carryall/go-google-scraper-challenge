package api_controllers_test

import (
	. "go-google-scraper-challenge/helpers/test"
	"go-google-scraper-challenge/initializers"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
}

type ErrorResponse struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

var _ = Describe("AuthController", func() {
	Describe("POST /login", func() {
		Context("given valid params", func() {
			It("responses with status ok", func() {
				FabricateUser("dev@nimblehq.co", "password")
				oauth_client := FabricateOAuthClient()
				body := RequestBody(map[string]string{
					"client_id":     oauth_client.ClientID,
					"client_secret": oauth_client.ClientSecret,
					"username":      "dev@nimblehq.co",
					"password":      "password",
					"grant_type":    "password",
				})
				response := MakeRequest("POST", "/api/v1/login", body)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
			})

			It("response with token data", func() {
				FabricateUser("dev@nimblehq.co", "password")
				oauth_client := FabricateOAuthClient()
				body := RequestBody(map[string]string{
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
				It("response with invalid request error", func() {
					FabricateUser("dev@nimblehq.co", "password")
					oauth_client := FabricateOAuthClient()
					body := RequestBody(map[string]string{
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
				It("response with invalid request error", func() {
					FabricateUser("dev@nimblehq.co", "password")
					oauth_client := FabricateOAuthClient()
					body := RequestBody(map[string]string{
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
				It("response with invalid request error", func() {
					FabricateUser("dev@nimblehq.co", "password")
					oauth_client := FabricateOAuthClient()
					body := RequestBody(map[string]string{
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
				It("response with invalid request error", func() {
					FabricateUser("dev@nimblehq.co", "password")
					oauth_client := FabricateOAuthClient()
					body := RequestBody(map[string]string{
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
				It("response with invalid request error", func() {
					FabricateUser("dev@nimblehq.co", "password")
					oauth_client := FabricateOAuthClient()
					body := RequestBody(map[string]string{
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
				It("response with invalid client error", func() {
					FabricateUser("dev@nimblehq.co", "password")
					oauth_client := FabricateOAuthClient()
					body := RequestBody(map[string]string{
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
				It("response with invalid client error", func() {
					FabricateUser("dev@nimblehq.co", "password")
					oauth_client := FabricateOAuthClient()
					body := RequestBody(map[string]string{
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
				It("response with unsupport grant type error", func() {
					FabricateUser("dev@nimblehq.co", "password")
					oauth_client := FabricateOAuthClient()
					body := RequestBody(map[string]string{
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
				It("response with invalid client error", func() {
					FabricateUser("dev@nimblehq.co", "password")
					oauth_client := FabricateOAuthClient()
					body := RequestBody(map[string]string{
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
				It("response with invalid client error", func() {
					FabricateUser("dev@nimblehq.co", "password")
					oauth_client := FabricateOAuthClient()
					body := RequestBody(map[string]string{
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
