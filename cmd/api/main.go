package main

import (
	"fmt"
	"os"

	"go-google-scraper-challenge/bootstrap"
	"go-google-scraper-challenge/database"
	"go-google-scraper-challenge/helpers/log"
	"go-google-scraper-challenge/lib/services/oauth"
	"go-google-scraper-challenge/view"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	bootstrap.LoadConfig()
	os.Setenv("TZ", "Asia/Bangkok")

	database.InitDatabase(database.GetDatabaseURL())

	engine := gin.Default()
	engine = bootstrap.SetupSession(engine)
	engine = bootstrap.SetupRouter(engine)

	oauth.SetUpOauth()
	bootstrap.InitCron()
	view.SetupView()

	err := engine.Run(getAppPort())
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}

func getAppPort() string {
	if gin.Mode() == gin.ReleaseMode {
		return fmt.Sprint(":", viper.GetString("PORT"))
	}

	return fmt.Sprint(":", viper.GetString("app_port"))
}
