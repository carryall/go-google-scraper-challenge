package jobs

import (
	"go-google-scraper-challenge/models"
	"go-google-scraper-challenge/services/scraper"

	"github.com/beego/beego/v2/core/logs"
	"github.com/gocraft/work"
)

type Context struct{}

// Number of retry
const MaxFails = 3

func (c *Context) Perform(job *work.Job) error {
	resultId := job.ArgInt64("resultID")
	result, err := models.GetResultById(resultId)
	if err != nil {
		logs.Error("Finding keyword failed: ", err)
		return err
	}

	scraperService := scraper.Scraper{
		Result: result,
	}
	err = scraperService.Run()
	if err != nil {
		logs.Error("Failed to run scraper:", err.Error())

		return err
	}

	logs.Info("Finished scraping keyword:", result.Keyword, "for result ID:", result.Id)

	return nil
}
