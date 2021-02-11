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
		Context("given user is not signed in", func() {
			It("renders with status 200", func() {
				response := MakeRequest("GET", "/signin", nil)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
			})
		})

		Context("given user already signed in", func() {
			It("redirects to root path", func() {
				user := FabricateUser("dev@nimblehq.co", "password")
				response := MakeAuthenticatedRequest("GET", "/signin", nil, user)
				currentPath := GetCurrentPath(response)

				Expect(response.StatusCode).To(Equal(http.StatusFound))
				Expect(currentPath).To(Equal("/"))
			})
		})
	})

	Describe("POST /sessions", func() {
		Context("given user not signed in", func() {
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

					Expect(flash.Data["success"]).To(Equal("Successfully signed in"))
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

					It("does NOT set user id to session", func() {
						FabricateUser("dev@nimblehq.co", "password")
						body := GenerateRequestBody(map[string]string{
							"email":    "dev@nimblehq.co",
							"password": "",
						})
						response := MakeRequest("POST", "/sessions", body)
						currentUserId := GetSession(response.Cookies(), controllers.CurrentUserKey)

						Expect(currentUserId).To(BeNil())
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

					It("does NOT set user id to session", func() {
						FabricateUser("dev@nimblehq.co", "password")
						body := GenerateRequestBody(map[string]string{
							"email":    "invalid@email.com",
							"password": "password",
						})
						response := MakeRequest("POST", "/sessions", body)
						currentUserId := GetSession(response.Cookies(), controllers.CurrentUserKey)

						Expect(currentUserId).To(BeNil())
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

					It("does NOT set user id to session", func() {
						FabricateUser("dev@nimblehq.co", "password")
						body := GenerateRequestBody(map[string]string{
							"email":    "dev@nimblehq.co",
							"password": "invalid password",
						})
						response := MakeRequest("POST", "/sessions", body)
						currentUserId := GetSession(response.Cookies(), controllers.CurrentUserKey)

						Expect(currentUserId).To(BeNil())
					})
				})
			})
		})

		Context("given user is already signed in", func() {
			It("returns error", func() {
				user := FabricateUser("dev@nimblehq.co", "password")
				body := GenerateRequestBody(map[string]string{
					"email":    "dev@nimblehq.co",
					"password": "password",
				})
				response := MakeAuthenticatedRequest("POST", "/sessions", body, user)

				Expect(response.StatusCode).To(Equal(http.StatusMethodNotAllowed))
			})
		})
	})

	Describe("DELETE /sessions", func() {
		Context("given user is already signed in", func() {
			It("redirects to sign in path", func() {
				user := FabricateUser("dev@nimblehq.co", "password")
				body := GenerateRequestBody(map[string]string{})
				response := MakeAuthenticatedRequest("DELETE", "/sessions", body, user)
				currentPath := GetCurrentPath(response)

				Expect(response.StatusCode).To(Equal(http.StatusFound))
				Expect(currentPath).To(Equal("/signin"))
			})

			It("sets the success message", func() {
				user := FabricateUser("dev@nimblehq.co", "password")
				body := GenerateRequestBody(map[string]string{})
				response := MakeAuthenticatedRequest("DELETE", "/sessions", body, user)
				flash := GetFlashMessage(response.Cookies())

				Expect(flash.Data["success"]).To(Equal("Successfully signed out"))
				Expect(flash.Data["error"]).To(BeEmpty())
			})

			It("removes user id from session", func() {
				user := FabricateUser("dev@nimblehq.co", "password")
				body := GenerateRequestBody(map[string]string{})
				response := MakeAuthenticatedRequest("DELETE", "/sessions", body, user)
				currentUserId := GetSession(response.Cookies(), controllers.CurrentUserKey)

				Expect(currentUserId).To(BeNil())
			})
		})

		Context("given user is NOT signed in", func() {
			It("returns error", func() {
				body := GenerateRequestBody(map[string]string{})
				response := MakeRequest("DELETE", "/sessions", body)

				Expect(response.StatusCode).To(Equal(http.StatusMethodNotAllowed))
			})
		})
	})

	AfterEach(func() {
		initializers.CleanupDatabase("users")
	})
})
