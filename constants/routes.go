package constants

var WebRoutes = map[string]map[string]string{
	"session": {
		"new":    "/signin",
		"create": "/sessions",
		"delete": "/signout",
	},
	"users": {
		"new":    "/signup",
		"create": "/users",
	},
	"results": {
		"index":  "/",
		"create": "/results",
		"show":   "/results/:id",
		"cache":  "/results/:id/cache",
	},
}
