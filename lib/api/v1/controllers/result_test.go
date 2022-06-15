package controllers_test

import (
	"fmt"
	"net/http"

	"go-google-scraper-challenge/errors"
	"go-google-scraper-challenge/lib/api/v1/controllers"
	"go-google-scraper-challenge/lib/api/v1/serializers"
	. "go-google-scraper-challenge/test"

	"github.com/bxcodec/faker/v3"
	"github.com/google/jsonapi"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ResultsController", func() {
	Describe("GET /results", func() {
		Context("given an authenticated request", func() {
			It("returns status OK", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				ctx, response := MakeJSONRequest("GET", "/results", nil, nil, user)

				resultsController := controllers.ResultsController{}
				resultsController.List(ctx)

				Expect(response.Code).To(Equal(http.StatusOK))
			})

			It("returns a list of results that belongs to the user", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				anotherUser := FabricateUser(faker.Email(), faker.Password())
				expectedResult := FabricateResult(user)
				FabricateResult(anotherUser)
				ctx, response := MakeJSONRequest("GET", "/results", nil, nil, user)

				resultsController := controllers.ResultsController{}
				resultsController.List(ctx)

				Expect(response.Code).To(Equal(http.StatusOK))

				jsonArrayResponse := &serializers.ResultsJSONResponse{}
				GetJSONResponseBody(response.Result(), &jsonArrayResponse)

				Expect(jsonArrayResponse.Data).To(HaveLen(1))
				Expect(jsonArrayResponse.Data[0].ID).To(Equal(fmt.Sprint(expectedResult.ID)))
				Expect(jsonArrayResponse.Data[0].Attributes.Keyword).To(Equal(expectedResult.Keyword))
				Expect(jsonArrayResponse.Data[0].Attributes.UserID).To(Equal(user.ID))
				Expect(jsonArrayResponse.Data[0].Attributes.CreatedAt).To(Equal(expectedResult.CreatedAt.Unix()))
				Expect(jsonArrayResponse.Data[0].Attributes.UpdatedAt).To(Equal(expectedResult.UpdatedAt.Unix()))
				Expect(jsonArrayResponse.Included[0].Type).To(Equal("user"))
				Expect(jsonArrayResponse.Included[0].ID).To(Equal(fmt.Sprint(user.ID)))
				Expect(jsonArrayResponse.Included[0].Attributes["email"]).To(Equal(user.Email))
			})

			It("returns number of result relations", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				result := FabricateResult(user)
				FabricateLink(result)
				FabricateAdLink(result)
				FabricateAdLink(result)
				ctx, response := MakeJSONRequest("GET", "/results", nil, nil, user)

				resultsController := controllers.ResultsController{}
				resultsController.List(ctx)

				Expect(response.Code).To(Equal(http.StatusOK))

				jsonArrayResponse := &serializers.ResultsJSONResponse{}
				GetJSONResponseBody(response.Result(), &jsonArrayResponse)

				Expect(jsonArrayResponse.Data).To(HaveLen(1))
				Expect(jsonArrayResponse.Data[0].Attributes.AdLinkCount).To(Equal(2))
				Expect(jsonArrayResponse.Data[0].Attributes.LinkCount).To(Equal(1))
			})
		})

		Context("given an unauthenticated request", func() {
			It("returns status Unauthorized", func() {
				ctx, response := MakeJSONRequest("GET", "/results", nil, nil, nil)

				resultsController := controllers.ResultsController{}
				resultsController.List(ctx)

				Expect(response.Code).To(Equal(http.StatusUnauthorized))

				jsonResponse := &jsonapi.ErrorsPayload{}
				GetJSONResponseBody(response.Result(), &jsonResponse)

				Expect(jsonResponse.Errors[0].Title).To(Equal(errors.Titles[errors.ErrUnauthorizedUser]))
				Expect(jsonResponse.Errors[0].Code).To(Equal(errors.ErrUnauthorizedUser.Error()))
				Expect(jsonResponse.Errors[0].Detail).To(Equal("invalid access token"))
			})
		})
	})

	Describe("GET /results/:id", func() {
		Context("given valid result ID", func() {
			It("returns status ok", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				result := FabricateResult(user)
				ctx, response := MakeJSONRequest("GET", fmt.Sprintf("/results/%d", result.ID), nil, nil, user)

				resultsController := controllers.ResultsController{}
				resultsController.Show(ctx)

				Expect(response.Code).To(Equal(http.StatusOK))
			})

			It("returns result detail", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				result := FabricateResult(user)
				ctx, response := MakeJSONRequest("GET", fmt.Sprintf("/results/%d", result.ID), nil, nil, user)

				resultsController := controllers.ResultsController{}
				resultsController.Show(ctx)

				Expect(response.Code).To(Equal(http.StatusOK))

				jsonResponse := &serializers.ResultDetailJSONResponse{}
				GetJSONResponseBody(response.Result(), &jsonResponse)

				Expect(jsonResponse.Data.ID).To(Equal(fmt.Sprint(result.ID)))
				Expect(jsonResponse.Data.Attributes.CreatedAt).To(Equal(result.CreatedAt.Unix()))
				Expect(jsonResponse.Data.Attributes.UpdatedAt).To(Equal(result.UpdatedAt.Unix()))
				Expect(jsonResponse.Data.Attributes.Keyword).To(Equal(result.Keyword))
				Expect(jsonResponse.Data.Attributes.PageCache).To(Equal(result.PageCache))
				Expect(jsonResponse.Data.Attributes.Status).To(Equal(result.Status))
				Expect(jsonResponse.Data.Attributes.UserID).To(Equal(result.UserID))
			})

			It("returns result relations", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				result := FabricateResult(user)
				adLink := FabricateAdLink(result)
				link := FabricateLink(result)
				ctx, response := MakeJSONRequest("GET", fmt.Sprintf("/results/%d", result.ID), nil, nil, user)

				resultsController := controllers.ResultsController{}
				resultsController.Show(ctx)

				Expect(response.Code).To(Equal(http.StatusOK))

				jsonResponse := &serializers.ResultDetailJSONResponse{}
				GetJSONResponseBody(response.Result(), &jsonResponse)

				Expect(jsonResponse.Data.ID).To(Equal(fmt.Sprint(result.ID)))
				Expect(jsonResponse.Data.Relationships.AdLinks.Data).To(HaveLen(1))
				Expect(jsonResponse.Data.Relationships.AdLinks.Data[0].ID).To(Equal(fmt.Sprint(adLink.ID)))
				Expect(jsonResponse.Data.Relationships.AdLinks.Data[0].Type).To(Equal("ad_link"))
				Expect(jsonResponse.Data.Relationships.Links.Data).To(HaveLen(1))
				Expect(jsonResponse.Data.Relationships.Links.Data[0].ID).To(Equal(fmt.Sprint(link.ID)))
				Expect(jsonResponse.Data.Relationships.Links.Data[0].Type).To(Equal("link"))
			})
		})

		Context("given INVALID result ID", func() {
			It("returns not found error", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				ctx, response := MakeJSONRequest("GET", fmt.Sprintf("/results/%d", 999), nil, nil, user)

				resultsController := controllers.ResultsController{}
				resultsController.Show(ctx)

				Expect(response.Code).To(Equal(http.StatusBadRequest))
			})
		})

		Context("given NO result ID", func() {
			It("returns not found", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				ctx, response := MakeJSONRequest("GET", fmt.Sprintf("/results/%s", "invalid"), nil, nil, user)

				resultsController := controllers.ResultsController{}
				resultsController.Show(ctx)

				Expect(response.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})

	Describe("POST /results", func() {
		Context("given an authenticated request", func() {
			Context("given a valid file", func() {
				It("returns status OK", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					header, body := CreateRequestInfoFormFile("test/fixtures/files/valid.csv")

					ctx, response := MakeJSONRequest("POST", "/results", header, body, user)

					resultsController := controllers.ResultsController{}
					resultsController.Create(ctx)

					Expect(response.Code).To(Equal(http.StatusOK))
				})

				It("returns list of result with the givern keyword", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					header, body := CreateRequestInfoFormFile("test/fixtures/files/valid.csv")

					ctx, response := MakeJSONRequest("POST", "/results", header, body, user)

					resultsController := controllers.ResultsController{}
					resultsController.Create(ctx)

					jsonArrayResponse := &serializers.ResultsJSONResponse{}
					GetJSONResponseBody(response.Result(), &jsonArrayResponse)

					Expect(jsonArrayResponse.Data[0].ID).NotTo(BeNil())
					Expect(jsonArrayResponse.Data[0].Attributes.Keyword).To(Equal("ergonomic chair"))
					Expect(jsonArrayResponse.Data[0].Attributes.UserID).To(Equal(user.ID))
					Expect(jsonArrayResponse.Data[0].Attributes.CreatedAt).To(BeNumerically(">", 0))
					Expect(jsonArrayResponse.Data[0].Attributes.UpdatedAt).To(BeNumerically(">", 0))
				})
			})

			Context("given an empty file", func() {
				It("returns status bad request", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					header, body := CreateRequestInfoFormFile("test/fixtures/files/empty.csv")

					ctx, response := MakeJSONRequest("POST", "/results", header, body, user)

					resultsController := controllers.ResultsController{}
					resultsController.Create(ctx)

					Expect(response.Code).To(Equal(http.StatusBadRequest))

					jsonResponse := &jsonapi.ErrorsPayload{}
					GetJSONResponseBody(response.Result(), &jsonResponse)

					Expect(jsonResponse.Errors[0].Title).To(Equal(errors.Titles[errors.ErrInvalidRequest]))
					Expect(jsonResponse.Errors[0].Code).To(Equal(errors.ErrInvalidRequest.Error()))
					Expect(jsonResponse.Errors[0].Detail).To(Equal("Keywords: cannot be blank."))
				})
			})

			Context("given a file with too many keywords", func() {
				It("returns status bad request", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					header, body := CreateRequestInfoFormFile("test/fixtures/files/invalid.csv")

					ctx, response := MakeJSONRequest("POST", "/results", header, body, user)

					resultsController := controllers.ResultsController{}
					resultsController.Create(ctx)

					Expect(response.Code).To(Equal(http.StatusBadRequest))

					jsonResponse := &jsonapi.ErrorsPayload{}
					GetJSONResponseBody(response.Result(), &jsonResponse)

					Expect(jsonResponse.Errors[0].Title).To(Equal(errors.Titles[errors.ErrInvalidRequest]))
					Expect(jsonResponse.Errors[0].Code).To(Equal(errors.ErrInvalidRequest.Error()))
					Expect(jsonResponse.Errors[0].Detail).To(Equal("Keywords: the length must be between 1 and 1000."))
				})
			})

			Context("given an INVALID file type", func() {
				It("returns status bad request", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					header, body := CreateRequestInfoFormFile("test/fixtures/files/text.txt")

					ctx, response := MakeJSONRequest("POST", "/results", header, body, user)

					resultsController := controllers.ResultsController{}
					resultsController.Create(ctx)

					Expect(response.Code).To(Equal(http.StatusBadRequest))

					jsonResponse := &jsonapi.ErrorsPayload{}
					GetJSONResponseBody(response.Result(), &jsonResponse)

					Expect(jsonResponse.Errors[0].Title).To(Equal(errors.Titles[errors.ErrInvalidRequest]))
					Expect(jsonResponse.Errors[0].Code).To(Equal(errors.ErrInvalidRequest.Error()))
					Expect(jsonResponse.Errors[0].Detail).To(Equal("FileHeader: invalid file type."))
				})
			})
		})

		Context("given an unauthenticated request", func() {
			It("returns status Unauthorized", func() {
				header, body := CreateRequestInfoFormFile("test/fixtures/files/valid.csv")

				ctx, response := MakeJSONRequest("POST", "/results", header, body, nil)

				resultsController := controllers.ResultsController{}
				resultsController.Create(ctx)

				Expect(response.Code).To(Equal(http.StatusUnauthorized))

				jsonResponse := &jsonapi.ErrorsPayload{}
				GetJSONResponseBody(response.Result(), &jsonResponse)

				Expect(jsonResponse.Errors[0].Title).To(Equal(errors.Titles[errors.ErrUnauthorizedUser]))
				Expect(jsonResponse.Errors[0].Code).To(Equal(errors.ErrUnauthorizedUser.Error()))
				Expect(jsonResponse.Errors[0].Detail).To(Equal("invalid access token"))
			})
		})
	})

	AfterEach(func() {
		CleanupDatabase([]string{"users", "results", "ad_links", "links"})
	})
})
