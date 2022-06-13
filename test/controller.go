package test

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"

	"go-google-scraper-challenge/lib/models"

	"github.com/gin-gonic/gin"
	"github.com/onsi/ginkgo"
)

func MakeJSONRequest(method string, url string, header http.Header, body io.Reader, user *models.User) (*gin.Context, *httptest.ResponseRecorder) {
	request := buildRequest(method, url, header, body)
	request.Header.Add("Content-Type", "application/json")

	if user != nil {
		accessToken := FabricateAuthToken(user.ID)
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	}

	return MakeRequest(request)
}

func MakeFormRequest(method string, url string, formData url.Values, user *models.User) (*gin.Context, *httptest.ResponseRecorder) {
	request := buildFormRequest(method, url, nil, formData)

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

func buildFormRequest(method string, url string, header http.Header, formData url.Values) (request *http.Request) {
	request = HTTPRequest(method, url, nil)

	if header != nil {
		request.Header = header
	}
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	if method == "POST" {
		request.PostForm = formData
	} else {
		request.Form = formData
	}

	return request
}

func buildRequest(method string, url string, header http.Header, body io.Reader) (request *http.Request) {
	request = HTTPRequest(method, url, body)

	if header != nil {
		request.Header = header
	}
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return request
}
