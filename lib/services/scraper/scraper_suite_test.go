package scraper_test

import (
	"testing"

	"go-google-scraper-challenge/test"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestScraper(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Scraper Suite")
}

var _ = BeforeSuite(func() {
	test.SetupTestEnvironment()
})
