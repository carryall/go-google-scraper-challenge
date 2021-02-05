package controllers_test

import (
	"net/http"

	"go-google-scraper-challenge/initializers"
	. "go-google-scraper-challenge/test/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UserController", func() {
	Describe("GET /signup", func() {
		It("renders with status 200", func() {
			response := MakeRequest("GET", "/signup", nil)

			Expect(response.StatusCode).To(Equal(http.StatusOK))
		})
	})

	Describe("POST /users", func() {
		Context("given valid params", func() {
			It("redirects to signup page", func() {
				body := RequestBody(map[string]string{
					"email":                 "dev@nimblehq.co",
					"password":              "password",
					"password_confirmation": "password",
				})

				response := MakeRequest("POST", "/users", body)
				currentPath := GetCurrentPath(response)

				Expect(response.StatusCode).To(Equal(http.StatusFound))
				Expect(currentPath.Path).To(Equal("/signup"))
			})

			It("sets the success message", func() {
				body := RequestBody(map[string]string{
					"email":                 "dev@nimblehq.co",
					"password":              "password",
					"password_confirmation": "password",
				})

				response := MakeRequest("POST", "/users", body)
				flash := GetFlashMessage(response.Cookies())

				Expect(flash.Data["success"]).To(Equal("The user was successfully created"))
				Expect(flash.Data["error"]).To(BeEmpty())
			})
		})

		Context("given INVALID params", func() {
			It("redirects to signup page", func() {
				body := RequestBody(map[string]string{
					"email":                 "",
					"password":              "",
					"password-confirmation": "",
				})

				response := MakeRequest("POST", "/users", body)
				currentPath := GetCurrentPath(response)

				Expect(response.StatusCode).To(Equal(http.StatusFound))
				Expect(currentPath.Path).To(Equal("/signup"))
			})

			It("sets error message", func() {
				body := RequestBody(map[string]string{
					"email":                 "",
					"password":              "",
					"password-confirmation": "",
				})

				response := MakeRequest("POST", "/users", body)
				flash := GetFlashMessage(response.Cookies())

				Expect(flash.Data["error"]).NotTo(BeEmpty())
				Expect(flash.Data["success"]).To(BeEmpty())
			})
		})
	})

	AfterEach(func() {
		initializers.CleanupDatabase("users")
	})
})
