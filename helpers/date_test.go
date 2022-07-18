package helpers_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"go-google-scraper-challenge/helpers"
)

var _ = Describe("Date", func() {
	Describe("#FormatDateTime", func() {
		Context("given a date time", func() {
			It("returns a formatted date time in string", func() {
				dateTime := time.Date(2022, time.Month(5), 20, 7, 30, 35, 0, time.UTC)
				htmlDate := helpers.FormatDateTime(dateTime)

				Expect(htmlDate).To(Equal("20/05/2022 14:30"))
			})
		})
	})
})
