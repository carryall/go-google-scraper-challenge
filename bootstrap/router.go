package bootstrap

import (
	apiv1router "go-google-scraper-challenge/lib/api/v1/routers"
	webrouter "go-google-scraper-challenge/lib/web/routers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(engine *gin.Engine) *gin.Engine {
	apiv1router.ComebineRoutes(engine)
	webrouter.ComebineRoutes(engine)

	return engine
}
