package oauth

import (
	"context"
	"fmt"
	"time"

	"go-google-scraper-challenge/database"
	"go-google-scraper-challenge/helpers"
	"go-google-scraper-challenge/helpers/log"
	"go-google-scraper-challenge/lib/models"

	"github.com/gin-gonic/gin"
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
	dbURL := database.GetDatabaseURL()

	pgxConn, err := pgx.Connect(context.TODO(), dbURL)
	if err != nil {
		log.Error("Failed to connect to database: ", err)
	}
	manager := manage.NewDefaultManager()

	// use PostgreSQL token store with pgx.Connection adapter
	adapter := pgx4adapter.NewConn(pgxConn)
	tokenStore, err := pg.NewTokenStore(adapter, pg.WithTokenStoreGCInterval(time.Minute))
	if err != nil {
		log.Error("Failed to create the token store: ", err)
	}
	defer tokenStore.Close()

	store, err := pg.NewClientStore(adapter)
	if err != nil {
		log.Error("Failed to create the client store: ", err)
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
func GenerateToken(c *gin.Context) (err error) {
	return oauthServer.HandleTokenRequest(c.Writer, c.Request)
}

// GetClientStore returns OAuth client store
func GetClientStore() *pg.ClientStore {
	return clientStore
}

func internalErrorHandler(err error) (response *errors.Response) {
	log.Info("Internal Error:", err.Error())

	response = errors.NewResponse(errors.ErrInvalidClient, errors.StatusCodes[errors.ErrInvalidClient])
	response.Description = errors.Descriptions[errors.ErrInvalidClient]
	return response
}

func responseErrorHandler(re *errors.Response) {
	log.Info("Oauth server response Error:", re.Error.Error())
}

func passwordAuthorizationHandler(email string, password string) (userID string, err error) {
	user, err := models.GetUserByEmail(email)
	if err != nil {
		return "", errors.ErrInvalidClient
	}

	if helpers.CompareHashWithPassword(user.HashedPassword, password) {
		return fmt.Sprint(user.ID), nil
	}
	return "", errors.ErrInvalidClient
}
