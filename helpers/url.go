package helpers

import (
	"net/url"

	"go-google-scraper-challenge/constants"
)

func IsActive(currentPath *url.URL, path string) bool {
	return currentPath.String() == path
}

func UrlFor(controllerName string, actionName string) string {
	controllerRoutes, ok := constants.WebRoutes[controllerName]
	if !ok {
		return ""
	}

	url, ok := controllerRoutes[actionName]
	if !ok {
		return ""
	}

	return url
}
