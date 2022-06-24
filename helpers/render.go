package helpers

import (
	"html/template"
	"io/ioutil"

	"go-google-scraper-challenge/helpers/log"
)

const ICON_PATH = "static/images/icons/"

func RenderFile(path string) template.HTML {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Error(err.Error())

		return template.HTML("")
	}

	return template.HTML(string(content))
}

func RenderIcon(iconName string, classNames string) template.HTML {
	iconPath := ICON_PATH + iconName + ".svg"
	iconTemplate := `<svg class="icon ` + classNames + `" viewBox="0 0 20 20">` + string(RenderFile(iconPath)) + `</svg>`

	return template.HTML(iconTemplate)
}
