package test_helpers

import (
	"go-google-scraper-challenge/models"

	"github.com/onsi/ginkgo"
	"golang.org/x/crypto/bcrypt"
)

// FabricateUser create a user with given User and return the user ID, will fail the test when there is any error
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
