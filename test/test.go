package test

import (
	"os"
	"path/filepath"
	"runtime"

	"go-google-scraper-challenge/bootstrap"
	"go-google-scraper-challenge/database"
	"go-google-scraper-challenge/helpers/log"
	"go-google-scraper-challenge/lib/services/oauth"
	"go-google-scraper-challenge/view"

	"github.com/gin-gonic/gin"
)

var Engine *gin.Engine

func SetupTestEnvironment() {
	gin.SetMode(gin.TestMode)
	os.Setenv("TZ", "Asia/Bangkok")

	setRootDir()

	bootstrap.LoadConfig()

	database.InitDatabase(database.GetDatabaseURL())

	engine := gin.Default()
	engine = bootstrap.SetupSession(engine)
	engine = bootstrap.SetupRouter(engine)
	Engine = engine

	oauth.SetUpOauth()
	view.SetupView()
}

func setRootDir() {
	_, currentFile, _, _ := runtime.Caller(0)
	root := filepath.Join(filepath.Dir(currentFile), "../")

	err := os.Chdir(root)
	if err != nil {
		log.Fatal("Failed to set root directory: ", err)
	}
}
