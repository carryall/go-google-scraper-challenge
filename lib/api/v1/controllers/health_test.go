package controllers_test

import (
	"net/http"

	"go-google-scraper-challenge/lib/api/v1/controllers"
	"go-google-scraper-challenge/test"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("HealthController", func() {
	Describe("GET /health", func() {
		It("returns status OK", func() {
			c, resp := test.CreateGinTestContext()
			healthController := controllers.HealthController{}

			healthController.HealthStatus(c)

			Expect(resp.Code).To(Equal(http.StatusOK))
		})

		It("returns correct response body", func() {
			c, resp := test.CreateGinTestContext()
			healthController := controllers.HealthController{}

			healthController.HealthStatus(c)

			Expect(resp.Body.String()).To(Equal("{\"status\":\"alive\"}"))
		})
	})
})
