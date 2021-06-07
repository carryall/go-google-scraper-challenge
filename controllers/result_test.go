package controllers_test

import (
	"fmt"
	"net/http"
	"net/url"

	"go-google-scraper-challenge/constants"
	"go-google-scraper-challenge/initializers"
	"go-google-scraper-challenge/models"
	. "go-google-scraper-challenge/tests"
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

			Context("given the user has results", func() {
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
			BeforeEach(func() {
				keyword := "ergonomic chair"
				visitURL := fmt.Sprintf("http://www.google.com/search?q=%s", url.QueryEscape(keyword))
				cassetteName := "scraper/success"

				RecordResponse(cassetteName, visitURL)
			})

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

				var query = map[string]interface{}{
					"user_id": user.Id,
				}
				results, err := models.GetResultsBy(query)
				if err != nil {
					Fail("Failed to get user results: " + err.Error())
				}

				Expect(results).To(HaveLen(1))
				Expect(results[0].Keyword).To(Equal("ergonomic chair"))
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
				Expect(flash.Data["error"]).To(Equal(constants.FileTypeInvalid))
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

	Describe("GET /results/:id", func() {
		Context("given the user already signed in", func() {
			Context("given a valid result id", func() {
				It("renders with status 200", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					result := FabricateResultWithParams(user, "some specific keyword", models.ResultStatusCompleted)

					response := MakeAuthenticatedRequest("GET", fmt.Sprintf("/results/%d", result.Id), nil, nil, user)

					Expect(response.StatusCode).To(Equal(http.StatusOK))
				})

				It("display result information", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					result := FabricateResultWithParams(user, "some specific keyword", models.ResultStatusCompleted)
					adLink1 := FabricateAdLink(result)
					adLink2 := FabricateAdLink(result)
					link := FabricateLink(result)

					response := MakeAuthenticatedRequest("GET", fmt.Sprintf("/results/%d", result.Id), nil, nil, user)
					responseBody := GetResponseBody(response)

					Expect(responseBody).To(ContainSubstring("some specific keyword"))
					Expect(responseBody).To(ContainSubstring("completed"))
					Expect(responseBody).To(ContainSubstring(adLink1.Link))
					Expect(responseBody).To(ContainSubstring(adLink2.Link))
					Expect(responseBody).To(ContainSubstring(link.Link))
				})
			})

			Context("given an INVALID result id", func() {
				It("display an error message with a link to root path", func() {
					user := FabricateUser(faker.Email(), faker.Password())

					response := MakeAuthenticatedRequest("GET", "/results/9999", nil, nil, user)
					responseBody := GetResponseBody(response)

					Expect(responseBody).To(ContainSubstring("We cannot find the result you are looking for"))
					Expect(responseBody).To(ContainSubstring("Back to Home"))
				})
			})
		})

		Context("given the user is NOT signed in", func() {
			It("redirects to sign in path", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				result := FabricateResult(user)

				response := MakeRequest("GET", fmt.Sprintf("/results/%d", result.Id), nil)
				currentPath := GetCurrentPath(response)

				Expect(response.StatusCode).To(Equal(http.StatusFound))
				Expect(currentPath).To(Equal("/signin"))
			})
		})
	})

	Describe("GET /results/:id/cache", func() {
		Context("given the user already signed in", func() {
			Context("given a valid result id", func() {
				It("renders with status 200", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					result := FabricateResultWithParams(user, "some specific keyword", models.ResultStatusCompleted)

					response := MakeAuthenticatedRequest("GET", fmt.Sprintf("/results/%d/cache", result.Id), nil, nil, user)

					Expect(response.StatusCode).To(Equal(http.StatusOK))
				})

				It("display result page cache", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					result := FabricateResultWithParams(user, "some specific keyword", models.ResultStatusCompleted)
					result.PageCache = "the page cache"
					err := models.UpdateResultById(result)
					if err != nil {
						Fail("Failed to update result by id")
					}

					response := MakeAuthenticatedRequest("GET", fmt.Sprintf("/results/%d/cache", result.Id), nil, nil, user)
					responseBody := GetResponseBody(response)

					Expect(responseBody).To(ContainSubstring("the page cache"))
				})
			})

			Context("given an INVALID result id", func() {
				It("display an error message with a link to root path", func() {
					user := FabricateUser(faker.Email(), faker.Password())

					response := MakeAuthenticatedRequest("GET", "/results/9999/cache", nil, nil, user)
					responseBody := GetResponseBody(response)

					Expect(responseBody).To(ContainSubstring("We cannot find the result you are looking for"))
					Expect(responseBody).To(ContainSubstring("Back to Home"))
				})
			})
		})

		Context("given the user is NOT signed in", func() {
			It("redirects to sign in path", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				result := FabricateResult(user)

				response := MakeRequest("GET", fmt.Sprintf("/results/%d/cache", result.Id), nil)
				currentPath := GetCurrentPath(response)

				Expect(response.StatusCode).To(Equal(http.StatusFound))
				Expect(currentPath).To(Equal("/signin"))
			})
		})
	})

	AfterEach(func() {
		initializers.CleanupDatabase([]string{"users", "results", "links", "ad_links"})
	})
})
