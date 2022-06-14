package bootstrap

import (
	"go-google-scraper-challenge/helpers/log"
	"go-google-scraper-challenge/lib/models"
	"go-google-scraper-challenge/lib/services/scraper"
	"time"

	"github.com/go-co-op/gocron"
)

var scheduler *gocron.Scheduler

func InitCron() {
	scheduler = gocron.NewScheduler(time.UTC)

	_, err := scheduler.Every(30).Seconds().Do(performScrape)
	if err != nil {
		log.Errorln("Fail to setup cron ", err.Error())

		return
	}

	scheduler.StartAsync()
}

func performScrape() error {
	result, err := models.GetOldestPendingResult()
	if err != nil {
		log.Infoln("No pending result:", err.Error())

		return nil
	}

	log.Infoln("Performing scraping task with result ID:", result.ID)
	scraperService := scraper.Scraper{
		Result: result,
	}
	err = scraperService.Run()
	if err != nil {
		err = models.UpdateResultStatus(result, models.ResultStatusFailed)
		if err != nil {
			log.Error("Failed to update result status:", err.Error())
		}

		return err
	}

	return nil
}
