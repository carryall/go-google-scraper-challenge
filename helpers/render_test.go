package helpers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"go-google-scraper-challenge/helpers"
)

var _ = Describe("Render Helpers", func() {
	Describe("#RenderFile", func() {
		Context("given a file name", func() {
			It("returns the given file content", func() {
				htmlStr := helpers.RenderFile("test/fixtures/files/dummy.txt")

				Expect(string(htmlStr)).To(ContainSubstring(`file content`))
			})
		})
	})

	Describe("#RenderIcon", func() {
		Context("given an  icon file name", func() {
			It("returns the svg html tag", func() {
				htmlStr := helpers.RenderIcon("search", "extra_class")

				Expect(string(htmlStr)).To(HavePrefix(`<svg class="icon extra_class" viewBox="0 0 20 20">`))
				Expect(string(htmlStr)).To(HaveSuffix(`</svg>`))
			})
		})
	})
})
