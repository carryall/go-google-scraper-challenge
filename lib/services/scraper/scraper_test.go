package scraper_test

import (
	"fmt"
	"net/url"

	"go-google-scraper-challenge/lib/models"
	"go-google-scraper-challenge/lib/services/scraper"
	. "go-google-scraper-challenge/test"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Scraper", func() {
	Describe("#Run", func() {
		Context("given valid keyword", func() {
			BeforeEach(func() {
				keyword := "ergonomic chair"
				visitURL := fmt.Sprintf("http://www.google.com/search?q=%s", url.QueryEscape(keyword))
				cassetteName := "scraper/success"

				RecordResponse(cassetteName, visitURL)
			})

			It("creates a result with the given keyword", func() {
				user := FabricateUser("dev@nimblehq.co", "password")
				result := FabricateResultWithParams(user, "ergonomic chair", models.ResultStatusPending)

				service := scraper.Scraper{
					Result: result,
				}
				err := service.Run()
				if err != nil {
					Fail("Failed to run scraper service")
				}

				userResults, err := models.GetUserResults(user.ID, []string{"User", "AdLinks", "Links"}, "")
				if err != nil {
					Fail("Failed to get user results")
				}

				Expect(userResults).To(HaveLen(1))
				Expect(userResults[0].Keyword).To(Equal("ergonomic chair"))
				Expect(userResults[0].Status).To(Equal(models.ResultStatusCompleted))
				Expect(userResults[0].PageCache).NotTo(BeEmpty())

				resultID := userResults[0].ID
				links, err := models.GetLinksByResultID(resultID)
				if err != nil {
					Fail("Failed to get result links")
				}
				Expect(links).NotTo(HaveLen(0))

				adLinks, err := models.GetAdLinksByResultID(resultID)
				if err != nil {
					Fail("Failed to get result ad links")
				}
				Expect(adLinks).NotTo(HaveLen(0))
			})
		})
	})

	AfterEach(func() {
		CleanupDatabase([]string{"users", "results", "links", "ad_links"})
	})
})
