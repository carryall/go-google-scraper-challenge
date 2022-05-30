package test

import (
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func CreateGinTestContext() (*gin.Context, *httptest.ResponseRecorder) {
	responseRecoder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(responseRecoder)

	return c, responseRecoder
}
