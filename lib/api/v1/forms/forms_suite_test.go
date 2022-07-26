package apiforms_test

import (
	"testing"

	"go-google-scraper-challenge/test"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAPIForms(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "API Forms Suite")
}

var _ = BeforeSuite(func() {
	test.SetupTestEnvironment()
})
