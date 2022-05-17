package bootstrap

import (
	"fmt"

	"go-google-scraper-challenge/helpers"
	"go-google-scraper-challenge/helpers/log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func InitDatabase(databaseURL string) {
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to %v database: %v", gin.Mode(), err)
	} else {
		viper.Set("database", db)
		database = db

		caser := cases.Title(language.English)
		log.Println(caser.String(gin.Mode()) + " database connected successfully.")
	}
}

func GetDB() *gorm.DB {
	if database == nil {
		InitDatabase(GetDatabaseURL())
	}

	return database
}

func GetDatabaseURL() string {
	if gin.Mode() == gin.ReleaseMode {
		return viper.GetString("DATABASE_URL")
	}

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		viper.GetString("db_username"),
		viper.GetString("db_password"),
		viper.GetString("db_host"),
		helpers.GetStringConfig("db_port"),
		helpers.GetStringConfig("db_name"),
	)
}
