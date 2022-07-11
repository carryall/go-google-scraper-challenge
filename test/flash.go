package test

import (
	"net/http"

	"go-google-scraper-challenge/helpers/log"

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
