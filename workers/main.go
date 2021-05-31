package main

import (
	"os"
	"os/signal"
	"syscall"

	"go-google-scraper-challenge/database"
	"go-google-scraper-challenge/helpers"
	"go-google-scraper-challenge/initializers"
	oauth_services "go-google-scraper-challenge/services/oauth"
	"go-google-scraper-challenge/workers/jobs"

	"github.com/beego/beego/v2/core/logs"
	"github.com/gocraft/work"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logs.Error("Error loading .env file")
	}

	initializers.SetUpDatabase()
	initializers.SetUpTemplateFunction()
	initializers.SetupStaticPaths()
	initializers.SetModelDefaultValueFilter()
	initializers.SetLowercaseValidationErrors()

	oauth_services.SetUpOauth()
}

func main() {
	pool := work.NewWorkerPool(jobs.Context{}, 5, helpers.GetRedisPoolNamespace(), database.GetRedisPool())
	pool.JobWithOptions(helpers.GetScrapingJobName(), work.JobOptions{MaxFails: jobs.MaxFails}, (*jobs.Context).Perform)

	// Start processing jobs
	pool.Start()

	// Wait for a signal to quit:
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan

	// Stop the pool
	pool.Stop()
}
