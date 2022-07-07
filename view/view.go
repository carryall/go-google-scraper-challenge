package view

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"path/filepath"

	"go-google-scraper-challenge/helpers"
	"go-google-scraper-challenge/helpers/log"

	"github.com/foolin/goview"
)

const ROOT_VIEW_PATH = "lib/web/views"
const PARTIAL_PATH = ROOT_VIEW_PATH + "/partials"

var viewEngines = map[string]*goview.ViewEngine{}

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
			"assetsCSS":  helpers.AssetsCSS,
			"isActive":   helpers.IsActive,
			"renderFile": helpers.RenderFile,
			"renderIcon": helpers.RenderIcon,
			"urlFor":     helpers.UrlFor,
			"toKebab":    helpers.ToKebabCase,
			"toTitle":    helpers.ToTitleCase,
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
