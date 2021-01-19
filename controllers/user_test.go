package controllers_test

import (
	"net/http"
	"net/http/httptest"

	"go-google-scraper-challenge/initializers"

	"github.com/beego/beego/v2/server/web"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UserController", func() {
	Describe("GET /signup", func() {
		It("renders with status 200", func() {
			request, _ := http.NewRequest("GET", "/signup", nil)
			response := httptest.NewRecorder()
			web.BeeApp.Handlers.ServeHTTP(response, request)

			Expect(response.Code).To(Equal(http.StatusOK))
		})
	})

	Describe("Post /users", func() {
		// TODO: add test cases
	})

	AfterEach(func() {
		initializers.CleanupDatabase("user")
	})
})
