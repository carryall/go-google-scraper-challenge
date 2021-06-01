package initializers

import (
	"go-google-scraper-challenge/tasks"

	"github.com/beego/beego/v2/task"
)

func SetupTask() {
	scrapingTask := tasks.ScrapingTask{}
	task.AddTask(scrapingTask.Name(), scrapingTask.GetTasker())
	task.StartTask()
}
