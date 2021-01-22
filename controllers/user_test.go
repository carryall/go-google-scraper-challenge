package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"

	. "go-google-scraper-challenge/helpers/test"
	"go-google-scraper-challenge/initializers"

	"github.com/beego/beego/v2/server/web"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UserController", func() {
	Describe("GET /signup", func() {
		It("renders with status 200", func() {

			request := HTTPRequest("GET", "/signup", nil)
			response := httptest.NewRecorder()
			web.BeeApp.Handlers.ServeHTTP(response, request)

			Expect(response.Code).To(Equal(http.StatusOK))
		})
	})

	Describe("POST /users", func() {
		Context("given valid params", func() {
			It("redirects to signup page", func() {
				form := url.Values{
					"email":                 {"dev@nimblehq.co"},
					"password":              {"password"},
					"password_confirmation": {"password"},
				}
				body := strings.NewReader(form.Encode())

				request := HTTPRequest("POST", "/users", body)
				request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
				response := httptest.NewRecorder()
				web.BeeApp.Handlers.ServeHTTP(response, request)
				currentPath := GetCurrentPath(response)

				Expect(response.Code).To(Equal(http.StatusFound))
				Expect(currentPath.Path).To(Equal("/signup"))
			})

			It("sets the success message", func() {
				form := url.Values{
					"email":                 {"dev@nimblehq.co"},
					"password":              {"password"},
					"password_confirmation": {"password"},
				}
				body := strings.NewReader(form.Encode())

				request := HTTPRequest("POST", "/users", body)
				request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
				response := httptest.NewRecorder()
				web.BeeApp.Handlers.ServeHTTP(response, request)

				flash := GetFlashMessage(response.Result().Cookies())

				Expect(flash.Data["success"]).To(HavePrefix("New User created with ID:"))
				Expect(flash.Data["error"]).To(BeEmpty())
			})
		})

		Context("given INVALID params", func() {
			It("redirects to signup page", func() {
				form := url.Values{
					"email":                 {""},
					"password":              {""},
					"password-confirmation": {""},
				}
				body := strings.NewReader(form.Encode())

				request := HTTPRequest("POST", "/users", body)
				request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
				response := httptest.NewRecorder()
				web.BeeApp.Handlers.ServeHTTP(response, request)
				currentPath := GetCurrentPath(response)

				Expect(response.Code).To(Equal(http.StatusFound))
				Expect(currentPath.Path).To(Equal("/signup"))
			})

			It("sets error message", func() {
				form := url.Values{
					"email":                 {""},
					"password":              {""},
					"password-confirmation": {""},
				}
				body := strings.NewReader(form.Encode())

				request := HTTPRequest("POST", "/users", body)
				request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
				response := httptest.NewRecorder()
				web.BeeApp.Handlers.ServeHTTP(response, request)

				flash := GetFlashMessage(response.Result().Cookies())

				Expect(flash.Data["error"]).NotTo(BeEmpty())
				Expect(flash.Data["success"]).To(BeEmpty())
			})
		})
	})

	AfterEach(func() {
		initializers.CleanupDatabase("user")
	})
})
