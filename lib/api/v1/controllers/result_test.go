package controllers_test

import (
	"net/http"

	"go-google-scraper-challenge/errors"
	"go-google-scraper-challenge/lib/api/v1/controllers"
	"go-google-scraper-challenge/test"
	. "go-google-scraper-challenge/test"

	"github.com/bxcodec/faker/v3"
	"github.com/google/jsonapi"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ResultsController", func() {
	Describe("POST /results", func() {
		Context("given an authenticated request", func() {
			It("returns status OK", func() {
				user := FabricateUser(faker.Email(), faker.Password())

				ctx, response := MakeAuthenticatedFormRequest("POST", "/results", nil, user)

				resultsController := controllers.ResultsController{}
				resultsController.Create(ctx)

				Expect(response.Code).To(Equal(http.StatusOK))
			})
		})

		Context("given an unauthenticated request", func() {
			It("returns status Unauthorized", func() {
				ctx, response := MakeFormRequest("POST", "/results", nil)

				resultsController := controllers.ResultsController{}
				resultsController.Create(ctx)

				Expect(response.Code).To(Equal(http.StatusUnauthorized))

				jsonResponse := &jsonapi.ErrorsPayload{}
				test.GetJSONResponseBody(response.Result(), &jsonResponse)

				Expect(jsonResponse.Errors[0].Title).To(Equal(errors.Titles[errors.ErrInvalidCredentials]))
				Expect(jsonResponse.Errors[0].Code).To(Equal(errors.ErrInvalidCredentials.Error()))
				Expect(jsonResponse.Errors[0].Detail).To(Equal(errors.Descriptions[errors.ErrInvalidCredentials]))
			})
		})
	})
})
