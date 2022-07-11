package test

import (
	"net/http"

	"go-google-scraper-challenge/helpers/log"
	"go-google-scraper-challenge/lib/sessions"

	"github.com/gorilla/securecookie"
)

func FabricateCookieWithFlashes(flashes map[string]interface{}) *http.Cookie {
	codecs := securecookie.CodecsFromPairs([]byte("secret"))
	data := make(map[interface{}]interface{}, len(flashes))
	for key, value := range flashes {
		var newKey interface{} = key
		data[newKey] = value
	}
	encoded, err := securecookie.EncodeMulti("google_scraper_session", data, codecs...)
	if err != nil {
		log.Error("Failed to encode multi: ", err)
	}

	cookie := http.Cookie{
		Name:  "google_scraper_session",
		Value: encoded,
	}

	return &cookie
}

func GetFlashMessage(cookies []*http.Cookie) map[string][]string {
	flashes := map[string][]string{}
	for _, c := range cookies {
		if c.Name == "google_scraper_session" {
			decodedCookie := DecodeCookieString(c.Value)

			if decodedCookie[sessions.FlashTypeSuccess] != nil {
				flashes[sessions.FlashTypeSuccess] = decodedCookie[sessions.FlashTypeSuccess].([]string)
			}

			if decodedCookie[sessions.FlashTypeInfo] != nil {
				flashes[sessions.FlashTypeInfo] = decodedCookie[sessions.FlashTypeInfo].([]string)
			}

			if decodedCookie[sessions.FlashTypeError] != nil {
				flashes[sessions.FlashTypeError] = decodedCookie[sessions.FlashTypeError].([]string)
			}
		}
	}

	return flashes
}
