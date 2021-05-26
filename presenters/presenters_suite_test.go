package presenters_test

import (
	"testing"

	"go-google-scraper-challenge/initializers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestModels(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Presenters Suite")
}

var _ = BeforeSuite(func() {
	initializers.SetupTestEnvironment()
})
