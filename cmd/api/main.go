package main

import (
	"fmt"

	"go-google-scraper-challenge/bootstrap"
	"go-google-scraper-challenge/database"
	"go-google-scraper-challenge/helpers/log"
	"go-google-scraper-challenge/lib/services/oauth"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	bootstrap.LoadConfig()

	database.InitDatabase(database.GetDatabaseURL())

	r := bootstrap.SetupRouter()

	oauth.SetUpOauth()

	err := r.Run(getAppPort())
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
