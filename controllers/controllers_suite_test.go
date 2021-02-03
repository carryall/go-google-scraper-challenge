package controllers_test

import (
	"testing"

	"go-google-scraper-challenge/initializers"
	_ "go-google-scraper-challenge/routers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestControllers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controllers Suite")
}

var _ = BeforeSuite(func() {
	initializers.SetupTestEnvironment()
})
