package bootstrap

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func SetupSession(engine *gin.Engine) *gin.Engine {
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{MaxAge: 60 * 60 * 24 * 3})
	engine.Use(sessions.Sessions("mysession", store))

	return engine
}
