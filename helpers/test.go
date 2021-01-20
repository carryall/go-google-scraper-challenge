package helpers

import (
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

// GetFlashMessage get Beego flash message out of array of http cookie
func GetFlashMessage(cookies []*http.Cookie) *web.FlashData {
	flash := web.NewFlash()

	for _, cookie := range cookies {
		if cookie.Name == "BEEGO_FLASH" {
			decodedCookie := DecodeQueryString(cookie.Value)
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
