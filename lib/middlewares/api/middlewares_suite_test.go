package apimiddlewares_test

import (
	"testing"

	"go-google-scraper-challenge/test"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAPIMiddlewares(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "API Middlewares Suite")
}

var _ = BeforeSuite(func() {
	test.SetupTestEnvironment()
})
