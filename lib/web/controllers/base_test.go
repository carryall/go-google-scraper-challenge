package webcontrollers_test

import (
	"net/url"

	"go-google-scraper-challenge/lib/models"
	webcontrollers "go-google-scraper-challenge/lib/web/controllers"
	. "go-google-scraper-challenge/test"

	"github.com/bxcodec/faker/v3"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type DummyController struct {
	webcontrollers.BaseController
}

func (c *DummyController) DummyAction(ctx *gin.Context) gin.H {
	return c.Data(ctx, gin.H{})
}

var _ = Describe("BaseController", func() {
	Describe("#Data", func() {
		It("returns the request URL as current path", func() {
			expectedURL := &url.URL{Path: "/url"}
			c, _ := CreateGinTestContext()
			c.Request = HTTPRequest("GET", "/url", nil)
			c.Request.URL = expectedURL
			c.Set("CurrentUser", nil)
			baseController := webcontrollers.BaseController{}

			data := baseController.Data(c, gin.H{})

			Expect(data["CurrentPath"]).To(Equal(expectedURL))
		})

		It("returns the correct controller and action name", func() {
			c, _ := CreateGinTestContext()
			c.Request = HTTPRequest("GET", "/url", nil)
			c.Request.URL = &url.URL{Path: "/url"}
			c.Set("CurrentUser", nil)
			dummyController := DummyController{}

			data := dummyController.DummyAction(c)

			Expect(data["ControllerName"]).To(Equal("dummy"))
			Expect(data["ActionName"]).To(Equal("dummy_action"))
		})

		Context("given a user in session", func() {
			It("returns the given user as current user", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				c, _ := CreateGinTestContext()
				c.Request = HTTPRequest("GET", "/url", nil)
				c.Request.URL = &url.URL{Path: "/url"}
				c.Set("CurrentUser", user)
				dummyController := DummyController{}

				data := dummyController.DummyAction(c)
				currentUser := data["CurrentUser"].(*models.User)

				Expect(currentUser.ID).To(Equal(user.ID))
			})
		})

		Context("given NO user in session", func() {
			It("returns NIL as current user", func() {
				c, _ := CreateGinTestContext()
				c.Request = HTTPRequest("GET", "/url", nil)
				c.Request.URL = &url.URL{Path: "/url"}
				c.Set("CurrentUser", nil)
				dummyController := DummyController{}

				data := dummyController.DummyAction(c)

				Expect(data["CurrentUser"]).To(BeNil())
			})
		})

		AfterEach(func() {
			CleanupDatabase([]string{"users"})
		})
	})
})
