package test

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/onsi/ginkgo"
)

func MakeFormRequest(method string, url string, formData url.Values) (*gin.Context, *httptest.ResponseRecorder) {
	request := HTTPRequest(method, url, nil)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	if method == "POST" {
		request.PostForm = formData
	} else {
		request.Form = formData
	}

	return MakeRequest(request)
}

func MakeJSONRequest(method string, url string, body io.Reader) (*gin.Context, *httptest.ResponseRecorder) {
	request := HTTPRequest(method, url, body)
	request.Header.Add("Content-Type", "application/json")

	return MakeRequest(request)
}

// MakeRequest make a HTTP request and return response
func MakeRequest(request *http.Request) (*gin.Context, *httptest.ResponseRecorder) {
	ctx, responseRecorder := CreateGinTestContext()
	ctx.Request = request

	return ctx, responseRecorder
}

// HTTPRequest initiate new HTTP request and handle the error, will fail the test if there is any error
func HTTPRequest(method string, url string, body io.Reader) *http.Request {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		ginkgo.Fail("Request failed: " + err.Error())
	}

	return request
}

// GetJSONResponseBody get response body from response, will fail the test if there is any error
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
