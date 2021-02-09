package tests

import (
	"github.com/onsi/ginkgo"
)

// LoginAs login using given email and password
func Login(email string, password string) {
	body := RequestBody(map[string]string{
		"email":    email,
		"password": password,
	})
	response := MakeRequest("POST", "/sessions", body)
	currentPath := GetCurrentPath(response)

	if currentPath != "/" {
		ginkgo.Fail("Failed to log user in")
	}
}
