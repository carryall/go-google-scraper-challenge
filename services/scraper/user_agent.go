package scraper

import (
	"fmt"
	"math/rand"
)

var userAgents = []func() string{
	getFirefoxUA,
	getChromeUA,
}

var ffVersions = []float32{
	58.0,
	57.0,
	56.0,
	52.0,
}

var chromeVersions = []string{
	"88.0.4324.192",
	"70.0.3538.77",
	"65.0.3325.146",
	"64.0.3282.0",
}

var osStrings = []string{
	"Macintosh; Intel Mac OS X 10_10",
	"Windows NT 10.0",
	"Windows NT 5.1",
	"Windows NT 6.1; WOW64",
	"Windows NT 6.1; Win64; x64",
	"X11; Linux x86_64",
}

func RandomUserAgent() string {
	return userAgents[rand.Intn(len(userAgents))]()
}

func getFirefoxUA() string {
	version := ffVersions[rand.Intn(len(ffVersions))]
	os := osStrings[rand.Intn(len(osStrings))]
	return fmt.Sprintf("Mozilla/5.0 (%s; rv:%.1f) Gecko/20100101 Firefox/%.1f", os, version, version)
}

func getChromeUA() string {
	version := chromeVersions[rand.Intn(len(chromeVersions))]
	os := osStrings[rand.Intn(len(osStrings))]
	return fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36", os, version)
}
