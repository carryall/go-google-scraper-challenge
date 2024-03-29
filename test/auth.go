package test

import (
	"fmt"
	"net/http"
	"time"

	"go-google-scraper-challenge/helpers/log"
	"go-google-scraper-challenge/lib/services/oauth"
	"go-google-scraper-challenge/lib/sessions"

	"github.com/bxcodec/faker/v3"
	"github.com/gorilla/securecookie"
	"github.com/onsi/ginkgo"
	"gopkg.in/oauth2.v3/models"
)

func FabricateAuthClient() oauth.OAuthClient {
	authClient, err := oauth.GenerateClient()
	if err != nil {
		ginkgo.Fail("Fail to fablicate auth client")
	}

	return authClient
}

func FabricateAuthToken(userID int64) string {
	client := FabricateAuthClient()
	tokenInfo := &models.Token{
		ClientID:         client.ClientID,
		UserID:           fmt.Sprint(userID),
		Access:           faker.Password(),
		AccessCreateAt:   time.Now().Local(),
		AccessExpiresIn:  time.Hour * 2,
		Refresh:          faker.Password(),
		RefreshCreateAt:  time.Now().Local(),
		RefreshExpiresIn: time.Hour * 2,
	}

	err := oauth.GetTokenStore().Create(tokenInfo)
	if err != nil {
		ginkgo.Fail("Add TokenInfo failed: " + err.Error())
	}

	return tokenInfo.GetAccess()
}

func FabricateAuthUserCookie(userID int64) *http.Cookie {
	codecs := securecookie.CodecsFromPairs([]byte("secret"))
	data := make(map[interface{}]interface{})
	data[sessions.CurrentUserKey] = userID
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

func GetSessionUserID(cookies []*http.Cookie) interface{} {
	for _, c := range cookies {
		if c.Name == "google_scraper_session" {
			decodedCookie := DecodeCookieString(c.Value)

			if decodedCookie[sessions.CurrentUserKey] != nil {
				return fmt.Sprint(decodedCookie[sessions.CurrentUserKey])
			}
		}
	}

	return nil
}
