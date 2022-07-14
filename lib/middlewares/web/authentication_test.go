package webmiddlewares_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"go-google-scraper-challenge/bootstrap"
	"go-google-scraper-challenge/constants"
	webmiddlewares "go-google-scraper-challenge/lib/middlewares/web"
	"go-google-scraper-challenge/lib/models"
	. "go-google-scraper-challenge/test"

	"github.com/bxcodec/faker/v3"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Middleware", func() {
	Describe("#CurrentUser", func() {
		engine := gin.Default()
		engine = bootstrap.SetupSession(engine)
		engine.Use(webmiddlewares.CurrentUser)

		engine.GET("/test-current-user", func(ctx *gin.Context) {
			currentUser, ok := ctx.Get(constants.ContextCurrentUser)
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

		Context("given the session has user ID", func() {
			It("sets the current user to the context", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				responseRecorder := httptest.NewRecorder()
				cookie := FabricateAuthUserCookie(user.ID)
				request, _ := http.NewRequest("GET", "/test-current-user", nil)
				request.Header.Set("Cookie", cookie.Name+"="+cookie.Value)
				engine.ServeHTTP(responseRecorder, request)

				response := responseRecorder.Result()
				responseBody := GetResponseBody(response)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
				Expect(responseBody).To(Equal(fmt.Sprint(user.ID)))
			})
		})

		Context("given NO user ID in the session", func() {
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

		Context("given a fake user ID in the session", func() {
			It("sets NO current user to the context", func() {
				responseRecorder := httptest.NewRecorder()
				cookie := FabricateAuthUserCookie(999)
				request, _ := http.NewRequest("GET", "/test-current-user", nil)
				request.Header.Set("Cookie", cookie.Name+"="+cookie.Value)
				engine.ServeHTTP(responseRecorder, request)

				response := responseRecorder.Result()
				responseBody := GetResponseBody(response)

				Expect(response.StatusCode).To(Equal(http.StatusNotFound))
				Expect(responseBody).To(BeEmpty())
			})
		})
	})

	Describe("#EnsureGuestUser", func() {
		engine := gin.Default()
		engine = bootstrap.SetupSession(engine)
		engine.Use(webmiddlewares.CurrentUser)
		engine.Use(webmiddlewares.EnsureGuestUser)

		engine.GET("/test-ensure-guest-user", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "Does not redirect")
		})

		Context("given the session has user ID", func() {
			It("redirects to dashboard", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				responseRecorder := httptest.NewRecorder()
				cookie := FabricateAuthUserCookie(user.ID)
				request, _ := http.NewRequest("GET", "/test-ensure-guest-user", nil)
				request.Header.Set("Cookie", cookie.Name+"="+cookie.Value)
				engine.ServeHTTP(responseRecorder, request)

				response := responseRecorder.Result()
				currentPath := GetCurrentPath(response)

				Expect(response.StatusCode).To(Equal(http.StatusFound))
				Expect(currentPath).To(Equal(constants.WebRoutes["results"]["index"]))
			})
		})

		Context("given NO user ID in the session", func() {
			It("does NOT redirect to dashboard", func() {
				responseRecorder := httptest.NewRecorder()
				request, _ := http.NewRequest("GET", "/test-ensure-guest-user", nil)
				engine.ServeHTTP(responseRecorder, request)

				response := responseRecorder.Result()
				responseBody := GetResponseBody(response)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
				Expect(responseBody).To(Equal("Does not redirect"))
			})
		})
	})

	Describe("#EnsureAuthenticatedUser", func() {
		engine := gin.Default()
		engine = bootstrap.SetupSession(engine)
		engine.Use(webmiddlewares.CurrentUser)
		engine.Use(webmiddlewares.EnsureAuthenticatedUser)

		engine.GET("/test-ensure-authenticated-user", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "Does not redirect")
		})

		Context("given the session has user ID", func() {
			It("does NOT redirects to sign in screen", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				responseRecorder := httptest.NewRecorder()
				cookie := FabricateAuthUserCookie(user.ID)
				request, _ := http.NewRequest("GET", "/test-ensure-authenticated-user", nil)
				request.Header.Set("Cookie", cookie.Name+"="+cookie.Value)
				engine.ServeHTTP(responseRecorder, request)

				response := responseRecorder.Result()
				responseBody := GetResponseBody(response)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
				Expect(responseBody).To(Equal("Does not redirect"))
			})
		})

		Context("given NO user ID in the session", func() {
			It("redirects to sign in screen", func() {
				responseRecorder := httptest.NewRecorder()
				request, _ := http.NewRequest("GET", "/test-ensure-authenticated-user", nil)
				engine.ServeHTTP(responseRecorder, request)

				response := responseRecorder.Result()
				currentPath := GetCurrentPath(response)

				Expect(response.StatusCode).To(Equal(http.StatusFound))
				Expect(currentPath).To(Equal(constants.WebRoutes["sessions"]["new"]))
			})
		})
	})
})
