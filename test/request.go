package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

	return makeAPIRequest(request, user)
}

func MakeFormRequest(method string, url string, formData url.Values, user *models.User) (*gin.Context, *httptest.ResponseRecorder) {
	request := buildFormRequest(method, url, nil, formData)

	return makeAPIRequest(request, user)
}

func makeAPIRequest(request *http.Request, user *models.User) (*gin.Context, *httptest.ResponseRecorder) {
	ctx, responseRecorder := CreateGinTestContext()
	ctx.Request = request
	ctx.Set("CurrentUser", user)

	return ctx, responseRecorder
}

func MakeWebRequest(method string, url string, body io.Reader, user *models.User) *http.Response {
	request := buildRequest(method, url, nil, body)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	if user != nil {
		cookie := FabricateAuthUserCookie(user.ID)
		request.Header.Set("Cookie", cookie.Name+"="+cookie.Value)
	}

	responseRecorder := httptest.NewRecorder()

	Engine.ServeHTTP(responseRecorder, request)

	return responseRecorder.Result()
}

func MakeWebFormRequest(method string, url string, formData url.Values, user *models.User) *http.Response {
	request := buildFormRequest(method, url, nil, formData)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	if user != nil {
		cookie := FabricateAuthUserCookie(user.ID)
		request.Header.Set("Cookie", cookie.Name+"="+cookie.Value)
	}

	responseRecorder := httptest.NewRecorder()

	Engine.ServeHTTP(responseRecorder, request)

	return responseRecorder.Result()
}

// HTTPRequest initiate new HTTP request and handle the error, will fail the test if there is any error
func HTTPRequest(method string, url string, body io.Reader) *http.Request {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		ginkgo.Fail("Request failed: " + err.Error())
	}

	return request
}

func GenerateRequestBody(params map[string]interface{}) *bytes.Buffer {
	queryParams, err := json.Marshal(params)
	if err != nil {
		ginkgo.Fail("Cannot parse the request body")
	}

	return bytes.NewBuffer(queryParams)
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
