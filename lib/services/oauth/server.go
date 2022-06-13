package oauth

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"go-google-scraper-challenge/database"
	"go-google-scraper-challenge/helpers"
	"go-google-scraper-challenge/helpers/log"
	"go-google-scraper-challenge/lib/models"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	pg "github.com/vgarvardt/go-oauth2-pg"
	"github.com/vgarvardt/go-pg-adapter/pgx4adapter"
	"gopkg.in/oauth2.v3"
	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/server"
)

var oauthServer *server.Server
var clientStore *pg.ClientStore
var tokenStore *pg.TokenStore

type TokenRequest struct {
	ClientID     string
	ClientSecret string
	UserID       string
}

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
	tokenStore, err = pg.NewTokenStore(adapter, pg.WithTokenStoreGCInterval(time.Minute))
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
	authServer.SetAllowedGrantType(oauth2.PasswordCredentials, oauth2.Refreshing)
	manager.SetRefreshTokenCfg(manage.DefaultRefreshTokenCfg)

	oauthServer = authServer
	clientStore = store
}

func ValidateToken(request *http.Request) (oauth2.TokenInfo, error) {
	return oauthServer.ValidationBearerToken(request)
}

func HandleTokenRequest(ctx *gin.Context) (tokenData map[string]interface{}, err error) {
	gt, tgr, err := oauthServer.ValidationTokenRequest(ctx.Request)
	if err != nil {
		return nil, err
	}

	ti, err := oauthServer.GetAccessToken(gt, tgr)
	if err != nil {
		return nil, err
	}

	tokenData = oauthServer.GetTokenData(ti)

	return tokenData, nil
}

// GenerateToken handle token request, will return error if fail
func GenerateToken(request *TokenRequest) (tokenInfo oauth2.TokenInfo, err error) {
	tokenRequest := &oauth2.TokenGenerateRequest{
		ClientID:     request.ClientID,
		ClientSecret: request.ClientSecret,
		UserID:       request.UserID,
	}

	return oauthServer.GetAccessToken(oauth2.PasswordCredentials, tokenRequest)
}

// GetClientStore returns OAuth client store
func GetClientStore() *pg.ClientStore {
	return clientStore
}

// GetTokenStore returns OAuth client store
func GetTokenStore() *pg.TokenStore {
	return tokenStore
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
