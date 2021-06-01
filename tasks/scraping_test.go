package tasks_test

import (
	"context"

	"go-google-scraper-challenge/tasks"

	"github.com/beego/beego/v2/task"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Scraping", func() {
	Describe("#Name", func() {
		It("returns the task name", func() {
			scrapingTask := tasks.ScrapingTask{}

			Expect(scrapingTask.Name()).To(Equal("scraping"))
		})
	})

	Describe("#GetTasker", func() {
		It("returns a runnable tasker", func() {
			scrapingTask := tasks.ScrapingTask{}
			tasker := scrapingTask.GetTasker()

			Expect(tasker).To(BeAssignableToTypeOf(&task.Task{}))
			err := tasker.Run(context.Background())
			Expect(err).To(BeNil())
		})
	})
})
