package helpers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"go-google-scraper-challenge/helpers"
)

var _ = Describe("Assets", func() {
	Describe("#AssetsCSS", func() {
		Context("given a css file name", func() {
			It("returns a html string", func() {
				htmlStr := helpers.AssetsCSS("file_name")

				Expect(string(htmlStr)).To(Equal(`<link href="../../static/stylesheets/file_name" rel="stylesheet" type="text/css" />`))
			})
		})
	})
})
