package initializers

import (
	"go-google-scraper-challenge/helpers"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq"
)

// SetUpDatabase setup database for the project
func SetUpDatabase() {
	runMode := helpers.GetAppRunMode()
	orm.Debug = runMode == "dev"

	dbURL, err := web.AppConfig.String("dbUrl")
	if err != nil {
		logs.Error("Database URL not found: ", err)
	}

	err = orm.RegisterDriver("postgres", orm.DRPostgres)
	if err != nil {
		logs.Error("Postgres Driver registration failed: ", err)
	}

	err = orm.RegisterDataBase("default", "postgres", dbURL)
	if err != nil {
		logs.Error("Database Registration failed: ", err)
	}
}
