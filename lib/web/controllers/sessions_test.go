package webcontrollers_test

import (
	. "go-google-scraper-challenge/test"
	"net/http"

	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SessionsController", func() {
	Describe("GET /signin", func() {
		Context("given user is not signed in", func() {
			It("renders with status 200", func() {
				response := MakeWebRequest("GET", "/signin", nil, nil)
				// ctx, response := MakeFormRequest("GET", "/signin", nil, nil)
				// sessions.Default(ctx)
				// sessionsController := webcontrollers.SessionsController{}
				// sessionsController.New(ctx)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
			})
		})

		// TODO: Test this when work on the result list screen
		XContext("given user already signed in", func() {
			It("redirects to root path", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				response := MakeWebRequest("GET", "/signin", nil, user)

				currentPath := GetCurrentPath(response)

				Expect(response.StatusCode).To(Equal(http.StatusFound))
				Expect(currentPath).To(Equal("/"))
			})
		})
	})
})
