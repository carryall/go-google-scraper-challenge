package test

import (
	"fmt"
	"net/http"

	"github.com/dnaeon/go-vcr/recorder"
	. "github.com/onsi/ginkgo"
)

func RecordResponse(name string, url string) {
	rec, err := recorder.New(fmt.Sprintf("test/fixtures/vcr/%s", name))
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
