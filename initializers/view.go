package initializers

import (
	"html/template"
	"io/ioutil"
	"log"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

func SetUpTemplateFunction() {
	web.AddFuncMap("titlecase", func(str string) string {
		return strings.Title(str)
	})

	web.AddFuncMap("render_file", func(path string) template.HTML {
		content, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatal(err)
		}

		return web.Str2html(string(content))
	})
}
