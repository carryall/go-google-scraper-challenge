package test

import (
	"fmt"
	database "go-google-scraper-challenge/bootstrap"
	"go-google-scraper-challenge/helpers/log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CleanupDatabase(tableNames []string) {
	truncateSQL := ""
	for _, t := range tableNames {
		truncateSQL += fmt.Sprintf("TRUNCATE TABLE \"%s\" CASCADE;", t)
	}

	db, err := gorm.Open(postgres.Open(database.GetDatabaseURL()), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to %v database: %v", gin.Mode(), err)
	} else {
		viper.Set("database", db)
		log.Println(strings.Title(gin.Mode()) + " database connected successfully.")
	}

	db.Exec(truncateSQL)

	// ormer := orm.NewOrm()
	// _, err := ormer.Raw(truncateSQL).Exec()
	if err != nil {
		log.Warn("FAILED TO TRUNCATE TABLE", tableNames, err.Error())
		// err := orm.RunSyncdb("default", true, false)
		// if err != nil {
		// 	log.Error("Failed to sync database", err)
		// }
	}
}
