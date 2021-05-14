package scraper_test

import (
	"fmt"
	"net/url"

	"go-google-scraper-challenge/initializers"
	"go-google-scraper-challenge/models"
	"go-google-scraper-challenge/services/scraper"
	. "go-google-scraper-challenge/tests"
	. "go-google-scraper-challenge/tests/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Scraper", func() {
	Describe("#Search", func() {
		Context("given valid keywords", func() {
			BeforeEach(func() {
				keyword := "keyword"
				visitURL := fmt.Sprintf("http://www.google.com/search?q=%s", url.QueryEscape(keyword))
				cassetteName := "scraper/success"

				RecordResponse(cassetteName, visitURL)
			})

			It("creates a result with the given keyword", func() {
				user := FabricateUser("dev@nimblehq.co", "password")
				keywords := []string{"keyword"}
				scraper.Search(keywords, user)

				userResults, err := models.GetResultsByUserId(user.Id)
				if err != nil {
					Fail("Failed to get user results")
				}

				Expect(userResults).To(HaveLen(1))
				Expect(userResults[0].Keyword).To(Equal("keyword"))
			})
		})
	})

	AfterEach(func() {
		initializers.CleanupDatabase("users")
		initializers.CleanupDatabase("results")
		initializers.CleanupDatabase("links")
		initializers.CleanupDatabase("ad_links")
	})
})
