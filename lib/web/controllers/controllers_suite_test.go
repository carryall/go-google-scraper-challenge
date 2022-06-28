package webcontrollers_test

import (
	"testing"

	"go-google-scraper-challenge/test"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestWebControllers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Web Controllers Suite")
}

var _ = BeforeSuite(func() {
	test.SetupTestEnvironment()
})
