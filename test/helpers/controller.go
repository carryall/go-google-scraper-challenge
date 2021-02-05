package tests

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"

	"github.com/beego/beego/v2/server/web"
	"github.com/onsi/ginkgo"
)

// RequestBody return a request body
func RequestBody(data map[string]string) (body io.Reader) {
	rawData := url.Values{}
	for k, v := range data {
		rawData.Set(k, v)
	}
	body = strings.NewReader(rawData.Encode())

	return body
}

// MakeRequest make a HTTP request and return response
func MakeRequest(method string, url string, body io.Reader) *http.Response {
	request := HTTPRequest(method, url, body)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	responseRecoder := httptest.NewRecorder()
	web.BeeApp.Handlers.ServeHTTP(responseRecoder, request)

	return responseRecoder.Result()
}

// HTTPRequest initiate new HTTP request and handle the error, will fail the test if there is any error
func HTTPRequest(method string, url string, body io.Reader) *http.Request {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		ginkgo.Fail("Request failed: " + err.Error())
	}

	return request
}

// GetResponseBody get response body from response recoder, will fail the test if there us any error
func GetResponseBody(response *http.Response) string {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		ginkgo.Fail("Failed to read response body")
	}

	return string(body)
}

// GetJSONResponseBody get response body from response recoder, will fail the test if there us any error
func GetJSONResponseBody(response *http.Response, v interface{}) {
	body := GetResponseBody(response)

	err := json.Unmarshal([]byte(body), v)
	if err != nil {
		ginkgo.Fail("Failed to unmarshal json response " + err.Error())
	}
}

// GetCurrentPath get current path from HTTP response and return the current url path
func GetCurrentPath(response *http.Response) *url.URL {
	path, err := response.Location()
	if err != nil {
		ginkgo.Fail("Failed to get current path: " + err.Error())
	}
	return path
}

// GetFlashMessage get Beego flash message out of array of http cookie
func GetFlashMessage(cookies []*http.Cookie) *web.FlashData {
	flash := web.NewFlash()

	for _, cookie := range cookies {
		if cookie.Name == "BEEGO_FLASH" {
			decodedCookie := decodeQueryString(cookie.Value)
			// Trim null character out of the docoded cookie value
			trimedCookie := strings.Trim(decodedCookie, "\x00")
			cookieParts := strings.Split(trimedCookie, "#BEEGOFLASH#")
			if len(cookieParts) == 2 {
				flash.Data[cookieParts[0]] = cookieParts[1]
			}
		}
	}

	return flash
}

// decodeQueryString decode query string to normal string,
// for example from %00error%23BEEGOFLASH%23PasswordConfirmation+Minimum+size+is+6%00 to
// error#BEEGOFLASH#PasswordConfirmation Minimum size is 6
func decodeQueryString(encodedString string) string {
	decodedString, err := url.QueryUnescape(encodedString)
	if err != nil {
		log.Println(err)

		return ""
	}

	return decodedString
}
