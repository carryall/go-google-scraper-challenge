package controllers_test

import (
	"net/http"

	"go-google-scraper-challenge/constants"
	"go-google-scraper-challenge/initializers"
	. "go-google-scraper-challenge/tests/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("OAuthClientController", func() {
	Describe("GET /oauth_client", func() {
		It("renders with status 200", func() {
			response := MakeRequest("GET", "/oauth_client", nil)

			Expect(response.StatusCode).To(Equal(http.StatusOK))
		})
	})

	Describe("POST /oauth_client", func() {
		It("redirects to oauth client detail page", func() {
			body := GenerateRequestBody(nil)
			response := MakeRequest("POST", "/oauth_client", body)
			currentPath := GetCurrentPath(response)

			Expect(response.StatusCode).To(Equal(http.StatusFound))
			Expect(currentPath).To(MatchRegexp(`\/oauth_client\/\b`))
		})

		It("sets the success message", func() {
			body := GenerateRequestBody(nil)
			response := MakeRequest("POST", "/oauth_client", body)
			flash := GetFlashMessage(response.Cookies())

			Expect(flash.Data["success"]).To(Equal(constants.OAuthClientCreateSuccess))
			Expect(flash.Data["error"]).To(BeEmpty())
		})
	})

	Describe("GET /oauth_client/:id", func() {
		Context("given valid client id", func() {
			It("renders with status 200", func() {
				client := FabricateOAuthClient()
				response := MakeRequest("GET", "/oauth_client/"+client.ClientID, nil)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
			})

			It("shows OAuth client detail", func() {
				client := FabricateOAuthClient()
				response := MakeRequest("GET", "/oauth_client/"+client.ClientID, nil)
				responseBody := GetResponseBody(response)

				Expect(responseBody).To(ContainSubstring(client.ClientID))
				Expect(responseBody).To(ContainSubstring(client.ClientSecret))
			})
		})

		Context("given INVALID client id", func() {
			It("redirects to oauth client page", func() {
				response := MakeRequest("GET", "/oauth_client/invalid_id", nil)

				Expect(response.StatusCode).To(Equal(http.StatusFound))
			})

			It("sets the error message", func() {
				response := MakeRequest("GET", "/oauth_client/invalid_id", nil)
				flash := GetFlashMessage(response.Cookies())

				Expect(flash.Data["error"]).To(Equal(constants.OAuthClientNotFound))
				Expect(flash.Data["success"]).To(BeEmpty())
			})
		})
	})

	AfterEach(func() {
		initializers.CleanupDatabase([]string{"oauth2_clients", "oauth2_tokens"})
	})
})
