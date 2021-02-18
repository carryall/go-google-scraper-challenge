package controllers_test

import (
	"net/http"

	"go-google-scraper-challenge/initializers"
	. "go-google-scraper-challenge/test/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("KeywordController", func() {
	Describe("GET /", func() {
		Context("given user already signed in", func() {
			It("renders with status 200", func() {
				user := FabricateUser("dev@nimblehq.co", "password")
				response := MakeAuthenticatedRequest("GET", "/", nil, user)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
			})
		})

		Context("given user is NOT signed in", func() {
			It("redirects to sign in path", func() {
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