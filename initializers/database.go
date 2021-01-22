package initializers

import (
	"log"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq"
)

// SetUpDatabase setup database for the project
func SetUpDatabase() {
	runMode, err := web.AppConfig.String("runmode")
	if err != nil {
		log.Fatal("Run mode not found: ", err)
	}

	dbURL, err := web.AppConfig.String("db_url")
	if err != nil {
		log.Fatal("Database URL not found: ", err)
	}

	err = orm.RegisterDriver("postgres", orm.DRPostgres)
	if err != nil {
		log.Fatal("Postgres Driver registration failed: ", err)
	}

	err = orm.RegisterDataBase("default", "postgres", dbURL)
	if err != nil {
		log.Fatal("Database Registration failed: ", err)
	}

	verbose := runMode == "dev"
	err = orm.RunSyncdb("default", false, verbose)
	if err != nil {
		log.Fatal("Database Sync failed: ", err)
	}
}
