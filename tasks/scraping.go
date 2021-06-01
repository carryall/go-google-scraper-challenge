package tasks

import (
	"context"

	"go-google-scraper-challenge/models"
	"go-google-scraper-challenge/services/scraper"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/task"
)

type ScrapingTask struct {
	Tasker *task.Task
}

func (t *ScrapingTask) Name() string {
	return "scraping"
}

func (t *ScrapingTask) GetTasker() *task.Task {
	if t.Tasker == nil {
		t.Tasker = task.NewTask(t.Name(), "0 * * * * *", perform)
	}

	return t.Tasker
}

func perform(_ context.Context) error {
	result, err := models.GetFirstPendingResult()
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
		return result.Fail()
	}

	return nil
}
