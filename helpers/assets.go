package helpers

import "html/template"

func AssetsCSS(path string) template.HTML {
	linkHTML := `<link href="static/stylesheets/` + path + `" rel="stylesheet" type="text/css" />`

	return template.HTML(linkHTML)
}
