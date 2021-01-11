package main

import (
	"fmt"
	"log"

	_ "go-google-scraper-challenge/routers"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq"
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

	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	beego.Run()
}
