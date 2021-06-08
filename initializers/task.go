package initializers

import (
	"go-google-scraper-challenge/tasks"

	"github.com/beego/beego/v2/task"
)

func SetupTask() {
	task.AddTask(tasks.ScrapingTaskName, tasks.GetScrapingTasker())
	task.StartTask()
}
