package helpers_test

import (
	"net/url"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"go-google-scraper-challenge/helpers"
)

var _ = Describe("URL Helpers", func() {
	Describe("#IsActive", func() {
		Context("given the URL and the path are the same", func() {
			It("returns true", func() {
				currentpath := &url.URL{Path: "/path"}

				Expect(helpers.IsActive(currentpath, "/path")).To(BeTrue())
			})
		})

		Context("given the URL and the path are NOT the same", func() {
			It("returns false", func() {
				currentpath := &url.URL{Path: "/path"}

				Expect(helpers.IsActive(currentpath, "/another_path")).To(BeFalse())
			})
		})
	})

	Describe("#UrlFor", func() {
		Context("given a valid controller and action name", func() {
			It("returns the correct path", func() {
				Expect(helpers.UrlFor("users", "new")).To(Equal("/signup"))
			})
		})

		Context("given an INVALID controller and action name", func() {
			It("returns the given file content", func() {
				Expect(helpers.UrlFor("unknown", "unknown")).To(BeNil())
			})
		})
	})
})
