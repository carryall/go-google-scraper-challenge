package test

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/onsi/ginkgo"
)

func CreateGinTestContext() (*gin.Context, *httptest.ResponseRecorder) {
	responseRecoder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(responseRecoder)

	return c, responseRecoder
}

func GetCurrentPath(response *http.Response) string {
	url, err := response.Location()
	if err != nil {
		ginkgo.Fail("Failed to get current path: " + err.Error())
	}
	return url.Path
}
