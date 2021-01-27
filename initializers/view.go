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
	templateFunctions := map[string]interface{}{
		"title_case":    strings.Title,
		"sentence_case": helpers.ToSentenceCase,
		"render_file":   renderFile,
		"render_icon":   renderIcon,
	}

	for name, fn := range templateFunctions {
		err := web.AddFuncMap(name, fn)
		if err != nil {
			log.Fatal("Failed to add template function", err.Error())
		}
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
