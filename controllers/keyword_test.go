package controllers_test

import (
	"go-google-scraper-challenge/initializers"
	. "go-google-scraper-challenge/test/helpers"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("KeywordController", func() {
	Describe("GET /", func() {
		Context("given user is already logged in", func() {
			It("renders with status 200", func() {
				FabricateUser("dev@nimblehq.co", "password")
				Login("dev@nimblehq.co", "password")

				response := MakeRequest("GET", "/", nil)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
			})
		})

		Context("given NO user logged in", func() {
			It("redirects to signin path", func() {
				response := MakeRequest("GET", "/", nil)
				currentPath := GetCurrentPath(response)

				Expect(response.StatusCode).To(Equal(http.StatusFound))
				Expect(currentPath).To(Equal("/signin"))
			})
		})
	})

	AfterEach(func() {
		initializers.CleanupDatabase("users")
	})
})
