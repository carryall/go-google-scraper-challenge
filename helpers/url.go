package helpers

import (
	"net/url"

	"go-google-scraper-challenge/constants"
)

func IsActive(currentPath *url.URL, path string) bool {
	return currentPath.String() == path
}

func UrlFor(controllerName string, actionName string) string {
	return constants.WebRoutes[controllerName][actionName]
}
