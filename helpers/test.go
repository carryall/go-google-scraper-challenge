package helpers

import (
	"log"
	"net/http"
	"net/url"
	"strings"
)

// GetFlashMessage get Beego flash message out of array of http cookie
func GetFlashMessage(cookies []*http.Cookie) map[string]string {
	mapCookie := map[string]string{}

	for _, cookie := range cookies {
		if cookie.Name == "BEEGO_FLASH" {
			decodedCookie := DecodeQueryString(cookie.Value)
			cookiePart := strings.Split(strings.TrimSpace(decodedCookie), "#BEEGOFLASH#")
			if len(cookiePart) >= 2 {
				mapCookie[cookiePart[0]] = cookiePart[1]
			}
		}
	}

	return mapCookie
}

// DecodeQueryString decode query string to normal string,
// for example from %00error%23BEEGOFLASH%23PasswordConfirmation+Minimum+size+is+6%00 to
// error#BEEGOFLASH#PasswordConfirmation Minimum size is 6
func DecodeQueryString(encodedString string) string {
	decodedString, err := url.QueryUnescape(encodedString)
	if err != nil {
		log.Println(err)

		return ""
	}

	return decodedString
}
