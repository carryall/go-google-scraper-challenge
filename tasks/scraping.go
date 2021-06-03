package tasks

import (
	"context"

	"go-google-scraper-challenge/models"
	"go-google-scraper-challenge/services/scraper"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/task"
)

const ScrapingTaskName = "scraping"

func GetScrapingTasker() *task.Task {
	return task.NewTask(ScrapingTaskName, "0 * * * * *", perform)
}

func perform(_ context.Context) error {
	result, err := models.GetOldestPendingResult()
	if err != nil {
		logs.Info("No pending result", err.Error())

		return nil
	}

	logs.Info("Performing scraping task with result ID:", result.Id)
	scraperService := scraper.Scraper{
		Result: result,
	}
	err = scraperService.Run()
	if err != nil {
		err = models.UpdateResultStatus(result, models.ResultStatusFailed)
		if err != nil {
			logs.Error("Failed to update result status", err.Error())
		}

		return err
	}

	return nil
}
