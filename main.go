package main

import (
	"go-google-scraper-challenge/initializers"
	_ "go-google-scraper-challenge/routers"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq"
)

func init() {
	initializers.SetUpDatabase()
}

func main() {
	web.Run()
}
