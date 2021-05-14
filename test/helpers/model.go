package tests

import (
	"go-google-scraper-challenge/models"
	"go-google-scraper-challenge/services/oauth"

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
