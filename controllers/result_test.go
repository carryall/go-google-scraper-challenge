package controllers_test

import (
	"net/http"

	"go-google-scraper-challenge/initializers"
	"go-google-scraper-challenge/models"
	. "go-google-scraper-challenge/tests/helpers"

	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ResultController", func() {
	Describe("GET /", func() {
		Context("given the user already signed in", func() {
			It("renders with status 200", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				response := MakeAuthenticatedRequest("GET", "/", nil, nil, user)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
			})

			Context("given the user have results", func() {
				It("renders user results", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					otherUser := FabricateUser(faker.Email(), faker.Password())
					result1 := FabricateResult(user)
					result2 := FabricateResult(user)
					result3 := FabricateResult(otherUser)

					response := MakeAuthenticatedRequest("GET", "/", nil, nil, user)
					responseBody := GetResponseBody(response)

					Expect(responseBody).To(ContainSubstring(result1.Keyword))
					Expect(responseBody).To(ContainSubstring(result2.Keyword))
					Expect(responseBody).NotTo(ContainSubstring(result3.Keyword))
				})
			})
		})

		Context("given the user is NOT signed in", func() {
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
			It("redirects to the root path", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				header, body := CreateRequestInfoFormFile("tests/fixtures/files/valid.csv")

				response := MakeAuthenticatedRequest("POST", "/results",  header, body, user)
				currentPath := GetCurrentPath(response)

				Expect(response.StatusCode).To(Equal(http.StatusFound))
				Expect(currentPath).To(Equal("/"))
			})

			It("sets the success message", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				header, body := CreateRequestInfoFormFile("tests/fixtures/files/valid.csv")

				response := MakeAuthenticatedRequest("POST", "/results", header, body, user)
				flash := GetFlashMessage(response.Cookies())

				Expect(flash.Data["success"]).To(Equal("Successfully uploaded the file, the result status would be updated soon"))
				Expect(flash.Data["error"]).To(BeEmpty())
			})

			It("creates results with the given keywords", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				header, body := CreateRequestInfoFormFile("tests/fixtures/files/valid.csv")

				MakeAuthenticatedRequest("POST", "/results", header, body, user)

				results, err := models.GetResultsByUserId(user.Id)
				if err != nil {
					Fail("Failed to get user results: " + err.Error())
				}

				for _, r := range results {
					Expect(r.Keyword).To(SatisfyAny(Equal("cloud computing service"), Equal("crypto currency")))
					Expect(r.Status).To(Equal(models.ResultStatusCompleted))
					Expect(r.PageCache).NotTo(BeEmpty())
				}
			})
		})

		Context("given a blank CSV file", func() {
			It("redirects to the root path", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				header, body := CreateRequestInfoFormFile("tests/fixtures/files/empty.csv")

				response := MakeAuthenticatedRequest("POST", "/results",  header, body, user)
				currentPath := GetCurrentPath(response)

				Expect(response.StatusCode).To(Equal(http.StatusFound))
				Expect(currentPath).To(Equal("/"))
			})

			It("sets the error message", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				header, body := CreateRequestInfoFormFile("tests/fixtures/files/empty.csv")

				response := MakeAuthenticatedRequest("POST", "/results", header, body, user)
				flash := GetFlashMessage(response.Cookies())

				Expect(flash.Data["success"]).To(BeEmpty())
				Expect(flash.Data["error"]).To(Equal("File should contains between 1 to 1000 keywords"))
			})
		})

		Context("given a CSV file that contains more than 1000 keywords", func() {
			It("redirects to the root path", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				header, body := CreateRequestInfoFormFile("tests/fixtures/files/invalid.csv")

				response := MakeAuthenticatedRequest("POST", "/results",  header, body, user)
				currentPath := GetCurrentPath(response)

				Expect(response.StatusCode).To(Equal(http.StatusFound))
				Expect(currentPath).To(Equal("/"))
			})

			It("sets the error message", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				header, body := CreateRequestInfoFormFile("tests/fixtures/files/invalid.csv")

				response := MakeAuthenticatedRequest("POST", "/results", header, body, user)
				flash := GetFlashMessage(response.Cookies())

				Expect(flash.Data["success"]).To(BeEmpty())
				Expect(flash.Data["error"]).To(Equal("File should contains between 1 to 1000 keywords"))
			})
		})

		Context("given an INVALID file type", func() {
			It("redirects to the root path", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				header, body := CreateRequestInfoFormFile("tests/fixtures/files/text.txt")

				response := MakeAuthenticatedRequest("POST", "/results",  header, body, user)
				currentPath := GetCurrentPath(response)

				Expect(response.StatusCode).To(Equal(http.StatusFound))
				Expect(currentPath).To(Equal("/"))
			})

			It("sets the error message", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				header, body := CreateRequestInfoFormFile("tests/fixtures/files/text.txt")

				response := MakeAuthenticatedRequest("POST", "/results", header, body, user)
				flash := GetFlashMessage(response.Cookies())

				Expect(flash.Data["success"]).To(BeEmpty())
				Expect(flash.Data["error"]).To(Equal("Incorrect file type"))
			})
		})

		Context("given user is NOT signed in", func() {
			It("returns an error", func() {
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
