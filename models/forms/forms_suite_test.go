package forms_test

import (
	"testing"

	"go-google-scraper-challenge/initializers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestForms(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Forms Suite")
}

var _ = BeforeSuite(func() {
	initializers.SetupTestEnvironment()
})

var _ = AfterSuite(func() {
	initializers.CleanupDatabase()
})
