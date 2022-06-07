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
			Context("given a valid file", func() {
				It("returns status OK", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					header, body := CreateRequestInfoFormFile("test/fixtures/files/valid.csv")

					ctx, response := MakeUploadRequest("POST", "/results", header, body, user)

					resultsController := controllers.ResultsController{}
					resultsController.Create(ctx)

					Expect(response.Code).To(Equal(http.StatusOK))
				})
			})

			Context("given an empty file", func() {
				It("returns status bad request", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					header, body := CreateRequestInfoFormFile("test/fixtures/files/empty.csv")

					ctx, response := MakeUploadRequest("POST", "/results", header, body, user)

					resultsController := controllers.ResultsController{}
					resultsController.Create(ctx)

					Expect(response.Code).To(Equal(http.StatusBadRequest))

					jsonResponse := &jsonapi.ErrorsPayload{}
					test.GetJSONResponseBody(response.Result(), &jsonResponse)

					Expect(jsonResponse.Errors[0].Title).To(Equal(errors.Titles[errors.ErrInvalidRequest]))
					Expect(jsonResponse.Errors[0].Code).To(Equal(errors.ErrInvalidRequest.Error()))
					Expect(jsonResponse.Errors[0].Detail).To(Equal("Keywords: cannot be blank."))
				})
			})

			Context("given an INVALID file type", func() {
				It("returns status bad request", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					header, body := CreateRequestInfoFormFile("test/fixtures/files/text.txt")

					ctx, response := MakeUploadRequest("POST", "/results", header, body, user)

					resultsController := controllers.ResultsController{}
					resultsController.Create(ctx)

					Expect(response.Code).To(Equal(http.StatusBadRequest))

					jsonResponse := &jsonapi.ErrorsPayload{}
					test.GetJSONResponseBody(response.Result(), &jsonResponse)

					Expect(jsonResponse.Errors[0].Title).To(Equal(errors.Titles[errors.ErrInvalidRequest]))
					Expect(jsonResponse.Errors[0].Code).To(Equal(errors.ErrInvalidRequest.Error()))
					Expect(jsonResponse.Errors[0].Detail).To(Equal("File: wrong file type"))
				})
			})
		})

		Context("given an unauthenticated request", func() {
			It("returns status Unauthorized", func() {
				header, body := CreateRequestInfoFormFile("test/fixtures/files/valid.csv")

				ctx, response := MakeUploadRequest("POST", "/results", header, body, nil)

				resultsController := controllers.ResultsController{}
				resultsController.Create(ctx)

				Expect(response.Code).To(Equal(http.StatusUnauthorized))

				jsonResponse := &jsonapi.ErrorsPayload{}
				test.GetJSONResponseBody(response.Result(), &jsonResponse)

				Expect(jsonResponse.Errors[0].Title).To(Equal(errors.Titles[errors.ErrUnauthorizedUser]))
				Expect(jsonResponse.Errors[0].Code).To(Equal(errors.ErrUnauthorizedUser.Error()))
				Expect(jsonResponse.Errors[0].Detail).To(Equal("invalid access token"))
			})
		})
	})
})
