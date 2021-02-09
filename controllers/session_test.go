package controllers_test

import (
	"net/http"

	"go-google-scraper-challenge/initializers"
	. "go-google-scraper-challenge/test/helpers"
	. "go-google-scraper-challenge/test/serializers"

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
			It("returns with status ok", func() {
				FabricateUser("dev@nimblehq.co", "password")
				body := RequestBody(map[string]string{
					"username": "dev@nimblehq.co",
					"password": "password",
				})
				response := MakeRequest("POST", "/sessions", body)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
			})

			It("returns with a valid token", func() {
				FabricateUser("dev@nimblehq.co", "password")
				body := RequestBody(map[string]string{
					"username": "dev@nimblehq.co",
					"password": "password",
				})
				response := MakeRequest("POST", "/sessions", body)
				responseBody := LoginResponse{}
				GetJSONResponseBody(response, &responseBody)

				Expect(responseBody.AccessToken).NotTo(BeNil())
				Expect(responseBody.RefreshToken).NotTo(BeNil())
				Expect(responseBody.ExpiresIn).NotTo(BeNil())
				Expect(responseBody.TokenType).NotTo(BeNil())
			})
		})

		Context("given INVALID params", func() {
			Context("given NO username", func() {
				It("returns with a bad request error", func() {
					FabricateUser("dev@nimblehq.co", "password")
					body := RequestBody(map[string]string{
						"username": "",
						"password": "password",
					})
					response := MakeRequest("POST", "/sessions", body)
					responseBody := ErrorResponse{}
					GetJSONResponseBody(response, &responseBody)

					Expect(response.StatusCode).To(Equal(http.StatusBadRequest))
					Expect(responseBody.Error).To(Equal("Bad Request"))
					Expect(responseBody.ErrorDescription).To(Equal("The request is missing a required parameter, includes an invalid parameter value, includes a parameter more than once, or is otherwise malformed"))
				})
			})

			Context("given NO password", func() {
				It("returns with a bad request error", func() {
					FabricateUser("dev@nimblehq.co", "password")
					body := RequestBody(map[string]string{
						"username": "dev@nimblehq.co",
						"password": "",
					})
					response := MakeRequest("POST", "/sessions", body)
					responseBody := ErrorResponse{}
					GetJSONResponseBody(response, &responseBody)

					Expect(response.StatusCode).To(Equal(http.StatusBadRequest))
					Expect(responseBody.Error).To(Equal("Bad Request"))
					Expect(responseBody.ErrorDescription).To(Equal("The request is missing a required parameter, includes an invalid parameter value, includes a parameter more than once, or is otherwise malformed"))
				})
			})

			Context("given INVALID username", func() {
				It("returns with an invalid client error", func() {
					FabricateUser("dev@nimblehq.co", "password")
					body := RequestBody(map[string]string{
						"username": "invalid@email.com",
						"password": "password",
					})
					response := MakeRequest("POST", "/sessions", body)
					responseBody := ErrorResponse{}
					GetJSONResponseBody(response, &responseBody)

					Expect(response.StatusCode).To(Equal(http.StatusUnauthorized))
					Expect(responseBody.Error).To(Equal("invalid_client"))
					Expect(responseBody.ErrorDescription).To(Equal("Client authentication failed"))
				})
			})

			Context("given INVALID password", func() {
				It("returns with an invalid client error", func() {
					FabricateUser("dev@nimblehq.co", "password")
					body := RequestBody(map[string]string{
						"username": "dev@nimblehq.co",
						"password": "invalid password",
					})
					response := MakeRequest("POST", "/sessions", body)
					responseBody := ErrorResponse{}
					GetJSONResponseBody(response, &responseBody)

					Expect(response.StatusCode).To(Equal(http.StatusUnauthorized))
					Expect(responseBody.Error).To(Equal("invalid_client"))
					Expect(responseBody.ErrorDescription).To(Equal("Client authentication failed"))
				})
			})
		})
	})

	AfterEach(func() {
		initializers.CleanupDatabase("user")
		initializers.CleanupDatabase("oauth2_clients")
		initializers.CleanupDatabase("oauth2_tokens")
	})
})
