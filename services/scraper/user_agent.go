package scraper

import (
	"fmt"
	"math/rand"
	"strings"
)

var browserVersions = []string{
	"firefox_58.0",
	"firefox_57.0",
	"firefox_56.0",
	"firefox_52.0",
	"chrome_88.0.4324.192",
	"chrome_70.0.3538.77",
	"chrome_65.0.3325.146",
	"chrome_64.0.3282.0",
}

var osStrings = []string{
	"Macintosh; Intel Mac OS X 10_10",
	"Windows NT 10.0",
	"Windows NT 5.1",
	"Windows NT 6.1; WOW64",
	"Windows NT 6.1; Win64; x64",
	"X11; Linux x86_64",
}

const (
	firefoxPattern = "Mozilla/5.0 (%s; rv:%s) Gecko/20100101 Firefox/%s"
	chromePattern = "Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36"
)

func RandomUserAgent() string {
	os := osStrings[rand.Intn(len(osStrings))]
	uaPattern := ""
	version := strings.Split(browserVersions[rand.Intn(len(browserVersions))], "_")
	switch version[0] {
	case "firefox":
		uaPattern = firefoxPattern
	case "chrome":
		uaPattern = chromePattern
	}

	return fmt.Sprintf(uaPattern, os, version[1], version[1])
}
