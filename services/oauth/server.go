package oauth_services

import (
	"context"
	"fmt"
	"go-google-scraper-challenge/helpers"
	"go-google-scraper-challenge/models"
	"log"
	"net/http"
	"time"

	"github.com/beego/beego/v2/server/web"
	"github.com/jackc/pgx/v4"
	pg "github.com/vgarvardt/go-oauth2-pg"
	"github.com/vgarvardt/go-pg-adapter/pgx4adapter"
	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/server"
)

var oauthServer *server.Server
var clientStore *pg.ClientStore

func SetUpOauth() {
	dbURL, err := web.AppConfig.String("db_url")
	if err != nil {
		log.Fatal("Database URL not found: ", err)
	}

	pgxConn, err := pgx.Connect(context.TODO(), dbURL)
	if err != nil {
		log.Fatal("Database URL not found: ", err)
	}
	manager := manage.NewDefaultManager()

	// use PostgreSQL token store with pgx.Connection adapter
	adapter := pgx4adapter.NewConn(pgxConn)
	tokenStore, err := pg.NewTokenStore(adapter, pg.WithTokenStoreGCInterval(time.Minute))
	if err != nil {
		log.Fatal("Database URL not found: ", err)
	}
	defer tokenStore.Close()

	store, err := pg.NewClientStore(adapter)
	if err != nil {
		log.Fatal("Database URL not found: ", err)
	}

	manager.MapTokenStorage(tokenStore)
	manager.MapClientStorage(store)

	authServer := server.NewDefaultServer(manager)
	authServer.SetAllowGetAccessRequest(true)
	authServer.SetClientInfoHandler(server.ClientFormHandler)
	authServer.SetInternalErrorHandler(internalErrorHandler)
	authServer.SetResponseErrorHandler(responseErrorHandler)
	authServer.SetPasswordAuthorizationHandler(passwordAuthorizationHandler)
	manager.SetRefreshTokenCfg(manage.DefaultRefreshTokenCfg)

	oauthServer = authServer
	clientStore = store
}

func internalErrorHandler(err error) (response *errors.Response) {
	log.Println("Internal Error:", err.Error())
	return
}

func responseErrorHandler(re *errors.Response) {
	log.Println("Response Error:", re.Error.Error())
}

func passwordAuthorizationHandler(email string, password string) (userID string, err error) {
	user, err := models.FindUserByEmail(email)
	if err != nil {
		return "", err
	}

	if helpers.CompareHashedPasswords(user.EncryptedPassword, password) {
		return fmt.Sprint(user.Id), nil
	}
	return "", nil
}

func AuthrizeRequest(w http.ResponseWriter, r *http.Request) (err error) {
	return oauthServer.HandleAuthorizeRequest(w, r)
}

func GenerateToken(w http.ResponseWriter, r *http.Request) (err error) {
	return oauthServer.HandleTokenRequest(w, r)
}
