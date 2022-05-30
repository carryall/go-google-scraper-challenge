package helpers

import (
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
var matchSnake = regexp.MustCompile("([a-z0-9])_([a-z0-9])")
var matchSentence = regexp.MustCompile("([A-Za-z0-9]) ([A-Za-z0-9])")
var matchKebab = regexp.MustCompile("([a-z0-9])-([a-z0-9])")

// ToKebabCase convert string from camel/snake/sentence case to kebab case
func ToKebabCase(str string) string {
	kebab := matchFirstCap.ReplaceAllString(str, "${1}-${2}")
	kebab = matchAllCap.ReplaceAllString(kebab, "${1}-${2}")
	kebab = matchSnake.ReplaceAllString(kebab, "${1}-${2}")
	kebab = matchSentence.ReplaceAllString(kebab, "${1}-${2}")

	return strings.ToLower(kebab)
}

// ToSentenceCase convert string from camel/snake/kebab case to kebab case
func ToSentenceCase(str string) string {
	sentence := matchFirstCap.ReplaceAllString(str, "${1} ${2}")
	sentence = matchAllCap.ReplaceAllString(sentence, "${1} ${2}")
	sentence = matchSnake.ReplaceAllString(sentence, "${1} ${2}")
	sentence = matchKebab.ReplaceAllString(sentence, "${1} ${2}")
	sentence = strings.ToLower(sentence)
	words := strings.Split(sentence, " ")

	caser := cases.Title(language.English)
	return strings.Replace(sentence, words[0], caser.String(words[0]), 1)
}
