package webmiddlewares_test

import (
	"testing"

	"go-google-scraper-challenge/test"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestWebMiddlewares(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Web Middlewares Suite")
}

var _ = BeforeSuite(func() {
	test.SetupTestEnvironment()
})
