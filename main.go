package main

import (
	"log"

	"go-google-scraper-challenge/initializers"
	_ "go-google-scraper-challenge/routers"
	oauth_services "go-google-scraper-challenge/services/oauth"

	"github.com/beego/beego/v2/server/web"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	initializers.SetUpDatabase()
	initializers.SetUpTemplateFunction()
	initializers.SetupStaticPaths()
	initializers.SetLowercaseValidationErrors()
	oauth_services.SetUpOauth()
}

func main() {
	web.Run()
}
