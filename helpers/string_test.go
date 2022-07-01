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

		Context("given a sentence case string", func() {
			It("returns a kebab case string", func() {
				Expect(helpers.ToKebabCase("Sentence case string")).To(Equal("sentence-case-string"))
			})
		})
	})

	Describe("#ToSentenceCase", func() {
		Context("given a camel case string", func() {
			It("returns a sentence case string", func() {
				Expect(helpers.ToSentenceCase("CamelCaseString")).To(Equal("Camel case string"))
			})
		})

		Context("given a snake case string", func() {
			It("returns a sentence case string", func() {
				Expect(helpers.ToSentenceCase("snake_case_string")).To(Equal("Snake case string"))
			})
		})

		Context("given a kebab case string", func() {
			It("returns a sentence case string", func() {
				Expect(helpers.ToSentenceCase("kebab-case-string")).To(Equal("Kebab case string"))
			})
		})
	})

	Describe("#ToSnakeCase", func() {
		Context("given a camel case string", func() {
			It("returns a snake case string", func() {
				Expect(helpers.ToSnakeCase("CamelCaseString")).To(Equal("camel_case_string"))
			})
		})

		Context("given a kebab case string", func() {
			It("returns a snake case string", func() {
				Expect(helpers.ToSnakeCase("kebab-case-string")).To(Equal("kebab_case_string"))
			})
		})

		Context("given a sentence case string", func() {
			It("returns a snake case string", func() {
				Expect(helpers.ToSnakeCase("Sentence case string")).To(Equal("sentence_case_string"))
			})
		})
	})

	Describe("#ToTitleCase", func() {
		Context("given a camel case string", func() {
			It("returns a title case string", func() {
				Expect(helpers.ToTitleCase("CamelCaseString")).To(Equal("CAMELCASESTRING"))
			})
		})

		Context("given a kebab case string", func() {
			It("returns a title case string", func() {
				Expect(helpers.ToTitleCase("kebab-case-string")).To(Equal("KEBAB-CASE-STRING"))
			})
		})

		Context("given a sentence case string", func() {
			It("returns a title case string", func() {
				Expect(helpers.ToTitleCase("Sentence case string")).To(Equal("SENTENCE CASE STRING"))
			})
		})
	})
})
