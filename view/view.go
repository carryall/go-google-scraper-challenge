package view

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
const ICON_PATH = "static/images/icons/"

var viewEngines = map[string]*goview.ViewEngine{}

// var defaultConfig = goview.Config{
// 	Root:      ROOT_VIEW_PATH,
// 	Extension: ".html",
// 	Master:    "layouts/default",
// 	Partials:  getPartialList(),
// 	Funcs: template.FuncMap{
// 		"assetsCSS":  assetsCSS,
// 		"isActive":   isActive,
// 		"renderFile": renderFile,
// 		"renderIcon": renderIcon,
// 		"urlFor":     urlFor,
// 	},
// }

func SetupView() {
	defaultEngine := goview.New(getDefaultConfig())
	viewEngines["default"] = defaultEngine
	viewEngines["authentication"] = setupNewEngine("authentication")

	goview.Use(defaultEngine)
}

func SetLayout(layoutName string) {
	engine := viewEngines[layoutName]

	if engine != nil {
		engine = setupNewEngine(layoutName)
	}

	viewEngines[layoutName] = engine
	goview.Use(engine)
}

func setupNewEngine(layoutName string) *goview.ViewEngine {
	newConfig := getDefaultConfig()
	newConfig.Master = "layouts/" + layoutName

	return goview.New(newConfig)
}

func getDefaultConfig() goview.Config {
	return goview.Config{
		Root:      ROOT_VIEW_PATH,
		Extension: ".html",
		Master:    "layouts/default",
		Partials:  getPartialList(),
		Funcs: template.FuncMap{
			"assetsCSS":  assetsCSS,
			"isActive":   isActive,
			"renderFile": renderFile,
			"renderIcon": renderIcon,
			"urlFor":     urlFor,
		},
	}
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

func assetsCSS(path string) template.HTML {
	linkHTML := `<link href="static/stylesheets/` + path + `" rel="stylesheet" type="text/css" />`

	return template.HTML(linkHTML)
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
	iconPath := ICON_PATH + iconName + ".svg"
	iconTemplate := `<svg class="icon ` + classNames + `" viewBox="0 0 20 20">` + string(renderFile(iconPath)) + `</svg>`

	return template.HTML(iconTemplate)
}

func urlFor(controller string, action string) string {
	return ""
}
