package tasks_test

import (
	"context"

	"go-google-scraper-challenge/tasks"

	"github.com/beego/beego/v2/task"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Scraping", func() {
	Describe("#GetScrapingTasker", func() {
		It("returns a runnable tasker", func() {
			tasker := tasks.GetScrapingTasker()

			Expect(tasker).To(BeAssignableToTypeOf(&task.Task{}))
			err := tasker.Run(context.Background())
			Expect(err).To(BeNil())
		})
	})
})
