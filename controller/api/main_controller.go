package api

import(
	"github.com/kataras/iris/v12/mvc"
)

type MainController struct {}

func (m *MainController) Get() mvc.Response  {
	return mvc.Response {
		ContentType: "application/json",
		Text:        "{message: 'Hello World'}",
	}
}
