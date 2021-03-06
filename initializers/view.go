package initializers

import (
	"html/template"
	"io/ioutil"
	"strings"
	"time"

	"go-google-scraper-challenge/helpers"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

// SetUpTemplateFunction register additional template functions
func SetUpTemplateFunction() {
	templateFunctions := map[string]interface{}{
		"title_case":    strings.Title,
		"sentence_case": helpers.ToSentenceCase,
		"render_file":   renderFile,
		"render_icon":   renderIcon,
		"format_datetime": formatDateTime,
	}

	for n, fn := range templateFunctions {
		err := web.AddFuncMap(n, fn)
		if err != nil {
			logs.Error("Failed to add template function", err.Error())
		}
	}
}

func renderFile(path string) template.HTML {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		logs.Error(err)
	}

	return web.Str2html(string(content))
}

func renderIcon(iconName string, classNames string) template.HTML {
	iconTemplate := `<svg class="icon `+ classNames +`" viewBox="0 0 20 20">
		<use xlink:href="/svg/sprite.symbol.svg#` + iconName + `" />
	</svg>`

	return web.Str2html(iconTemplate)
}

func formatDateTime(dateTime time.Time) string {
	return web.Date(dateTime.Local(), "d/m/y H:i:s")
}
