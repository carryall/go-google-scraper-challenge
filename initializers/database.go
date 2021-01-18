package initializers

import (
	"fmt"
	"log"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq"
)

// SetUpDatabase setup database for the project
func SetUpDatabase() {
	dbURL, err := web.AppConfig.String("db_url")
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
		fmt.Println("Database Sync failed: ", err)
	}
}
