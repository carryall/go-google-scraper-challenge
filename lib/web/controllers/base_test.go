package webcontrollers_test

import (
	"net/url"

	webcontrollers "go-google-scraper-challenge/lib/web/controllers"
	. "go-google-scraper-challenge/test"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type DummyController struct {
	webcontrollers.BaseController
}

func (c DummyController) DummyAction(ctx *gin.Context) gin.H {
	return c.Data(ctx, gin.H{})
}

var _ = Describe("BaseController", func() {
	Describe("#Data", func() {
		It("returns the request URL as current path", func() {
			expectedURL := &url.URL{Path: "/url"}
			c, _ := CreateGinTestContext()
			c.Request.URL = expectedURL
			baseController := webcontrollers.BaseController{}

			data := baseController.Data(c, gin.H{})

			Expect(data["CurrentPath"]).To(Equal(expectedURL))
		})

		It("returns the correct controller and action name", func() {
			c, _ := CreateGinTestContext()
			dummyController := DummyController{}

			data := dummyController.DummyAction(c)

			Expect(data["ControllerName"]).To(Equal("dummy"))
			Expect(data["ActionName"]).To(Equal("dummy_action"))
		})
	})
})