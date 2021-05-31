package enqueuer

import (
	"go-google-scraper-challenge/database"
	"go-google-scraper-challenge/helpers"

	"github.com/beego/beego/v2/core/logs"
	"github.com/gocraft/work"
)

var enqueuer *work.Enqueuer

func init() {
	database.SetupRedisPool()
	enqueuer = work.NewEnqueuer(helpers.GetRedisPoolNamespace(), database.GetRedisPool())
}

// EnqueueScraping enqueues a scraping
func EnqueueScraping(resultId int64) error {
	job, err := enqueuer.EnqueueIn(helpers.GetScrapingJobName(), 3, work.Q{"resultID": resultId})
	if err != nil {
		return err
	}

	logs.Info("Enqueued ", job.Name, "job for resultID", job.ArgInt64("resultID"))

	return nil
}
