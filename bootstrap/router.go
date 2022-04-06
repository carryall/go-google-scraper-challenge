package bootstrap

import (
	apiv1router "github.com/nimblehq/google_scraper/lib/api/v1/routers"
	webrouter "github.com/nimblehq/google_scraper/lib/web/routers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	apiv1router.ComebineRoutes(r)
	webrouter.ComebineRoutes(r)

	return r
}
