package api_controllers_test

import (
	. "go-google-scraper-challenge/helpers/test"
	"go-google-scraper-challenge/initializers"
	"net/http"
	"net/url"
	"strings"

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
	Error            string `json: "error"`
	ErrorDescription string `json: "error_description"`
}

var _ = Describe("AuthController", func() {
	Describe("POST /login", func() {
		Context("given valid params", func() {
			It("responses with status ok", func() {
				FabricateUser("dev@nimblehq.co", "password")
				oauth_client := FabricateOAuthClient()
				data := url.Values{}
				data.Set("client_id", oauth_client.ClientID)
				data.Set("client_secret", oauth_client.ClientSecret)
				data.Set("username", "dev@nimblehq.co")
				data.Set("password", "password")
				data.Set("grant_type", "password")
				body := strings.NewReader(data.Encode())
				response := MakeRequest("POST", "/api/v1/login", body)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
			})

			It("response with token data", func() {
				FabricateUser("dev@nimblehq.co", "password")
				oauth_client := FabricateOAuthClient()
				body := MakeRequestBody(map[string]string{
					"client_id":     oauth_client.ClientID,
					"client_secret": oauth_client.ClientSecret,
					"username":      "dev@nimblehq.co",
					"password":      "password",
					"grant_type":    "password",
				})
				response := MakeRequest("POST", "/api/v1/login", body)
				responseBody := LoginResponse{}
				GetJSONResponseBody(response, responseBody)

				Expect(responseBody.AccessToken).NotTo(BeNil())
				Expect(responseBody.RefreshToken).NotTo(BeNil())
				Expect(responseBody.ExpiresIn).NotTo(BeNil())
				Expect(responseBody.TokenType).NotTo(BeNil())
			})
		})

		Context("given INVALID params", func() {
			Context("given INVALID client id", func() {
				It("response with error", func() {
					FabricateUser("dev@nimblehq.co", "password")
					oauth_client := FabricateOAuthClient()
					body := MakeRequestBody(map[string]string{
						"client_id":     "invalid client id",
						"client_secret": oauth_client.ClientSecret,
						"username":      "dev@nimblehq.co",
						"password":      "password",
						"grant_type":    "password",
					})
					response := MakeRequest("POST", "/api/v1/login", body)
					responseBody := ErrorResponse{}
					GetJSONResponseBody(response, responseBody)

					Expect(response.StatusCode).To(Equal(http.StatusUnauthorized))
					Expect(responseBody.Error).To(Equal("unauthorized_client"))
					Expect(responseBody.ErrorDescription).To(Equal("The client is not authorized to request an authorization code using this method"))
				})
			})

			Context("given INVALID client secret", func() {
				It("response with error", func() {
					FabricateUser("dev@nimblehq.co", "password")
					oauth_client := FabricateOAuthClient()
					body := MakeRequestBody(map[string]string{
						"client_id":     oauth_client.ClientID,
						"client_secret": "invalid client secret",
						"username":      "dev@nimblehq.co",
						"password":      "password",
						"grant_type":    "password",
					})
					response := MakeRequest("POST", "/api/v1/login", body)
					responseBody := ErrorResponse{}
					GetJSONResponseBody(response, responseBody)

					Expect(response.StatusCode).To(Equal(http.StatusUnauthorized))
					Expect(responseBody.Error).To(Equal("unauthorized_client"))
					Expect(responseBody.ErrorDescription).To(Equal("The client is not authorized to request an authorization code using this method"))
				})
			})

			Context("given INVALID grant type", func() {
				It("response with error", func() {
					FabricateUser("dev@nimblehq.co", "password")
					oauth_client := FabricateOAuthClient()
					body := MakeRequestBody(map[string]string{
						"client_id":     oauth_client.ClientID,
						"client_secret": oauth_client.ClientSecret,
						"username":      "dev@nimblehq.co",
						"password":      "password",
						"grant_type":    "invalid grant type",
					})
					response := MakeRequest("POST", "/api/v1/login", body)
					responseBody := ErrorResponse{}
					GetJSONResponseBody(response, responseBody)

					Expect(response.StatusCode).To(Equal(http.StatusUnauthorized))
					Expect(responseBody.Error).To(Equal("invalid_grant"))
					Expect(responseBody.ErrorDescription).To(Equal("The provided authorization grant (e.g., authorization code, resource owner credentials) or refresh token is invalid, expired, revoked, does not match the redirection URI used in the authorization request, or was issued to another client"))
				})
			})

			Context("given INVALID username", func() {
				It("response with error", func() {
					FabricateUser("dev@nimblehq.co", "password")
					oauth_client := FabricateOAuthClient()
					body := MakeRequestBody(map[string]string{
						"client_id":     oauth_client.ClientID,
						"client_secret": oauth_client.ClientSecret,
						"username":      "invalid@email.com",
						"password":      "password",
						"grant_type":    "password",
					})
					response := MakeRequest("POST", "/api/v1/login", body)
					responseBody := ErrorResponse{}
					GetJSONResponseBody(response, responseBody)

					Expect(response.StatusCode).To(Equal(http.StatusUnauthorized))
					Expect(responseBody.Error).To(Equal("unauthorized_client"))
					Expect(responseBody.ErrorDescription).To(Equal("The client is not authorized to request an authorization code using this method"))
				})
			})

			Context("given INVALID password", func() {
				It("response with error", func() {
					FabricateUser("dev@nimblehq.co", "password")
					oauth_client := FabricateOAuthClient()
					body := MakeRequestBody(map[string]string{
						"client_id":     oauth_client.ClientID,
						"client_secret": oauth_client.ClientSecret,
						"username":      "dev@nimblehq.co",
						"password":      "invalid password",
						"grant_type":    "password",
					})
					response := MakeRequest("POST", "/api/v1/login", body)
					responseBody := ErrorResponse{}
					GetJSONResponseBody(response, responseBody)

					Expect(response.StatusCode).To(Equal(http.StatusUnauthorized))
					Expect(responseBody.Error).To(Equal("unauthorized_client"))
					Expect(responseBody.ErrorDescription).To(Equal("The client is not authorized to request an authorization code using this method"))
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
