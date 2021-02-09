package controllers

import (
	"log"
	"net/http"

	"go-google-scraper-challenge/helpers"
	"go-google-scraper-challenge/models"

	"github.com/beego/beego/v2/server/web"
)

const (
	CurrentUserKey = "CURRENT_USER_ID"
)

type BaseController struct {
	web.Controller
}

func (base *BaseController) Prepare() {
	helpers.SetControllerAttributes(&base.Controller)
}

func (base *BaseController) SetCurrentUser(user *models.User) {
	err := base.Controller.SetSession(CurrentUserKey, user.Id)
	if err != nil {
		log.Fatal("Fail to set current user", err.Error())
	}
}

func (base *BaseController) CurrentUser() (user *models.User) {
	userId := base.Controller.GetSession(CurrentUserKey)
	if userId == nil {
		return nil
	}

	user, err := models.GetUserById(userId.(int64))
	if err != nil {
		return nil
	}
	return user
}

func (base *BaseController) EnsureAuthenticatedUser() {
	currentUser := base.CurrentUser()
	if currentUser == nil {
		base.Controller.Redirect("/signin", http.StatusFound)
	}
}

func (base *BaseController) EnsureGuestUser() {
	currentUser := base.CurrentUser()
	if currentUser != nil {
		base.Controller.Redirect("/", http.StatusFound)
	}
}
