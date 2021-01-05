package main

import (
	"fmt"
	_ "go-google-scraper-challenge/routers"
	"log"

	_ "github.com/lib/pq"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	dbURL, err := beego.AppConfig.String("db_url")
	if err != nil {
		log.Fatal("Database URL not found: ", err)
	}

	err = orm.RegisterDriver("postgres", orm.DRPostgres)
	if err != nil {
		fmt.Println("Postgres Driver registration failed: ", err)
	}

	err = orm.RegisterDataBase("default", "postgres", dbURL)
	if err != nil {
		fmt.Println("Database Registration failed: ", err)
	}

	// TODO: register models after this
}

func main() {
	beego.Run()
}
