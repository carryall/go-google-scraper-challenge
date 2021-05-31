package initializers

import (
	"go-google-scraper-challenge/database"
)

// SetUpDatabase setup database for the project
func SetUpDatabase() {
	database.SetupPostgresDB()
	database.SetupRedisPool()
}
