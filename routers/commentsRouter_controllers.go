package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["go-google-scraper-challenge/controllers/api:AuthController"] = append(beego.GlobalControllerRouter["go-google-scraper-challenge/controllers/api:AuthController"],
        beego.ControllerComments{
            Method: "Login",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-google-scraper-challenge/controllers:OAuthClientController"] = append(beego.GlobalControllerRouter["go-google-scraper-challenge/controllers:OAuthClientController"],
        beego.ControllerComments{
            Method: "New",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-google-scraper-challenge/controllers:OAuthClientController"] = append(beego.GlobalControllerRouter["go-google-scraper-challenge/controllers:OAuthClientController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-google-scraper-challenge/controllers:OAuthClientController"] = append(beego.GlobalControllerRouter["go-google-scraper-challenge/controllers:OAuthClientController"],
        beego.ControllerComments{
            Method: "Show",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-google-scraper-challenge/controllers:SessionController"] = append(beego.GlobalControllerRouter["go-google-scraper-challenge/controllers:SessionController"],
        beego.ControllerComments{
            Method: "New",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-google-scraper-challenge/controllers:SessionController"] = append(beego.GlobalControllerRouter["go-google-scraper-challenge/controllers:SessionController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-google-scraper-challenge/controllers:SessionController"] = append(beego.GlobalControllerRouter["go-google-scraper-challenge/controllers:SessionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-google-scraper-challenge/controllers:UserController"] = append(beego.GlobalControllerRouter["go-google-scraper-challenge/controllers:UserController"],
        beego.ControllerComments{
            Method: "New",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-google-scraper-challenge/controllers:UserController"] = append(beego.GlobalControllerRouter["go-google-scraper-challenge/controllers:UserController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
