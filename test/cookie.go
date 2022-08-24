package test

import (
	"fmt"
	"net/http"

	"go-google-scraper-challenge/helpers/log"
	"go-google-scraper-challenge/lib/sessions"

	"github.com/gorilla/securecookie"
)

func GetResponseCookie(response *http.Response) map[string]interface{} {
	encodedSession := ""
	for _, cookie := range response.Cookies() {
		if cookie.Name == "google_scraper_session" {
			encodedSession = cookie.Value
		}
	}

	return DecodeCookieString(encodedSession)
}

func DecodeCookieString(encodedString string) map[string]interface{} {
	codecs := securecookie.CodecsFromPairs([]byte("secret"))
	data := map[interface{}]interface{}{}
	err := securecookie.DecodeMulti("google_scraper_session", encodedString, &data, codecs...)
	if err != nil {
		log.Errorln(err.Error())

		return nil
	}

	decodedCookie := map[string]interface{}{}
	for key, value := range data {
		strKey := fmt.Sprint(key)
		if strKey != sessions.CurrentUserKey && value != nil {
			messages := []string{}
			for _, value := range value.([]interface{}) {
				messages = append(messages, fmt.Sprint(value))
			}
			decodedCookie[strKey] = messages
		} else {
			decodedCookie[strKey] = fmt.Sprint(value)
		}
	}

	return decodedCookie
}
