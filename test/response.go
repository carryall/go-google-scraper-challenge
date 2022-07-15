package test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/onsi/ginkgo"
)

func GetResponseBody(response *http.Response) string {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		ginkgo.Fail("Failed to read response body")
	}

	return string(body)
}

func GetJSONResponseBody(response *http.Response, v interface{}) {
	body := responseBody(response)

	err := json.Unmarshal([]byte(body), v)

	if err != nil {
		ginkgo.Fail("Failed to unmarshal json response " + err.Error())
	}
}

func responseBody(response *http.Response) string {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		ginkgo.Fail("Failed to read response body")
	}

	return string(body)
}
