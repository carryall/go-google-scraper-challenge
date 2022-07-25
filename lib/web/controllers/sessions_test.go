package webcontrollers_test

import (
	"fmt"
	"net/http"
	"net/url"

	"go-google-scraper-challenge/constants"
	. "go-google-scraper-challenge/test"

	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SessionsController", func() {
	Describe("GET /signin", func() {
		Context("given user is not signed in", func() {
			It("renders with status 200", func() {
				response := MakeWebRequest("GET", "/signin", nil, nil)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
			})
		})

		Context("given user already signed in", func() {
			It("redirects to root path", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				response := MakeWebRequest("GET", "/signin", nil, user)

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
					password := faker.Password()
					user := FabricateUser(faker.Email(), password)
					formData := url.Values{
						"email":    {user.Email},
						"password": {password},
					}
					response := MakeWebFormRequest("POST", "/sessions", formData, nil)
					currentPath := GetCurrentPath(response)

					Expect(response.StatusCode).To(Equal(http.StatusFound))
					Expect(currentPath).To(Equal("/"))
				})

				It("sets the success message", func() {
					password := faker.Password()
					user := FabricateUser(faker.Email(), password)
					formData := url.Values{
						"email":    {user.Email},
						"password": {password},
					}
					response := MakeWebFormRequest("POST", "/sessions", formData, nil)
					flashes := GetFlashMessage(response.Cookies())

					Expect(flashes["success"]).To(ContainElement("Successfully signed in"))
					Expect(flashes["error"]).To(BeEmpty())
				})

				It("sets user id to session", func() {
					password := faker.Password()
					user := FabricateUser(faker.Email(), password)
					formData := url.Values{
						"email":    {user.Email},
						"password": {password},
					}
					response := MakeWebFormRequest("POST", "/sessions", formData, nil)
					currentUserId := GetSessionUserID(response.Cookies())

					Expect(currentUserId).To(Equal(fmt.Sprint(user.ID)))
				})
			})

			Context("given INVALID params", func() {
				Context("given NO email", func() {
					It("redirects to signin page", func() {
						password := faker.Password()
						FabricateUser(faker.Email(), password)
						formData := url.Values{
							"email":    {""},
							"password": {password},
						}
						response := MakeWebFormRequest("POST", "/sessions", formData, nil)
						currentPath := GetCurrentPath(response)

						Expect(response.StatusCode).To(Equal(http.StatusFound))
						Expect(currentPath).To(Equal("/signin"))
					})

					It("sets the error message", func() {
						password := faker.Password()
						FabricateUser(faker.Email(), password)
						formData := url.Values{
							"email":    {""},
							"password": {password},
						}
						response := MakeWebFormRequest("POST", "/sessions", formData, nil)
						flashes := GetFlashMessage(response.Cookies())

						Expect(flashes["error"]).To(ContainElement("Email: cannot be blank."))
						Expect(flashes["success"]).To(BeEmpty())
					})

					It("does NOT set user id to session", func() {
						password := faker.Password()
						FabricateUser(faker.Email(), password)
						formData := url.Values{
							"email":    {""},
							"password": {password},
						}
						response := MakeWebFormRequest("POST", "/sessions", formData, nil)
						currentUserId := GetSessionUserID(response.Cookies())

						Expect(currentUserId).To(BeNil())
					})
				})

				Context("given NO password", func() {
					It("redirects to signin page", func() {
						user := FabricateUser(faker.Email(), faker.Password())
						formData := url.Values{
							"email":    {user.Email},
							"password": {""},
						}
						response := MakeWebFormRequest("POST", "/sessions", formData, nil)
						currentPath := GetCurrentPath(response)

						Expect(response.StatusCode).To(Equal(http.StatusFound))
						Expect(currentPath).To(Equal("/signin"))
					})

					It("sets the error message", func() {
						user := FabricateUser(faker.Email(), faker.Password())
						formData := url.Values{
							"email":    {user.Email},
							"password": {""},
						}
						response := MakeWebFormRequest("POST", "/sessions", formData, nil)
						flashes := GetFlashMessage(response.Cookies())

						Expect(flashes["error"]).To(ContainElement("Password: cannot be blank."))
						Expect(flashes["success"]).To(BeEmpty())
					})

					It("does NOT set user id to session", func() {
						user := FabricateUser(faker.Email(), faker.Password())
						formData := url.Values{
							"email":    {user.Email},
							"password": {""},
						}
						response := MakeWebFormRequest("POST", "/sessions", formData, nil)
						currentUserId := GetSessionUserID(response.Cookies())

						Expect(currentUserId).To(BeNil())
					})
				})

				Context("given INVALID email", func() {
					It("redirects to signin page", func() {
						FabricateUser(faker.Email(), faker.Password())
						formData := url.Values{
							"email":    {faker.Email()},
							"password": {faker.Password()},
						}
						response := MakeWebFormRequest("POST", "/sessions", formData, nil)
						currentPath := GetCurrentPath(response)

						Expect(response.StatusCode).To(Equal(http.StatusFound))
						Expect(currentPath).To(Equal("/signin"))
					})

					It("sets the error message", func() {
						FabricateUser(faker.Email(), faker.Password())
						formData := url.Values{
							"email":    {faker.Email()},
							"password": {faker.Password()},
						}
						response := MakeWebFormRequest("POST", "/sessions", formData, nil)
						flashes := GetFlashMessage(response.Cookies())

						Expect(flashes["error"]).To(ContainElement(constants.SignInFail))
						Expect(flashes["success"]).To(BeEmpty())
					})

					It("does NOT set user id to session", func() {
						FabricateUser(faker.Email(), faker.Password())
						formData := url.Values{
							"email":    {faker.Email()},
							"password": {faker.Password()},
						}
						response := MakeWebFormRequest("POST", "/sessions", formData, nil)
						currentUserId := GetSessionUserID(response.Cookies())

						Expect(currentUserId).To(BeNil())
					})
				})

				Context("given INVALID password", func() {
					It("redirects to signin page", func() {
						user := FabricateUser(faker.Email(), faker.Password())
						formData := url.Values{
							"email":    {user.Email},
							"password": {faker.Password()},
						}
						response := MakeWebFormRequest("POST", "/sessions", formData, nil)
						currentPath := GetCurrentPath(response)

						Expect(response.StatusCode).To(Equal(http.StatusFound))
						Expect(currentPath).To(Equal("/signin"))
					})

					It("sets the error message", func() {
						user := FabricateUser(faker.Email(), faker.Password())
						formData := url.Values{
							"email":    {user.Email},
							"password": {faker.Password()},
						}
						response := MakeWebFormRequest("POST", "/sessions", formData, nil)
						flashes := GetFlashMessage(response.Cookies())

						Expect(flashes["error"]).To(ContainElement(constants.SignInFail))
						Expect(flashes["success"]).To(BeEmpty())
					})

					It("does NOT set user id to session", func() {
						user := FabricateUser(faker.Email(), faker.Password())
						formData := url.Values{
							"email":    {user.Email},
							"password": {faker.Password()},
						}
						response := MakeWebFormRequest("POST", "/sessions", formData, nil)
						currentUserId := GetSessionUserID(response.Cookies())

						Expect(currentUserId).To(BeNil())
					})
				})
			})
		})

		Context("given a user is already signed in", func() {
			It("redirects to root path", func() {
				password := faker.Password()
				user := FabricateUser(faker.Email(), password)
				formData := url.Values{
					"email":    {user.Email},
					"password": {password},
				}
				response := MakeWebFormRequest("POST", "/sessions", formData, user)
				currentPath := GetCurrentPath(response)

				Expect(response.StatusCode).To(Equal(http.StatusFound))
				Expect(currentPath).To(Equal("/"))
			})
		})
	})

	Describe("POST /signout", func() {
		Context("given a user is already signed in", func() {
			It("redirects to sign in screen", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				response := MakeWebFormRequest("POST", "/signout", nil, user)
				currentPath := GetCurrentPath(response)

				Expect(response.StatusCode).To(Equal(http.StatusFound))
				Expect(currentPath).To(Equal(constants.WebRoutes["sessions"]["new"]))
			})
		})
	})

	AfterEach(func() {
		CleanupDatabase([]string{"users"})
	})
})
