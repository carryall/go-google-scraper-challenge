package webcontrollers_test

import (
	"fmt"
	"net/http"

	"go-google-scraper-challenge/constants"
	"go-google-scraper-challenge/lib/models"
	"go-google-scraper-challenge/lib/sessions"
	. "go-google-scraper-challenge/test"

	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ResultsController", func() {
	Describe("GET /", func() {
		Context("given a user is already signed in", func() {
			It("renders with status 200", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				response := MakeWebRequest("GET", "/", nil, nil, user)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
			})
		})

		Context("given NO user is signed in", func() {
			It("redirects to signin path", func() {
				response := MakeWebRequest("GET", "/", nil, nil, nil)

				currentPath := GetCurrentPath(response)

				Expect(response.StatusCode).To(Equal(http.StatusFound))
				Expect(currentPath).To(Equal("/signin"))
			})
		})
	})

	Describe("GET /results/:id", func() {
		Context("given a valid result ID", func() {
			It("renders the result details", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				result := FabricateResult(user)
				topAdLink1 := FabricateAdLinkWithParams(result, models.AdLinkPositionTop)
				topAdLink2 := FabricateAdLinkWithParams(result, models.AdLinkPositionTop)
				sideAdLink := FabricateAdLinkWithParams(result, models.AdLinkPositionSide)
				link1 := FabricateLink(result)
				link2 := FabricateLink(result)

				response := MakeWebRequest("GET", fmt.Sprintf(`/results/%d`, result.ID), nil, nil, user)
				responseBody := GetResponseBody(response)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
				Expect(responseBody).To(ContainSubstring(result.Keyword))
				Expect(responseBody).To(ContainSubstring(topAdLink1.Link))
				Expect(responseBody).To(ContainSubstring(topAdLink2.Link))
				Expect(responseBody).To(ContainSubstring(sideAdLink.Link))
				Expect(responseBody).To(ContainSubstring(link1.Link))
				Expect(responseBody).To(ContainSubstring(link2.Link))
			})
		})

		Context("given a non numberic result ID", func() {
			It("renders not found error", func() {
				user := FabricateUser(faker.Email(), faker.Password())

				response := MakeWebRequest("GET", "/results/invalid_id", nil, nil, user)
				responseBody := GetResponseBody(response)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
				Expect(responseBody).To(ContainSubstring("We cannot find the result you are looking for"))
			})
		})

		Context("given a result ID does NOT exist", func() {
			It("renders not found error", func() {
				user := FabricateUser(faker.Email(), faker.Password())

				response := MakeWebRequest("GET", "/results/999", nil, nil, user)
				responseBody := GetResponseBody(response)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
				Expect(responseBody).To(ContainSubstring("We cannot find the result you are looking for"))
			})
		})

		Context("given a result ID that does NOT belong to the user", func() {
			It("renders not found error", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				otherUser := FabricateUser(faker.Email(), faker.Password())
				result := FabricateResult(otherUser)

				response := MakeWebRequest("GET", fmt.Sprintf(`/results/%d`, result.ID), nil, nil, user)
				responseBody := GetResponseBody(response)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
				Expect(responseBody).To(ContainSubstring("We cannot find the result you are looking for"))
			})
		})
	})

	Describe("GET /results/:id/cache", func() {
		Context("given a valid result ID", func() {
			It("renders the result page cache", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				result := FabricateResult(user)

				response := MakeWebRequest("GET", fmt.Sprintf(`/results/%d/cache`, result.ID), nil, nil, user)
				responseBody := GetResponseBody(response)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
				Expect(responseBody).To(ContainSubstring(result.PageCache))
			})
		})

		Context("given a non numberic result ID", func() {
			It("renders not found error", func() {
				user := FabricateUser(faker.Email(), faker.Password())

				response := MakeWebRequest("GET", "/results/invalid_id/cache", nil, nil, user)
				responseBody := GetResponseBody(response)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
				Expect(responseBody).To(ContainSubstring("We cannot find the result you are looking for"))
			})
		})

		Context("given a result ID does NOT exist", func() {
			It("renders not found error", func() {
				user := FabricateUser(faker.Email(), faker.Password())

				response := MakeWebRequest("GET", "/results/999/cache", nil, nil, user)
				responseBody := GetResponseBody(response)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
				Expect(responseBody).To(ContainSubstring("We cannot find the result you are looking for"))
			})
		})

		Context("given a result ID that does NOT belong to the user", func() {
			It("renders not found error", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				otherUser := FabricateUser(faker.Email(), faker.Password())
				result := FabricateResult(otherUser)

				response := MakeWebRequest("GET", fmt.Sprintf(`/results/%d/cache`, result.ID), nil, nil, user)
				responseBody := GetResponseBody(response)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
				Expect(responseBody).To(ContainSubstring("We cannot find the result you are looking for"))
			})
		})
	})

	Describe("POST /results", func() {
		Context("given a valid file", func() {
			It("redirects to result list with the success message", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				header, body := CreateRequestInfoFormFile("test/fixtures/files/valid.csv")

				response := MakeWebRequest("POST", "/results", header, body, user)
				currentPath := GetCurrentPath(response)

				Expect(response.StatusCode).To(Equal(http.StatusFound))
				Expect(currentPath).To(Equal(constants.WebRoutes["results"]["index"]))
			})

			It("sets a success flash message", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				header, body := CreateRequestInfoFormFile("test/fixtures/files/valid.csv")

				response := MakeWebRequest("POST", "/results", header, body, user)
				cookie := GetResponseCookie(response)

				Expect(cookie).To(HaveKey(sessions.FlashTypeSuccess))
				Expect(cookie[sessions.FlashTypeSuccess]).To(ConsistOf("Successfully uploaded"))
			})
		})

		Context("given an empty file", func() {
			It("sets an error flashes message", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				header, body := CreateRequestInfoFormFile("test/fixtures/files/empty.csv")

				response := MakeWebRequest("POST", "/results", header, body, user)
				cookie := GetResponseCookie(response)

				Expect(response.StatusCode).To(Equal(http.StatusFound))
				Expect(cookie).To(HaveKey(sessions.FlashTypeError))
				Expect(cookie[sessions.FlashTypeError]).To(ConsistOf("Keywords: cannot be blank."))
			})
		})

		Context("given a file with too many keywords", func() {
			It("sets an error flashes message", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				header, body := CreateRequestInfoFormFile("test/fixtures/files/invalid.csv")

				response := MakeWebRequest("POST", "/results", header, body, user)
				cookie := GetResponseCookie(response)

				Expect(response.StatusCode).To(Equal(http.StatusFound))
				Expect(cookie).To(HaveKey(sessions.FlashTypeError))
				Expect(cookie[sessions.FlashTypeError]).To(ConsistOf("Keywords: the length must be between 1 and 1000."))
			})
		})
	})

	AfterEach(func() {
		CleanupDatabase([]string{"users", "results", "ad_links", "links"})
	})
})
