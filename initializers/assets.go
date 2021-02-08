package initializers

import "github.com/beego/beego/v2/server/web"

// SetupStaticPaths set static paths for assets
func SetupStaticPaths() {
	web.SetStaticPath("/css", "static/css")
	web.SetStaticPath("/js", "static/js")
	web.SetStaticPath("/svg", "static/symbol/svg")
}
