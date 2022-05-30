package test

import (
	"os"
	"path/filepath"
	"runtime"

	"go-google-scraper-challenge/bootstrap"
	"go-google-scraper-challenge/database"
	"go-google-scraper-challenge/helpers/log"
	"go-google-scraper-challenge/lib/services/oauth"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func SetupTestEnvironment() {
	gin.SetMode(gin.TestMode)

	setRootDir()

	bootstrap.LoadConfig()

	database.InitDatabase(database.GetDatabaseURL())

	oauth.SetUpOauth()
}

func setRootDir() {
	_, currentFile, _, _ := runtime.Caller(0)
	root := filepath.Join(filepath.Dir(currentFile), "../")

	err := os.Chdir(root)
	if err != nil {
		log.Fatal("Failed to set root directory: ", err)
	}
}
