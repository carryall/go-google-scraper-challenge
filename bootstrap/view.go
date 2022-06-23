package bootstrap

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/url"
	"path/filepath"

	"go-google-scraper-challenge/helpers/log"

	"github.com/foolin/goview"
)

const ROOT_VIEW_PATH = "lib/web/views"
const PARTIAL_PATH = ROOT_VIEW_PATH + "/partials"

func SetupView() {
	gv := goview.New(goview.Config{
		Root:      ROOT_VIEW_PATH,
		Extension: ".html",
		Master:    "layouts/default",
		Partials:  getPartialList(),
		Funcs: template.FuncMap{
			"isActive":   isActive,
			"renderFile": renderFile,
			"renderIcon": renderIcon,
		},
	})

	goview.Use(gv)
}

func getPartialList() []string {
	partials := []string{}
	files, err := ioutil.ReadDir(PARTIAL_PATH)
	if err != nil {
		log.Info("Fail to get partial files", err.Error())
	}

	for _, file := range files {
		fileName := file.Name()
		fileName = fileName[:len(fileName)-len(filepath.Ext(fileName))]
		partials = append(partials, fmt.Sprintf(`partials/%s`, fileName))
	}

	return partials
}

func isActive(currentPath *url.URL, path string) bool {
	return currentPath.String() == path
}

func renderFile(path string) template.HTML {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Error(err.Error())

		return template.HTML("")
	}

	return template.HTML(string(content))
}

func renderIcon(iconName string, classNames string) template.HTML {
	iconTemplate := `<svg class="icon ` + classNames + `" viewBox="0 0 20 20">
		<use xlink:href="/svg/sprite.symbol.svg#` + iconName + `" />
	</svg>`

	return template.HTML(iconTemplate)
}
