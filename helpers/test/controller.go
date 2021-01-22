package test_helpers

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"

	"github.com/beego/beego/v2/server/web"
	"github.com/onsi/ginkgo"
)

// HTTPRequest initiate new HTTP request and handle the error, will fail the test if there is any error
func HTTPRequest(method, url string, body io.Reader) *http.Request {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		ginkgo.Fail("Request failed: " + err.Error())
	}

	return request
}

// GetCurrentPath get current path from HTTP response and return the current url path
func GetCurrentPath(response *httptest.ResponseRecorder) *url.URL {
	path, err := response.Result().Location()
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
