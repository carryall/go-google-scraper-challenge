package bootstrap

import (
	apiv1router "go-google-scraper-challenge/lib/api/v1/routers"
	webrouter "go-google-scraper-challenge/lib/web/routers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	apiv1router.ComebineRoutes(r)
	webrouter.ComebineRoutes(r)

	return r
}
