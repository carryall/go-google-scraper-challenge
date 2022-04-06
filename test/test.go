package test

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/nimblehq/google_scraper/bootstrap"
	"github.com/nimblehq/google_scraper/helpers/log"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func SetupTestEnvironment() {
	gin.SetMode(gin.TestMode)

	setRootDir()

	bootstrap.LoadConfig()

	bootstrap.InitDatabase()
}

func setRootDir() {
	_, currentFile, _, _ := runtime.Caller(0)
	root := filepath.Join(filepath.Dir(currentFile), "../")

	err := os.Chdir(root)
	if err != nil {
		log.Fatal("Failed to set root directory: ", err)
	}
}
