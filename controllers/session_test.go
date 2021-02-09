package controllers_test

import (
	"net/http"

	"go-google-scraper-challenge/controllers"
	"go-google-scraper-challenge/initializers"
	. "go-google-scraper-challenge/test/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SessionController", func() {
	Describe("GET /signin", func() {
		It("renders with status 200", func() {
			response := MakeRequest("GET", "/signin", nil)

			Expect(response.StatusCode).To(Equal(http.StatusOK))
		})
	})

	Describe("POST /sessions", func() {
		Context("given valid params", func() {
			It("redirects to root path", func() {
				FabricateUser("dev@nimblehq.co", "password")
				body := GenerateRequestBody(map[string]string{
					"email":    "dev@nimblehq.co",
					"password": "password",
				})
				response := MakeRequest("POST", "/sessions", body)
				currentPath := GetCurrentPath(response)

				Expect(response.StatusCode).To(Equal(http.StatusFound))
				Expect(currentPath).To(Equal("/"))
			})

			It("sets the success message", func() {
				FabricateUser("dev@nimblehq.co", "password")
				body := GenerateRequestBody(map[string]string{
					"email":    "dev@nimblehq.co",
					"password": "password",
				})
				response := MakeRequest("POST", "/sessions", body)
				flash := GetFlashMessage(response.Cookies())

				Expect(flash.Data["success"]).To(Equal("Successfully logged in"))
				Expect(flash.Data["error"]).To(BeEmpty())
			})

			It("sets user id to session", func() {
				user := FabricateUser("dev@nimblehq.co", "password")
				body := GenerateRequestBody(map[string]string{
					"email":    "dev@nimblehq.co",
					"password": "password",
				})
				response := MakeRequest("POST", "/sessions", body)
				currentUserId := GetSession(response.Cookies(), controllers.CurrentUserKey)

				Expect(currentUserId).To(Equal(user.Id))
			})
		})

		Context("given INVALID params", func() {
			Context("given NO email", func() {
				It("redirects to signin page", func() {
					FabricateUser("dev@nimblehq.co", "password")
					body := GenerateRequestBody(map[string]string{
						"email":    "",
						"password": "password",
					})
					response := MakeRequest("POST", "/sessions", body)
					currentPath := GetCurrentPath(response)

					Expect(response.StatusCode).To(Equal(http.StatusFound))
					Expect(currentPath).To(Equal("/signin"))
				})

				It("sets the error message", func() {
					FabricateUser("dev@nimblehq.co", "password")
					body := GenerateRequestBody(map[string]string{
						"email":    "",
						"password": "password",
					})
					response := MakeRequest("POST", "/sessions", body)
					flash := GetFlashMessage(response.Cookies())

					Expect(flash.Data["error"]).To(Equal("Email can not be empty"))
					Expect(flash.Data["success"]).To(BeEmpty())
				})

				It("does NOT set user id to session", func() {
					FabricateUser("dev@nimblehq.co", "password")
					body := GenerateRequestBody(map[string]string{
						"email":    "",
						"password": "password",
					})
					response := MakeRequest("POST", "/sessions", body)
					currentUserId := GetSession(response.Cookies(), controllers.CurrentUserKey)

					Expect(currentUserId).To(BeNil())
				})
			})

			Context("given NO password", func() {
				It("redirects to signin page", func() {
					FabricateUser("dev@nimblehq.co", "password")
					body := GenerateRequestBody(map[string]string{
						"email":    "dev@nimblehq.cp",
						"password": "",
					})
					response := MakeRequest("POST", "/sessions", body)
					currentPath := GetCurrentPath(response)

					Expect(response.StatusCode).To(Equal(http.StatusFound))
					Expect(currentPath).To(Equal("/signin"))
				})

				It("sets the error message", func() {
					FabricateUser("dev@nimblehq.co", "password")
					body := GenerateRequestBody(map[string]string{
						"email":    "dev@nimblehq.co",
						"password": "",
					})
					response := MakeRequest("POST", "/sessions", body)
					flash := GetFlashMessage(response.Cookies())

					Expect(flash.Data["error"]).To(Equal("Password can not be empty"))
					Expect(flash.Data["success"]).To(BeEmpty())
				})
			})

			Context("given INVALID email", func() {
				It("redirects to signin page", func() {
					FabricateUser("dev@nimblehq.co", "password")
					body := GenerateRequestBody(map[string]string{
						"email":    "invalid@email.com",
						"password": "password",
					})
					response := MakeRequest("POST", "/sessions", body)
					currentPath := GetCurrentPath(response)

					Expect(response.StatusCode).To(Equal(http.StatusFound))
					Expect(currentPath).To(Equal("/signin"))
				})

				It("sets the error message", func() {
					FabricateUser("dev@nimblehq.co", "password")
					body := GenerateRequestBody(map[string]string{
						"email":    "invalid@email.com",
						"password": "password",
					})
					response := MakeRequest("POST", "/sessions", body)
					flash := GetFlashMessage(response.Cookies())

					Expect(flash.Data["error"]).To(Equal("Incorrect email or password"))
					Expect(flash.Data["success"]).To(BeEmpty())
				})
			})

			Context("given INVALID password", func() {
				It("redirects to signin page", func() {
					FabricateUser("dev@nimblehq.co", "password")
					body := GenerateRequestBody(map[string]string{
						"email":    "dev@nimblehq.co",
						"password": "invalid password",
					})
					response := MakeRequest("POST", "/sessions", body)
					currentPath := GetCurrentPath(response)

					Expect(response.StatusCode).To(Equal(http.StatusFound))
					Expect(currentPath).To(Equal("/signin"))
				})

				It("sets the error message", func() {
					FabricateUser("dev@nimblehq.co", "password")
					body := GenerateRequestBody(map[string]string{
						"email":    "dev@nimblehq.co",
						"password": "invalid password",
					})
					response := MakeRequest("POST", "/sessions", body)
					flash := GetFlashMessage(response.Cookies())

					Expect(flash.Data["error"]).To(Equal("Incorrect email or password"))
					Expect(flash.Data["success"]).To(BeEmpty())
				})
			})
		})
	})

	AfterEach(func() {
		initializers.CleanupDatabase("users")
	})
})
