package controllers

import (
	"fmt"
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

	CurrentUser *models.User
	controllerName string
	actionName string
}

func (b *BaseController) Prepare() {
	controller := &b.Controller
	helpers.SetControllerAttributes(controller)

	b.controllerName, b.actionName = controller.GetControllerAndAction()
	b.Layout = "layouts/default.html"
}

func (b *BaseController) SetCurrentUser(user *models.User) {
	err := b.SetSession(CurrentUserKey, user.Id)
	if err != nil {
		log.Fatal("Fail to set current user", err.Error())
	}
}

func (b *BaseController) GetCurrentUser() (*models.User) {
	userId := b.GetSession(CurrentUserKey)
	if userId == nil {
		return nil
	}

	user, err := models.GetUserById(userId.(int64))
	if err != nil {
		return nil
	}
	return user
}

func (b *BaseController) ClearCurrentUser() error {
	return b.DelSession(CurrentUserKey)
}

func (b *BaseController) EnsureAuthenticatedUser() {
	currentUser := b.GetCurrentUser()
	if currentUser == nil {
		if helpers.IsActionWithGetMethod(b.actionName) {
			b.Controller.Redirect("/signin", http.StatusFound)
		} else {
			b.Controller.Abort(fmt.Sprint(http.StatusMethodNotAllowed))
		}
	}
	b.CurrentUser = currentUser
	b.Controller.Data["CurrentUser"] = currentUser
}

func (b *BaseController) EnsureGuestUser() {
	currentUser := b.GetCurrentUser()
	if currentUser != nil {
		if helpers.IsActionWithGetMethod(b.actionName) {
			b.Controller.Redirect("/", http.StatusFound)
		} else {
			b.Controller.Abort(fmt.Sprint(http.StatusMethodNotAllowed))
		}
	}
}
