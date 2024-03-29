package controllers_test

import (
	"net/http"

	controllers "go-google-scraper-challenge/lib/api/v1/controllers/oauth"
	"go-google-scraper-challenge/lib/api/v1/serializers"
	"go-google-scraper-challenge/test"
	. "go-google-scraper-challenge/test"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("OAuthClientsController", func() {
	Describe("POST /oauth/clients", func() {
		It("returns status OK", func() {
			ctx, response := MakeJSONRequest("POST", "/oauth/clients", nil, nil, nil)

			oauthClientsController := controllers.OAuthClientsController{}
			oauthClientsController.Create(ctx)

			jsonResponse := serializers.OAuthClientJSONResponse{}
			test.GetJSONResponseBody(response.Result(), &jsonResponse)

			Expect(response.Code).To(Equal(http.StatusOK))
		})

		It("returns correct response body", func() {
			ctx, response := MakeJSONRequest("POST", "/oauth/clients", nil, nil, nil)

			oauthClientsController := controllers.OAuthClientsController{}
			oauthClientsController.Create(ctx)

			jsonResponse := serializers.OAuthClientJSONResponse{}
			test.GetJSONResponseBody(response.Result(), &jsonResponse)

			Expect(jsonResponse.Data.ID).NotTo(BeEmpty())
			Expect(jsonResponse.Data.Attributes.ClientID).NotTo(BeEmpty())
			Expect(jsonResponse.Data.Attributes.ClientSecret).NotTo(BeEmpty())
		})
	})

	AfterEach(func() {
		CleanupDatabase([]string{"oauth2_clients", "oauth2_tokens"})
	})
})
