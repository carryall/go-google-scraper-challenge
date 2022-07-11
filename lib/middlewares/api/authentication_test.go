package apimiddlewares_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"go-google-scraper-challenge/bootstrap"
	"go-google-scraper-challenge/errors"
	apimiddlewares "go-google-scraper-challenge/lib/middlewares/api"
	"go-google-scraper-challenge/lib/models"
	. "go-google-scraper-challenge/test"

	"github.com/bxcodec/faker/v3"
	"github.com/gin-gonic/gin"
	"github.com/google/jsonapi"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Middleware", func() {
	Describe("#CurrentUser", func() {
		engine := gin.Default()
		engine = bootstrap.SetupSession(engine)
		engine.Use(apimiddlewares.CurrentUser)

		engine.GET("/test-current-user", func(ctx *gin.Context) {
			currentUser, ok := ctx.Get("CurrentUser")
			if !ok {
				ctx.String(http.StatusNotFound, "")

				return
			}

			if currentUser != nil {
				user := currentUser.(*models.User)
				ctx.String(http.StatusOK, fmt.Sprint(user.ID))
			} else {
				ctx.String(http.StatusNotFound, "")
			}
		})

		Context("given a valid access token in request header", func() {
			It("sets the current user to the context", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				responseRecorder := httptest.NewRecorder()
				accessToken := FabricateAuthToken(user.ID)
				request, _ := http.NewRequest("GET", "/test-current-user", nil)
				request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
				engine.ServeHTTP(responseRecorder, request)

				response := responseRecorder.Result()
				responseBody := GetResponseBody(response)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
				Expect(responseBody).To(Equal(fmt.Sprint(user.ID)))
			})
		})

		Context("given NO valid access token in request header", func() {
			It("sets NO current user to the context", func() {
				responseRecorder := httptest.NewRecorder()
				request, _ := http.NewRequest("GET", "/test-current-user", nil)
				engine.ServeHTTP(responseRecorder, request)

				response := responseRecorder.Result()
				responseBody := GetResponseBody(response)

				Expect(response.StatusCode).To(Equal(http.StatusNotFound))
				Expect(responseBody).To(BeEmpty())
			})
		})

		Context("given an INVALID access token in request header", func() {
			It("sets NO current user to the context", func() {
				responseRecorder := httptest.NewRecorder()
				accessToken := FabricateAuthToken(999)
				request, _ := http.NewRequest("GET", "/test-current-user", nil)
				request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
				engine.ServeHTTP(responseRecorder, request)

				response := responseRecorder.Result()
				responseBody := GetResponseBody(response)

				Expect(response.StatusCode).To(Equal(http.StatusNotFound))
				Expect(responseBody).To(BeEmpty())
			})
		})
	})

	Describe("#EnsureAuthenticatedUser", func() {
		engine := gin.Default()
		engine = bootstrap.SetupSession(engine)
		engine.Use(apimiddlewares.CurrentUser)
		engine.Use(apimiddlewares.EnsureAuthenticatedUser)

		engine.GET("/test-ensure-authenticated-user", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "Success response")
		})

		Context("given a valid access token in request header", func() {
			It("response with no error", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				responseRecorder := httptest.NewRecorder()
				accessToken := FabricateAuthToken(user.ID)
				request, _ := http.NewRequest("GET", "/test-ensure-authenticated-user", nil)
				request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
				engine.ServeHTTP(responseRecorder, request)

				response := responseRecorder.Result()
				responseBody := GetResponseBody(response)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
				Expect(responseBody).To(Equal("Success response"))
			})
		})

		Context("given NO access token in request header", func() {
			It("responses with unauthorized error", func() {
				responseRecorder := httptest.NewRecorder()
				request, _ := http.NewRequest("GET", "/test-ensure-authenticated-user", nil)
				engine.ServeHTTP(responseRecorder, request)

				response := responseRecorder.Result()

				Expect(response.StatusCode).To(Equal(http.StatusUnauthorized))

				jsonResponse := &jsonapi.ErrorsPayload{}
				GetJSONResponseBody(response, &jsonResponse)

				Expect(jsonResponse.Errors[0].Title).To(Equal(errors.Titles[errors.ErrUnauthorizedUser]))
				Expect(jsonResponse.Errors[0].Code).To(Equal(errors.ErrUnauthorizedUser.Error()))
				Expect(jsonResponse.Errors[0].Detail).To(Equal(errors.Descriptions[errors.ErrUnauthorizedUser]))
			})
		})
	})
})
