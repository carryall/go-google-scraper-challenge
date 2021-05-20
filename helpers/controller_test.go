package helpers_test

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"go-google-scraper-challenge/controllers"
	"go-google-scraper-challenge/helpers"
)

var _ = Describe("Controller", func() {
	Describe("#SetControllerAttributes", func() {
		Context("given a valid controller", func() {
			type TestController struct {
				controllers.BaseController
			}

			It("sets controller name and action name", func() {
				testController := TestController{}
				testController.Init(context.NewContext(), "TestController", "ActionName", web.BeeApp)
				helpers.SetControllerAttributes(&testController.Controller)

				Expect(testController.Data["ControllerName"]).To(Equal("test"))
				Expect(testController.Data["ActionName"]).To(Equal("action-name"))
			})
		})
	})
})
