package webforms_test

import (
	"testing"

	"go-google-scraper-challenge/test"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestForms(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Web Forms Suite")
}

var _ = BeforeSuite(func() {
	test.SetupTestEnvironment()
})
