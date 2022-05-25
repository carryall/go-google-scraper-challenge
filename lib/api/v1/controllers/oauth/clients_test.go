package controllers_test

import (
	"net/http"

	controllers "go-google-scraper-challenge/lib/api/v1/controllers/oauth"
	"go-google-scraper-challenge/lib/services/oauth"
	"go-google-scraper-challenge/test"
	. "go-google-scraper-challenge/test"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("OAuthClientsController", func() {
	Describe("POST /oauth/clients", func() {
		It("returns status OK", func() {
			ctx, response := MakeJSONRequest("POST", "/oauth/clients", nil)

			oauthClientsController := controllers.OAuthClientsController{}
			oauthClientsController.Create(ctx)

			jsonResponse := oauth.OAuthClient{}
			test.GetJSONResponseBody(response.Result(), &jsonResponse)

			Expect(response.Code).To(Equal(http.StatusOK))
		})

		It("returns correct response body", func() {
			ctx, response := MakeJSONRequest("POST", "/oauth/clients", nil)

			oauthClientsController := controllers.OAuthClientsController{}
			oauthClientsController.Create(ctx)

			jsonResponse := oauth.OAuthClient{}
			test.GetJSONResponseBody(response.Result(), &jsonResponse)

			Expect(jsonResponse.ClientID).NotTo(BeEmpty())
			Expect(jsonResponse.ClientSecret).NotTo(BeEmpty())
		})
	})

	AfterEach(func() {
		CleanupDatabase([]string{"oauth2_clients", "oauth2_tokens"})
	})
})
