package tests

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"

	"go-google-scraper-challenge/controllers"
	"go-google-scraper-challenge/models"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/onsi/ginkgo"
)

// GenerateRequestBody return a request body
func GenerateRequestBody(data map[string]string) (body io.Reader) {
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

func MakeAuthenticatedRequest(method string, url string, header http.Header, body io.Reader, user *models.User) *http.Response {
	request := HTTPRequest(method, url, body)
	if header != nil {
		request.Header = header
	}
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	responseRecoder := httptest.NewRecorder()
	store, err := web.GlobalSessions.SessionStart(responseRecoder, request)
	if err != nil {
		ginkgo.Fail("Failed to start session" + err.Error())
	}

	err = store.Set(context.Background(), controllers.CurrentUserKey, user.Id)
	if err != nil {
		ginkgo.Fail("Failed to set current user" + err.Error())
	}

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

// GetResponseBody get response body from response, will fail the test if there is any error
func GetResponseBody(response *http.Response) string {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		ginkgo.Fail("Failed to read response body")
	}

	return string(body)
}

// GetJSONResponseBody get response body from response, will fail the test if there is any error
func GetJSONResponseBody(response *http.Response, v interface{}) {
	body := GetResponseBody(response)

	err := json.Unmarshal([]byte(body), v)
	if err != nil {
		ginkgo.Fail("Failed to unmarshal json response " + err.Error())
	}
}

// GetCurrentPath get current path from HTTP response and return the current url path
func GetCurrentPath(response *http.Response) string {
	url, err := response.Location()
	if err != nil {
		ginkgo.Fail("Failed to get current path: " + err.Error())
	}
	return url.Path
}

// GetSession get session with given key from cookie, will fail the test if cannot get session store
func GetSession(cookies []*http.Cookie, key string) interface{} {
	backgroundContext := context.Background()
	for _, c := range cookies {
		if c.Name == web.BConfig.WebConfig.Session.SessionName {
			store, err := web.GlobalSessions.GetSessionStore(c.Value)
			if err != nil {
				ginkgo.Fail("Failed to get store " + err.Error())
			}

			return store.Get(backgroundContext, key)
		}
	}
	return nil
}

// GetFlashMessage get Beego flash message out of array of http cookie
func GetFlashMessage(cookies []*http.Cookie) *web.FlashData {
	flash := web.NewFlash()

	for _, c := range cookies {
		if c.Name == "BEEGO_FLASH" {
			decodedCookie := decodeQueryString(c.Value)
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
		logs.Info(err)

		return ""
	}

	return decodedString
}
