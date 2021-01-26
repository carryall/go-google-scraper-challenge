package test_helpers

import (
	"go-google-scraper-challenge/models"

	"github.com/onsi/ginkgo"
)

// FabricateUser create a user with given User and return the user ID, will fail the test when there is any error
func FabricateUser(user *models.User) int64 {
	userID, err := models.CreateUser(user)
	if err != nil {
		ginkgo.Fail("Failed to add user " + err.Error())
	}

	return userID
}
