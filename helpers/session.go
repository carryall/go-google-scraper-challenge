package helpers

import (
	"go-google-scraper-challenge/helpers/log"
	"go-google-scraper-challenge/lib/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const CURRENT_USER_KEY = "CURRENT_USER_ID"

func GetCurrentUser(ctx *gin.Context) *models.User {
	session := sessions.Default(ctx)
	currentUserID := session.Get(CURRENT_USER_KEY)
	if currentUserID == nil {
		return nil
	}

	user, err := models.GetUserByID(currentUserID.(int64))
	if err != nil {
		log.Error("Fail to get current user", err.Error())

		return nil
	}

	return user
}

func SetCurrentUser(ctx *gin.Context, userID int64) {
	session := sessions.Default(ctx)
	session.Set(CURRENT_USER_KEY, userID)
	err := session.Save()
	if err != nil {
		log.Error("Fail to set current user", err.Error())
	}
}

func Clear(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	err := session.Save()
	if err != nil {
		log.Error("Fail to clear session", err.Error())
	}
}
