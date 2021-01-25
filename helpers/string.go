package helpers

import (
	"regexp"
	"strings"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
var matchSnake = regexp.MustCompile("([a-z0-9])_([a-z0-9])")

// ToKebabCase convert string from camel case to kebab case
func ToKebabCase(str string) string {
	kebab := matchFirstCap.ReplaceAllString(str, "${1}-${2}")
	kebab = matchAllCap.ReplaceAllString(kebab, "${1}-${2}")
	kebab = matchSnake.ReplaceAllString(kebab, "${1}-${2}")

	return strings.ToLower(kebab)
}
