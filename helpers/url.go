package helpers

import (
	"fmt"
	"net/url"
	"strings"

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

func UrlForID(controllerName string, actionName string, id int64) string {
	controllerRoutes, ok := constants.WebRoutes[controllerName]
	if !ok {
		return ""
	}

	url, ok := controllerRoutes[actionName]
	if !ok {
		return ""
	}

	return strings.Replace(url, ":id", fmt.Sprint(id), -1)
}
