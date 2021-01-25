package initializers

import (
	"html/template"
	"io/ioutil"
	"log"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

// SetUpTemplateFunction register additional template functions
func SetUpTemplateFunction() {
	err := web.AddFuncMap("titlecase", toTitleCase)
	if err != nil {
		log.Fatal("Failed to add template function", err.Error())
	}

	err = web.AddFuncMap("render_file", renderFile)
	if err != nil {
		log.Fatal("Failed to add template function", err.Error())
	}
}

func toTitleCase(str string) string {
	return strings.Title(str)
}

func renderFile(path string) template.HTML {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return web.Str2html(string(content))
}
