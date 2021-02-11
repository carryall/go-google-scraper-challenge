package oauth

import (
	"context"
	"fmt"
	"log"
	"time"

	"go-google-scraper-challenge/helpers"
	"go-google-scraper-challenge/models"

	"github.com/beego/beego/v2/server/web"
	app_context "github.com/beego/beego/v2/server/web/context"
	"github.com/jackc/pgx/v4"
	pg "github.com/vgarvardt/go-oauth2-pg"
	"github.com/vgarvardt/go-pg-adapter/pgx4adapter"
	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/server"
)

var oauthServer *server.Server
var clientStore *pg.ClientStore

// SetUpOauth setup OAuth server
func SetUpOauth() {
	dbURL, err := web.AppConfig.String("db_url")
	if err != nil {
		log.Fatal("Database URL not found: ", err)
	}

	pgxConn, err := pgx.Connect(context.TODO(), dbURL)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	manager := manage.NewDefaultManager()

	// use PostgreSQL token store with pgx.Connection adapter
	adapter := pgx4adapter.NewConn(pgxConn)
	tokenStore, err := pg.NewTokenStore(adapter, pg.WithTokenStoreGCInterval(time.Minute))
	if err != nil {
		log.Fatal("Failed to create the token store: ", err)
	}
	defer tokenStore.Close()

	store, err := pg.NewClientStore(adapter)
	if err != nil {
		log.Fatal("Failed to create the client store: ", err)
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

// GenerateToken handle token request, will return error if request from given context is invalid
func GenerateToken(c *app_context.Context) (err error) {
	return oauthServer.HandleTokenRequest(c.ResponseWriter, c.Request)
}

// GetClientStore returns OAuth client store
func GetClientStore() *pg.ClientStore {
	return clientStore
}

func internalErrorHandler(err error) (response *errors.Response) {
	log.Println("Internal Error:", err.Error())

	response = errors.NewResponse(errors.ErrInvalidClient, errors.StatusCodes[errors.ErrInvalidClient])
	response.Description = errors.Descriptions[errors.ErrInvalidClient]
	return response
}

func responseErrorHandler(re *errors.Response) {
	log.Println("Oauth server response Error:", re.Error.Error())
}

func passwordAuthorizationHandler(email string, password string) (userID string, err error) {
	user, err := models.GetUserByEmail(email)
	if err != nil {
		return "", errors.ErrInvalidClient
	}

	if helpers.CompareHashWithPassword(user.HashedPassword, password) {
		return fmt.Sprint(user.Id), nil
	}
	return "", errors.ErrInvalidClient
}
