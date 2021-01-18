package main

import (
	"go-google-scraper-challenge/initializers"
	_ "go-google-scraper-challenge/routers"
	"log"

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
}

func main() {
	web.Run()
}
