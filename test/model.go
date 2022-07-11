package test

import (
	"fmt"

	"go-google-scraper-challenge/helpers/log"
	"go-google-scraper-challenge/lib/models"
	"go-google-scraper-challenge/lib/services/oauth"

	"github.com/bxcodec/faker/v3"
)

// FabricateUser create a user with given email and password, will fail the test when there is any error
func FabricateUser(email string, password string) (user *models.User) {
	user = &models.User{
		Email:    email,
		Password: password,
	}

	userID, err := models.CreateUser(user)
	if err != nil {
		log.Fatal("Failed to add user " + err.Error())
	}

	user, err = models.GetUserByID(userID)
	if err != nil {
		log.Fatal("Failed to get user " + err.Error())
	}

	return user
}

// FabricateOAuthClient create a OAuth Client, will fail the test when there is any error
func FabricateOAuthClient() (client oauth.OAuthClient) {
	client, err := oauth.GenerateClient()
	if err != nil {
		log.Fatal("Failed to fabricate OAuth Client")
	}

	return client
}

func FabricateResult(user *models.User) (result *models.Result) {
	return FabricateResultWithParams(user, fmt.Sprintf("Keyword %s", faker.Word()), models.ResultStatusPending)
}

func FabricateResultWithParams(user *models.User, keyword string, status string) (result *models.Result) {
	result = &models.Result{
		UserID:  int64(user.ID),
		Keyword: keyword,
		Status:  status,
	}

	resultID, err := models.CreateResult(result)
	if err != nil {
		log.Fatal("Failed to add result " + err.Error())
	}

	result, err = models.GetResultByID(resultID, nil, []string{})
	if err != nil {
		log.Fatal("Failed to get result " + err.Error())
	}

	return result
}

func FabricateLink(result *models.Result) (link *models.Link) {
	link = &models.Link{
		ResultID: int64(result.ID),
		Link:     faker.URL(),
	}

	linkID, err := models.CreateLink(link)
	if err != nil {
		log.Fatal("Failed to add adLink " + err.Error())
	}

	link, err = models.GetLinkByID(linkID)
	if err != nil {
		log.Fatal("Failed to get link " + err.Error())
	}

	return link
}

func FabricateAdLink(result *models.Result) (adLink *models.AdLink) {
	return FabricateAdLinkWithParams(result, models.AdLinkPositionTop)
}

func FabricateAdLinkWithParams(result *models.Result, position string) (adLink *models.AdLink) {
	adLink = &models.AdLink{
		ResultID: int64(result.ID),
		Link:     faker.URL(),
		Position: position,
		Type:     models.AdLinkTypeLink,
	}

	adLinkID, err := models.CreateAdLink(adLink)
	if err != nil {
		log.Fatal("Failed to add adLink " + err.Error())
	}

	adLink, err = models.GetAdLinkByID(adLinkID)
	if err != nil {
		log.Fatal("Failed to get adLink " + err.Error())
	}

	return adLink
}
