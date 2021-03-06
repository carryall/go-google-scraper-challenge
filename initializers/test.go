package initializers

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"go-google-scraper-challenge/services/oauth"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/joho/godotenv"
)

// SetupTestEnvironment setup environment for testing
func SetupTestEnvironment() {
	appRoot := AppRootDir()

	logs.SetLevel(logs.LevelWarning)

	SetWorkingDirectory(appRoot)
	OverloadTestConfig()
	SetUpTemplateFunction()
	web.TestBeegoInit(appRoot)
	SetUpDatabase()
	SetupStaticPaths()
	SetModelDefaultValueFilter()
	SetLowercaseValidationErrors()
	oauth.SetUpOauth()
	SetupTask()
}

// AppRootDir returns the app root path of the project
func AppRootDir() string {
	_, currentFile, _, _ := runtime.Caller(0)
	currentFilePath := path.Join(path.Dir(currentFile))
	return filepath.Dir(currentFilePath)
}

// SetWorkingDirectory set current working directory
func SetWorkingDirectory(dir string) {
	err := os.Chdir(dir)
	if err != nil {
		logs.Error("Failed to set working directory", err.Error())
	}
}

// OverloadTestConfig load and override environment variables for test
func OverloadTestConfig() {
	err := godotenv.Overload(".env.test")
	if err != nil {
		logs.Error("Failed to overload test environment file", err.Error())
	}
}

// CleanupDatabase cleanup the given database table
func CleanupDatabase(tableNames []string) {
	truncateSQL := ""
	for _, t := range tableNames {
		truncateSQL += fmt.Sprintf("TRUNCATE TABLE \"%s\" CASCADE;", t)
	}

	ormer := orm.NewOrm()
	_, err := ormer.Raw(truncateSQL).Exec()
	if err != nil {
		logs.Warn("FAILED TO TRUNCATE TABLE", tableNames, err.Error())
		err := orm.RunSyncdb("default", true, false)
		if err != nil {
			logs.Error("Failed to sync database", err)
		}
	}
}
