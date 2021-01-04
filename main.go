package main

import (
	_ "go-google-scraper-challenge/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}
