package sessions

import (
	"go-google-scraper-challenge/helpers/log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const CurrentUserKey = "CURRENT_USER_ID"
const (
	FlashTypeSuccess = "success"
	FlashTypeInfo    = "info"
	FlashTypeError   = "error"
)

type Flash struct {
	Type    string
	Message string
}

func GetCurrentUserID(ctx *gin.Context) *int64 {
	session := sessions.Default(ctx)
	currentUserID := session.Get(CurrentUserKey)
	if currentUserID == nil {
		return nil
	}

	userID := currentUserID.(int64)

	return &userID
}

func SetCurrentUser(ctx *gin.Context, userID int64) {
	session := sessions.Default(ctx)
	session.Set(CurrentUserKey, userID)
	err := session.Save()
	if err != nil {
		log.Error("Fail to set current user", err.Error())
	}
}

func SetFlash(ctx *gin.Context, flashType string, flashMessage string) {
	session := sessions.Default(ctx)
	session.AddFlash(flashMessage, flashType)
	err := session.Save()
	if err != nil {
		log.Error(err)
	}
}

func GetFlash(ctx *gin.Context) map[string]interface{} {
	session := sessions.Default(ctx)

	flashes := map[string]interface{}{}
	flashes[FlashTypeInfo] = session.Flashes(FlashTypeInfo)
	flashes[FlashTypeError] = session.Flashes(FlashTypeError)
	flashes[FlashTypeSuccess] = session.Flashes(FlashTypeSuccess)

	err := session.Save()
	if err != nil {
		log.Error(err)
	}

	return flashes
}

func Clear(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	err := session.Save()
	if err != nil {
		log.Error("Fail to clear session", err.Error())
	}
}
