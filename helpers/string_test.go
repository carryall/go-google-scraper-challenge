package helpers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"go-google-scraper-challenge/helpers"
)

var _ = Describe("String", func() {
	Describe("#ToKebabCase", func() {
		Context("given a camel case string", func() {
			It("returns a kebab case string", func() {
				Expect(helpers.ToKebabCase("CamelCaseString")).To(Equal("camel-case-string"))
			})
		})

		Context("given a snake case string", func() {
			It("returns a kebab case string", func() {
				Expect(helpers.ToKebabCase("snake_case_string")).To(Equal("snake-case-string"))
			})
		})
	})
})
