package tests

import (
	"fmt"

	"go-google-scraper-challenge/models"
	"go-google-scraper-challenge/services/oauth"

	"github.com/bxcodec/faker/v3"
	"github.com/onsi/ginkgo"
	"golang.org/x/crypto/bcrypt"
)

// FabricateUser create a user with given email and password, will fail the test when there is any error
func FabricateUser(email string, password string) (user *models.User) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ginkgo.Fail("failed to generate hashed password " + err.Error())
	}

	user = &models.User{
		Email:          email,
		HashedPassword: string(hashedPassword),
	}

	userID, err := models.CreateUser(user)
	if err != nil {
		ginkgo.Fail("Failed to add user " + err.Error())
	}

	user, err = models.GetUserById(userID)
	if err != nil {
		ginkgo.Fail("Failed to get user " + err.Error())
	}

	return user
}

// FabricateOAuthClient create a OAuth Client, will fail the test when there is any error
func FabricateOAuthClient() (client oauth.OAuthClient) {
	client, err := oauth.GenerateClient()
	if err != nil {
		ginkgo.Fail("Failed to fabricate OAuth Client")
	}

	return client
}

func FabricateResult(user *models.User) (result *models.Result) {
	return FabricateResultWithParams(user, fmt.Sprintf("Keyword %s", faker.Word()), models.ResultStatusPending)
}

func FabricateResultWithParams(user *models.User, keyword string, status string) (result *models.Result) {
	result = &models.Result{
		User: user,
		Keyword: keyword,
		Status: status,
	}

	resultID, err := models.CreateResult(result)
	if err != nil {
		ginkgo.Fail("Failed to add result " + err.Error())
	}

	result, err = models.GetResultById(resultID)
	if err != nil {
		ginkgo.Fail("Failed to get result " + err.Error())
	}

	return result
}

func FabricateLink(result *models.Result) (link *models.Link)  {
	link = &models.Link{
		Result: result,
		Link: faker.URL(),
	}

	linkID, err := models.CreateLink(link)
	if err != nil {
		ginkgo.Fail("Failed to add adLink " + err.Error())
	}

	link, err = models.GetLinkById(linkID)
	if err != nil {
		ginkgo.Fail("Failed to get link " + err.Error())
	}

	return link
}

func FabricateAdLink(result *models.Result) (adLink *models.AdLink)  {
	adLink = &models.AdLink{
		Result: result,
		Link: faker.URL(),
		Position: models.AdLinkPositionTop,
		Type: models.AdLinkTypeLink,
	}

	adLinkID, err := models.CreateAdLink(adLink)
	if err != nil {
		ginkgo.Fail("Failed to add adLink " + err.Error())
	}

	adLink, err = models.GetAdLinkById(adLinkID)
	if err != nil {
		ginkgo.Fail("Failed to get adLink " + err.Error())
	}

	return adLink
}
