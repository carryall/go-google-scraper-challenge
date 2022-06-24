package helpers

import "net/url"

func IsActive(currentPath *url.URL, path string) bool {
	return currentPath.String() == path
}

func UrlFor(controller string, action string) string {
	return ""
}
