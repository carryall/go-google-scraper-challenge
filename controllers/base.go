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
}

func (base *BaseController) Prepare() {
	helpers.SetControllerAttributes(&base.Controller)

	base.Layout = "layouts/default.html"
}

func (base *BaseController) SetCurrentUser(user *models.User) {
	err := base.SetSession(CurrentUserKey, user.Id)
	if err != nil {
		log.Fatal("Fail to set current user", err.Error())
	}
}

func (base *BaseController) GetCurrentUser() (user *models.User) {
	userId := base.GetSession(CurrentUserKey)
	if userId == nil {
		return nil
	}

	user, err := models.GetUserById(userId.(int64))
	if err != nil {
		return nil
	}
	return user
}

func (base *BaseController) ClearCurrentUser() error {
	return base.DelSession(CurrentUserKey)
}

func (base *BaseController) EnsureAuthenticatedUser(redirect bool) {
	currentUser := base.GetCurrentUser()
	if currentUser == nil {
		if redirect {
			base.Controller.Redirect("/signin", http.StatusFound)
		} else {
			base.Controller.Abort(fmt.Sprint(http.StatusMethodNotAllowed))
		}
	}
	base.Controller.Data["CurrentUser"] = currentUser
}

func (base *BaseController) EnsureGuestUser(redirect bool) {
	currentUser := base.GetCurrentUser()
	if currentUser != nil {
		if redirect {
			base.Controller.Redirect("/", http.StatusFound)
		} else {
			base.Controller.Abort(fmt.Sprint(http.StatusMethodNotAllowed))
		}
	}
}
