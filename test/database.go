package test

import (
	"fmt"

	"go-google-scraper-challenge/database"
	"go-google-scraper-challenge/helpers/log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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
		caser := cases.Title(language.English)
		log.Println(caser.String(gin.Mode()) + " database connected successfully.")
	}

	db.Exec(truncateSQL)

	if err != nil {
		log.Warn("FAILED TO TRUNCATE TABLE", tableNames, err.Error())
	}
}
