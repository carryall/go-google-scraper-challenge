package helpers_test

import (
	"testing"

	"go-google-scraper-challenge/initializers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHelpers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Helpers Suite")
}

var _ = BeforeSuite(func() {
	initializers.SetupTestEnvironment()
})
