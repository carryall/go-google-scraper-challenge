package test

import (
	"go-google-scraper-challenge/lib/services/oauth"

	"github.com/onsi/ginkgo"
)

func FabricateAuthClient() oauth.OAuthClient {
	authClient, err := oauth.GenerateClient()
	if err != nil {
		ginkgo.Fail("Fail to fablicate auth client")
	}

	return authClient
}
