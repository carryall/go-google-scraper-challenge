package webcontrollers_test

import (
	"fmt"
	"net/http"

	"go-google-scraper-challenge/constants"
	"go-google-scraper-challenge/lib/models"
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
				response := MakeWebRequest("GET", "/", nil, user)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
			})
		})

		Context("given NO user is signed in", func() {
			It("redirects to signin path", func() {
				response := MakeWebRequest("GET", "/", nil, nil)

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

				response := MakeWebRequest("GET", fmt.Sprintf(`/results/%d`, result.ID), nil, user)
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

				response := MakeWebRequest("GET", "/results/invalid_id", nil, user)
				responseBody := GetResponseBody(response)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
				Expect(responseBody).To(ContainSubstring(constants.ResultNotFound))
			})
		})

		Context("given a result ID does NOT exist", func() {
			It("renders not found error", func() {
				user := FabricateUser(faker.Email(), faker.Password())

				response := MakeWebRequest("GET", "/results/999", nil, user)
				responseBody := GetResponseBody(response)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
				Expect(responseBody).To(ContainSubstring(constants.ResultNotFound))
			})
		})

		Context("given a result ID that does NOT belong to the user", func() {
			It("renders not found error", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				otherUser := FabricateUser(faker.Email(), faker.Password())
				result := FabricateResult(otherUser)

				response := MakeWebRequest("GET", fmt.Sprintf(`/results/%d`, result.ID), nil, user)
				responseBody := GetResponseBody(response)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
				Expect(responseBody).To(ContainSubstring(constants.ResultNotFound))
			})
		})
	})

	AfterEach(func() {
		CleanupDatabase([]string{"users", "results", "ad_links", "links"})
	})
})
