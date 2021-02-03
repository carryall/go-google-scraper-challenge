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
				data := url.Values{}
				data.Set("client_id", oauth_client.ClientID)
				data.Set("client_secret", oauth_client.ClientSecret)
				data.Set("username", "dev@nimblehq.co")
				data.Set("password", "password")
				data.Set("grant_type", "password")
				body := strings.NewReader(data.Encode())
				response := MakeRequest("POST", "/api/v1/login", body)
				responseBody := LoginResponse{}
				GetJSONResponseBody(response, responseBody)

				Expect(responseBody.AccessToken).NotTo(BeNil())
				Expect(responseBody.RefreshToken).NotTo(BeNil())
				Expect(responseBody.ExpiresIn).NotTo(BeNil())
				Expect(responseBody.TokenType).NotTo(BeNil())
			})
		})
	})

	AfterEach(func() {
		initializers.CleanupDatabase("user")
	})
})
