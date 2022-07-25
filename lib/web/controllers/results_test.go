package webcontrollers_test

import (
	"net/http"

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
})
