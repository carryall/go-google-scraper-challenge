package scraper_test

import (
	"go-google-scraper-challenge/lib/services/scraper"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UserAgent", func() {
	Describe("#RandomUserAgent", func() {
		It("returns a user agent with os name", func() {
			Expect(scraper.RandomUserAgent()).To(MatchRegexp(`(Macintosh|Windows|Linux)`))
		})

		It("returns a user agent with browser version", func() {
			Expect(scraper.RandomUserAgent()).To(Or(
				MatchRegexp(`(Firefox\/\d{2}.\d)`),
				MatchRegexp(`(Chrome\/\d{2}.\d.\d{4}.\d{1,3})`)))
		})
	})
})
