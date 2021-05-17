package controllers_test

import (
	"net/http"

	"go-google-scraper-challenge/initializers"
	. "go-google-scraper-challenge/tests/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UserController", func() {
	Describe("GET /signup", func() {
		Context("given user is not signed in", func() {
			It("renders with status 200", func() {
				response := MakeRequest("GET", "/signup", nil)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
			})
		})

		Context("given user is already signed in", func() {
			It("redirects to root path", func() {
				user := FabricateUser("dev@nimblehq.co", "password")
				response := MakeAuthenticatedRequest("GET", "/signup", nil, user)
				currentPath := GetCurrentPath(response)

				Expect(response.StatusCode).To(Equal(http.StatusFound))
				Expect(currentPath).To(Equal("/"))
			})
		})
	})

	Describe("POST /users", func() {
		Context("given user is not signed in", func() {
			Context("given valid params", func() {
				It("redirects to signup page", func() {
					body := GenerateRequestBody(map[string]string{
						"email":                 "dev@nimblehq.co",
						"password":              "password",
						"password_confirmation": "password",
					})

					response := MakeRequest("POST", "/users", body)
					currentPath := GetCurrentPath(response)

					Expect(response.StatusCode).To(Equal(http.StatusFound))
					Expect(currentPath).To(Equal("/signin"))
				})

				It("sets the success message", func() {
					body := GenerateRequestBody(map[string]string{
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
					body := GenerateRequestBody(map[string]string{
						"email":                 "",
						"password":              "",
						"password-confirmation": "",
					})

					response := MakeRequest("POST", "/users", body)
					currentPath := GetCurrentPath(response)

					Expect(response.StatusCode).To(Equal(http.StatusFound))
					Expect(currentPath).To(Equal("/signup"))
				})

				It("sets error message", func() {
					body := GenerateRequestBody(map[string]string{
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

		Context("given user is already signed on", func() {
			It("returns error", func() {
				user := FabricateUser("dev@nimblehq.co", "password")
				body := GenerateRequestBody(map[string]string{
					"email":                 "dev@nimblehq.co",
					"password":              "password",
					"password_confirmation": "password",
				})

				response := MakeAuthenticatedRequest("POST", "/users", body, user)

				Expect(response.StatusCode).To(Equal(http.StatusMethodNotAllowed))
			})
		})
	})

	AfterEach(func() {
		initializers.CleanupDatabase([]string{"users"})
	})
})
