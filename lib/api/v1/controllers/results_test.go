package controllers_test

import (
	"fmt"
	"net/http"

	"go-google-scraper-challenge/errors"
	"go-google-scraper-challenge/lib/api/v1/controllers"
	"go-google-scraper-challenge/lib/api/v1/serializers"
	"go-google-scraper-challenge/lib/models"
	. "go-google-scraper-challenge/test"

	"github.com/bxcodec/faker/v3"
	"github.com/gin-gonic/gin"
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

			It("returns the number of result relations", func() {
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

			Context("given a keyword param", func() {
				It("returns the results that contain the keyword", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					result := FabricateResultWithParams(user, "mechanical keyboard", models.ResultStatusCompleted)
					FabricateResultWithParams(user, "Khao Yai Hotel", models.ResultStatusCompleted)

					ctx, response := MakeJSONRequest("GET", fmt.Sprintf("/results?keyword=%s", "key"), nil, nil, user)

					resultsController := controllers.ResultsController{}
					resultsController.List(ctx)

					Expect(response.Code).To(Equal(http.StatusOK))

					jsonArrayResponse := &serializers.ResultsJSONResponse{}
					GetJSONResponseBody(response.Result(), &jsonArrayResponse)

					Expect(jsonArrayResponse.Data).To(HaveLen(1))
					Expect(jsonArrayResponse.Data[0].ID).To(Equal(fmt.Sprint(result.ID)))
					Expect(jsonArrayResponse.Data[0].Attributes.Keyword).To(Equal(result.Keyword))
				})

				Context("given no result match the keyword", func() {
					It("returns an empty array", func() {
						user := FabricateUser(faker.Email(), faker.Password())
						FabricateResultWithParams(user, "mechanical keyboard", models.ResultStatusCompleted)
						FabricateResultWithParams(user, "Khao Yai Hotel", models.ResultStatusCompleted)

						ctx, response := MakeJSONRequest("GET", fmt.Sprintf("/results?keyword=%s", "other"), nil, nil, user)

						resultsController := controllers.ResultsController{}
						resultsController.List(ctx)

						Expect(response.Code).To(Equal(http.StatusOK))

						jsonArrayResponse := &serializers.ResultsJSONResponse{}
						GetJSONResponseBody(response.Result(), &jsonArrayResponse)

						Expect(jsonArrayResponse.Data).To(HaveLen(0))
					})
				})
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
		Context("given a valid result ID", func() {
			It("returns status ok", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				result := FabricateResult(user)
				ctx, response := MakeJSONRequest("GET", fmt.Sprintf("/results/%d", result.ID), nil, nil, user)
				ctx.Params = append(ctx.Params, gin.Param{Key: "id", Value: fmt.Sprint(result.ID)})

				resultsController := controllers.ResultsController{}
				resultsController.Show(ctx)

				Expect(response.Code).To(Equal(http.StatusOK))
			})

			It("returns the result detail", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				result := FabricateResult(user)
				ctx, response := MakeJSONRequest("GET", fmt.Sprintf("/results/%d", result.ID), nil, nil, user)
				ctx.Params = append(ctx.Params, gin.Param{Key: "id", Value: fmt.Sprint(result.ID)})

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

			It("returns the result relations", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				result := FabricateResult(user)
				adLink := FabricateAdLink(result)
				link := FabricateLink(result)
				ctx, response := MakeJSONRequest("GET", fmt.Sprintf("/results/%d", result.ID), nil, nil, user)
				ctx.Params = append(ctx.Params, gin.Param{Key: "id", Value: fmt.Sprint(result.ID)})

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

		Context("given a result ID that does NOT belong to the authenticated user", func() {
			It("returns error not found", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				anotherUser := FabricateUser(faker.Email(), faker.Password())
				result := FabricateResult(anotherUser)
				ctx, response := MakeJSONRequest("GET", fmt.Sprintf("/results/%d", result.ID), nil, nil, user)
				ctx.Params = append(ctx.Params, gin.Param{Key: "id", Value: fmt.Sprint(result.ID)})

				resultsController := controllers.ResultsController{}
				resultsController.Show(ctx)

				Expect(response.Code).To(Equal(http.StatusNotFound))

				jsonResponse := &jsonapi.ErrorsPayload{}
				GetJSONResponseBody(response.Result(), &jsonResponse)

				Expect(jsonResponse.Errors[0].Title).To(Equal(errors.Titles[errors.ErrNotFound]))
				Expect(jsonResponse.Errors[0].Code).To(Equal(errors.ErrNotFound.Error()))
				Expect(jsonResponse.Errors[0].Detail).To(Equal("record not found"))
			})
		})

		Context("given non-existing result ID", func() {
			It("returns error not found", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				ctx, response := MakeJSONRequest("GET", fmt.Sprintf("/results/%d", 999), nil, nil, user)
				ctx.Params = append(ctx.Params, gin.Param{Key: "id", Value: "999"})

				resultsController := controllers.ResultsController{}
				resultsController.Show(ctx)

				Expect(response.Code).To(Equal(http.StatusNotFound))

				jsonResponse := &jsonapi.ErrorsPayload{}
				GetJSONResponseBody(response.Result(), &jsonResponse)

				Expect(jsonResponse.Errors[0].Title).To(Equal(errors.Titles[errors.ErrNotFound]))
				Expect(jsonResponse.Errors[0].Code).To(Equal(errors.ErrNotFound.Error()))
				Expect(jsonResponse.Errors[0].Detail).To(Equal("record not found"))
			})
		})

		Context("given an INVALID result ID", func() {
			It("returns error bad request", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				ctx, response := MakeJSONRequest("GET", fmt.Sprintf("/results/%s", "invalid"), nil, nil, user)
				ctx.Params = append(ctx.Params, gin.Param{Key: "id", Value: "invalid"})

				resultsController := controllers.ResultsController{}
				resultsController.Show(ctx)

				Expect(response.Code).To(Equal(http.StatusBadRequest))

				jsonResponse := &jsonapi.ErrorsPayload{}
				GetJSONResponseBody(response.Result(), &jsonResponse)

				Expect(jsonResponse.Errors[0].Title).To(Equal(errors.Titles[errors.ErrInvalidRequest]))
				Expect(jsonResponse.Errors[0].Code).To(Equal(errors.ErrInvalidRequest.Error()))
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

					Expect(response.Code).To(Equal(http.StatusCreated))
				})

				It("returns list of result with the givern keyword", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					header, body := CreateRequestInfoFormFile("test/fixtures/files/valid.csv")

					ctx, response := MakeJSONRequest("POST", "/results", header, body, user)

					resultsController := controllers.ResultsController{}
					resultsController.Create(ctx)

					jsonArrayResponse := &serializers.ResultsJSONResponse{}
					GetJSONResponseBody(response.Result(), &jsonArrayResponse)

					Expect(jsonArrayResponse.Data).To(HaveLen(2))
					Expect(jsonArrayResponse.Data[0].ID).NotTo(BeNil())
					Expect(jsonArrayResponse.Data[0].Attributes.Keyword).To(Equal("ergonomic chair"))
					Expect(jsonArrayResponse.Data[0].Attributes.UserID).To(Equal(user.ID))
					Expect(jsonArrayResponse.Data[1].ID).NotTo(BeNil())
					Expect(jsonArrayResponse.Data[1].Attributes.Keyword).To(Equal("mechanical keyboard"))
					Expect(jsonArrayResponse.Data[1].Attributes.UserID).To(Equal(user.ID))
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
