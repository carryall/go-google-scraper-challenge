package controllers_test

import (
	"net/http"

	"go-google-scraper-challenge/initializers"
	. "go-google-scraper-challenge/tests/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ResultController", func() {
	Describe("GET /", func() {
		Context("given user already signed in", func() {
			It("renders with status 200", func() {
				user := FabricateUser("dev@nimblehq.co", "password")
				response := MakeAuthenticatedRequest("GET", "/", nil, nil, user)

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

	Describe("POST /results", func() {
		Context("given a valid CSV file", func() {
			It("redirects to root path", func() {
				user := FabricateUser("dev@nimblehq.co", "password")
				header, body := CreateRequestInfoFormFile("tests/fixtures/files/valid.csv")

				response := MakeAuthenticatedRequest("POST", "/results",  header, body, user)
				currentPath := GetCurrentPath(response)

				Expect(response.StatusCode).To(Equal(http.StatusFound))
				Expect(currentPath).To(Equal("/"))
			})

			It("sets the success message", func() {
				user := FabricateUser("dev@nimblehq.co", "password")
				header, body := CreateRequestInfoFormFile("tests/fixtures/files/valid.csv")

				response := MakeAuthenticatedRequest("POST", "/results", header, body, user)
				flash := GetFlashMessage(response.Cookies())

				Expect(flash.Data["success"]).To(Equal("Successfully uploaded the file, the result status would be update soon"))
				Expect(flash.Data["error"]).To(BeEmpty())
			})
		})

		Context("given a CSV file with more than 1000 keywords", func() {
			It("redirects to root path", func() {
				user := FabricateUser("dev@nimblehq.co", "password")
				header, body := CreateRequestInfoFormFile("tests/fixtures/files/invalid.csv")

				response := MakeAuthenticatedRequest("POST", "/results",  header, body, user)
				currentPath := GetCurrentPath(response)

				Expect(response.StatusCode).To(Equal(http.StatusFound))
				Expect(currentPath).To(Equal("/"))
			})

			It("sets the error message", func() {
				user := FabricateUser("dev@nimblehq.co", "password")
				header, body := CreateRequestInfoFormFile("tests/fixtures/files/invalid.csv")

				response := MakeAuthenticatedRequest("POST", "/results", header, body, user)
				flash := GetFlashMessage(response.Cookies())

				Expect(flash.Data["success"]).To(BeEmpty())
				Expect(flash.Data["error"]).To(Equal("File contains too many keywords"))
			})
		})

		Context("given an INVALID file type", func() {
			It("redirects to root path", func() {
				user := FabricateUser("dev@nimblehq.co", "password")
				header, body := CreateRequestInfoFormFile("tests/fixtures/files/text.txt")

				response := MakeAuthenticatedRequest("POST", "/results",  header, body, user)
				currentPath := GetCurrentPath(response)

				Expect(response.StatusCode).To(Equal(http.StatusFound))
				Expect(currentPath).To(Equal("/"))
			})

			It("sets the error message", func() {
				user := FabricateUser("dev@nimblehq.co", "password")
				header, body := CreateRequestInfoFormFile("tests/fixtures/files/text.txt")

				response := MakeAuthenticatedRequest("POST", "/results", header, body, user)
				flash := GetFlashMessage(response.Cookies())

				Expect(flash.Data["success"]).To(BeEmpty())
				Expect(flash.Data["error"]).To(Equal("Incorrect file type"))
			})
		})

		Context("given user is NOT signed in", func() {
			It("returns error", func() {
				body := GenerateRequestBody(nil)
				response := MakeRequest("POST", "/results", body)

				Expect(response.StatusCode).To(Equal(http.StatusMethodNotAllowed))
			})
		})
	})

	AfterEach(func() {
		initializers.CleanupDatabase([]string{"users", "results", "links", "ad_links"})
	})
})
