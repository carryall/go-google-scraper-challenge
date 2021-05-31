package helpers

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

func GetAppRunMode() string {
	runMode, err := web.AppConfig.String("runmode")
	if err != nil {
		logs.Error("Run mode not found: ", err)
	}

	return runMode
}

func GetPaginationPerPage() int {
	perPage, err := web.AppConfig.Int("PaginationPerPage")
	if err != nil {
		logs.Error("Pagination per page not found: ", err)
	}

	return perPage
}

func GetRedisPoolNamespace() string {
	namespace, err := web.AppConfig.String("redisPoolNamespace")
	if err != nil {
		logs.Error("Redis pool namespace not found: ", err)
	}

	return namespace
}

func GetScrapingJobName() string {
	jobName, err := web.AppConfig.String("scrapingJobName")
	if err != nil {
		logs.Error("Job name not found: ", err)
	}

	return jobName
}
