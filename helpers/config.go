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
