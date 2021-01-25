package initializers

import (
	"go-google-scraper-challenge/helpers"
	"html/template"
	"io/ioutil"
	"log"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

// SetUpTemplateFunction register additional template functions
func SetUpTemplateFunction() {
	err := web.AddFuncMap("titlecase", strings.Title)
	if err != nil {
		log.Fatal("Failed to add template function", err.Error())
	}

	err = web.AddFuncMap("sentencecase", helpers.ToSentenceCase)
	if err != nil {
		log.Fatal("Failed to add template function", err.Error())
	}

	err = web.AddFuncMap("render_file", renderFile)
	if err != nil {
		log.Fatal("Failed to add template function", err.Error())
	}

	err = web.AddFuncMap("render_icon", renderIcon)
	if err != nil {
		log.Fatal("Failed to add template function", err.Error())
	}
}

func renderFile(path string) template.HTML {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return web.Str2html(string(content))
}

func renderIcon(iconName string) template.HTML {
	iconTemplate := `<svg class="icon" viewBox="0 0 20 20">
		<use xlink:href="svg/sprite.symbol.svg#` + iconName + `" />
	</svg>`

	return web.Str2html(iconTemplate)
}
