package bootstrap

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

const sessionMaxAge = 60 * 60 * 24 * 3 // 3 days

func SetupSession(engine *gin.Engine) *gin.Engine {
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{MaxAge: sessionMaxAge})
	engine.Use(sessions.Sessions("google_scraper_session", store))

	return engine
}
