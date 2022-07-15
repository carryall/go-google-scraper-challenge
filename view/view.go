package view

import (
	"html/template"
	"io/fs"
	"path/filepath"
	"strings"

	"go-google-scraper-challenge/helpers"
	"go-google-scraper-challenge/helpers/log"

	"github.com/foolin/goview"
)

const ROOT_VIEW_PATH = "lib/web/views/"

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
			"assetsCSS":      helpers.AssetsCSS,
			"isActive":       helpers.IsActive,
			"renderFile":     helpers.RenderFile,
			"renderIcon":     helpers.RenderIcon,
			"urlFor":         helpers.UrlFor,
			"urlForID":       helpers.UrlForID,
			"toKebab":        helpers.ToKebabCase,
			"toTitle":        helpers.ToTitleCase,
			"formatDateTime": helpers.FormatDateTime,
		},
	}
}

func getPartialList() []string {
	partials := []string{}

	err := filepath.Walk(ROOT_VIEW_PATH, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			fileName := info.Name()
			fileExtension := filepath.Ext(fileName)

			if fileExtension == ".html" && strings.HasPrefix(fileName, "_") {
				filePath := strings.Split(path, ROOT_VIEW_PATH)[1]
				filePathWithoutExt := filePath[:len(filePath)-len(fileExtension)]
				partials = append(partials, filePathWithoutExt)
			}
		}

		return nil
	})

	if err != nil {
		log.Info("Fail to get partial files", err.Error())
	}

	return partials
}
