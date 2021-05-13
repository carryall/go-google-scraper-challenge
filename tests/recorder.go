package tests

import (
	"fmt"
	"net/http"

	"go-google-scraper-challenge/initializers"

	"github.com/dnaeon/go-vcr/recorder"
	. "github.com/onsi/ginkgo"
)

func RecordResponse(name string, url string)  {
	rec, err := recorder.New(fmt.Sprintf("%s/tests/fixtures/vcr/%s", initializers.AppRootDir(), name))
	if err != nil {
		Fail(err.Error())
	}
	defer func() {
		err := rec.Stop()
		if err != nil {
			Fail(err.Error())
		}
	}()

	// Create an HTTP client and inject our transport
	client := &http.Client{Transport: rec}
	_, err = client.Get(url)
	if err != nil {
		Fail(err.Error())
	}
}
