package controllers_test

import (
	"testing"

	"go-google-scraper-challenge/test"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestOAuthControllers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "OAuth Controllers Suite")
}

var _ = BeforeSuite(func() {
	test.SetupTestEnvironment()
})
