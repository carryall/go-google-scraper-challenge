package webcontrollers_test

import (
	"net/http"
	"net/url"
	"strconv"

	"go-google-scraper-challenge/constants"
	. "go-google-scraper-challenge/test"

	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UsersController", func() {
	Describe("GET /signup", func() {
		Context("given user is not signed in", func() {
			It("renders with status 200", func() {
				response := MakeWebRequest("GET", "/signup", nil, nil)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
			})
		})

		// TODO: Test this when work on the result list screen
		XContext("given user already signed in", func() {
			It("redirects to root path", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				response := MakeWebRequest("GET", "/signup", nil, user)

				currentPath := GetCurrentPath(response)

				Expect(response.StatusCode).To(Equal(http.StatusFound))
				Expect(currentPath).To(Equal("/"))
			})
		})
	})

	Describe("POST /users", func() {
		Context("given user not signed in", func() {
			Context("given valid params", func() {
				It("redirects to root path", func() {
					password := faker.Password()
					formData := url.Values{
						"email":                 {faker.Email()},
						"password":              {password},
						"password_confirmation": {password},
					}
					response := MakeWebFormRequest("POST", "/users", formData, nil)
					currentPath := GetCurrentPath(response)

					Expect(response.StatusCode).To(Equal(http.StatusFound))
					Expect(currentPath).To(Equal("/"))
				})

				It("sets the success message", func() {
					password := faker.Password()
					formData := url.Values{
						"email":                 {faker.Email()},
						"password":              {password},
						"password_confirmation": {password},
					}
					response := MakeWebFormRequest("POST", "/users", formData, nil)
					flashes := GetFlashMessage(response.Cookies())

					Expect(flashes["success"]).To(ContainElement("Successfully signed up"))
					Expect(flashes["error"]).To(BeEmpty())
				})

				It("sets user id to session", func() {
					password := faker.Password()
					formData := url.Values{
						"email":                 {faker.Email()},
						"password":              {password},
						"password_confirmation": {password},
					}
					response := MakeWebFormRequest("POST", "/users", formData, nil)
					currentUserIdStr := GetSessionUserID(response.Cookies())
					currentUserId, err := strconv.Atoi(currentUserIdStr.(string))
					if err != nil {
						Fail("Fail to convert current user ID")
					}

					Expect(currentUserId).To(BeNumerically(">", 0))
				})
			})

			Context("given INVALID params", func() {
				Context("given NO email", func() {
					It("redirects to signup page", func() {
						password := faker.Password()
						FabricateUser(faker.Email(), password)
						formData := url.Values{
							"email":                 {""},
							"password":              {password},
							"password_confirmation": {password},
						}
						response := MakeWebFormRequest("POST", "/users", formData, nil)
						currentPath := GetCurrentPath(response)

						Expect(response.StatusCode).To(Equal(http.StatusFound))
						Expect(currentPath).To(Equal("/signup"))
					})

					It("sets the error message", func() {
						password := faker.Password()
						FabricateUser(faker.Email(), password)
						formData := url.Values{
							"email":                 {""},
							"password":              {password},
							"password_confirmation": {password},
						}
						response := MakeWebFormRequest("POST", "/users", formData, nil)
						flashes := GetFlashMessage(response.Cookies())

						Expect(flashes["error"]).To(ContainElement("Email: cannot be blank."))
						Expect(flashes["success"]).To(BeEmpty())
					})

					It("does NOT set user id to session", func() {
						password := faker.Password()
						FabricateUser(faker.Email(), password)
						formData := url.Values{
							"email":                 {""},
							"password":              {password},
							"password_confirmation": {password},
						}
						response := MakeWebFormRequest("POST", "/users", formData, nil)
						currentUserId := GetSessionUserID(response.Cookies())

						Expect(currentUserId).To(BeNil())
					})
				})

				Context("given NO password", func() {
					It("redirects to signup page", func() {
						user := FabricateUser(faker.Email(), faker.Password())
						formData := url.Values{
							"email":                 {user.Email},
							"password":              {""},
							"password_confirmation": {faker.Password()},
						}
						response := MakeWebFormRequest("POST", "/users", formData, nil)
						currentPath := GetCurrentPath(response)

						Expect(response.StatusCode).To(Equal(http.StatusFound))
						Expect(currentPath).To(Equal("/signup"))
					})

					It("sets the error message", func() {
						user := FabricateUser(faker.Email(), faker.Password())
						formData := url.Values{
							"email":                 {user.Email},
							"password":              {""},
							"password_confirmation": {faker.Password()},
						}
						response := MakeWebFormRequest("POST", "/users", formData, nil)
						flashes := GetFlashMessage(response.Cookies())

						Expect(flashes["error"]).To(ContainElement("Password: cannot be blank."))
						Expect(flashes["success"]).To(BeEmpty())
					})

					It("does NOT set user id to session", func() {
						user := FabricateUser(faker.Email(), faker.Password())
						formData := url.Values{
							"email":                 {user.Email},
							"password":              {""},
							"password_confirmation": {faker.Password()},
						}
						response := MakeWebFormRequest("POST", "/users", formData, nil)
						currentUserId := GetSessionUserID(response.Cookies())

						Expect(currentUserId).To(BeNil())
					})
				})

				Context("given password confirmation does NOT match the password", func() {
					It("redirects to signup page", func() {
						formData := url.Values{
							"email":                 {faker.Email()},
							"password":              {faker.Password()},
							"password_confirmation": {faker.Password()},
						}
						response := MakeWebFormRequest("POST", "/users", formData, nil)
						currentPath := GetCurrentPath(response)

						Expect(response.StatusCode).To(Equal(http.StatusFound))
						Expect(currentPath).To(Equal("/signup"))
					})

					It("sets the error message", func() {
						formData := url.Values{
							"email":                 {faker.Email()},
							"password":              {faker.Password()},
							"password_confirmation": {faker.Password()},
						}
						response := MakeWebFormRequest("POST", "/users", formData, nil)
						flashes := GetFlashMessage(response.Cookies())

						Expect(flashes["error"]).To(ContainElement("PasswordConfirmation: does not match the password."))
						Expect(flashes["success"]).To(BeEmpty())
					})

					It("does NOT set user id to session", func() {
						formData := url.Values{
							"email":                 {faker.Email()},
							"password":              {faker.Password()},
							"password_confirmation": {faker.Password()},
						}
						response := MakeWebFormRequest("POST", "/users", formData, nil)
						currentUserId := GetSessionUserID(response.Cookies())

						Expect(currentUserId).To(BeNil())
					})
				})

				Context("given email is already exist", func() {
					It("redirects to signup page", func() {
						user := FabricateUser(faker.Email(), faker.Password())
						password := faker.Password()
						formData := url.Values{
							"email":                 {user.Email},
							"password":              {password},
							"password_confirmation": {password},
						}
						response := MakeWebFormRequest("POST", "/users", formData, nil)
						currentPath := GetCurrentPath(response)

						Expect(response.StatusCode).To(Equal(http.StatusFound))
						Expect(currentPath).To(Equal("/signup"))
					})

					It("sets the error message", func() {
						user := FabricateUser(faker.Email(), faker.Password())
						password := faker.Password()
						formData := url.Values{
							"email":                 {user.Email},
							"password":              {password},
							"password_confirmation": {password},
						}
						response := MakeWebFormRequest("POST", "/users", formData, nil)
						flashes := GetFlashMessage(response.Cookies())

						Expect(flashes["error"]).To(ContainElement(constants.UserAlreadyExist))
						Expect(flashes["success"]).To(BeEmpty())
					})

					It("does NOT set user id to session", func() {
						user := FabricateUser(faker.Email(), faker.Password())
						password := faker.Password()
						formData := url.Values{
							"email":                 {user.Email},
							"password":              {password},
							"password_confirmation": {password},
						}
						response := MakeWebFormRequest("POST", "/users", formData, nil)
						currentUserId := GetSessionUserID(response.Cookies())

						Expect(currentUserId).To(BeNil())
					})
				})
			})
		})

		Context("given user is already signed in", func() {
			It("redirects to root path", func() {
				password := faker.Password()
				user := FabricateUser(faker.Email(), password)
				formData := url.Values{
					"email":                 {faker.Email()},
					"password":              {password},
					"password_confirmation": {password},
				}
				response := MakeWebFormRequest("POST", "/users", formData, user)
				currentPath := GetCurrentPath(response)

				Expect(response.StatusCode).To(Equal(http.StatusFound))
				Expect(currentPath).To(Equal("/"))
			})
		})
	})

	AfterEach(func() {
		CleanupDatabase([]string{"users"})
	})
})
