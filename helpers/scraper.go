package helpers

import (
	"log"
	"net/url"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

var selectors = map[string]string {
	"topImageAds": "#tvcap .plantl a.pla-unit-title-link",
	"topLinkAds": "#tvcap .d5oMvf > a",
	"sideImageAds": ".commercial-unit-desktop-rhs a.pla-unit-single-clickable-target",
	"nonAds": "#search .yuRUbf > a",
	"bottomAds": "#bottomads .uEierd > a",
}

func Scrape(keyword string) {
	keyword = url.QueryEscape(keyword)
	collector := colly.NewCollector()

	extensions.RandomUserAgent(collector)

	collector.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	collector.OnResponse(func(r *colly.Response) {
		log.Println("Visited", r.Request.URL)
	})

	collector.OnHTML(selectors["nonAds"], func(e *colly.HTMLElement) {
		log.Println("NON ADs ===============")
		href := e.Attr("href")
		if len(href) > 0 {
			log.Println("	Title", e.Text)
			log.Println("	Link", e.Attr("href"))
		}
	})

	collector.OnHTML(selectors["topImageAds"], func(e *colly.HTMLElement) {
		log.Println("TOP IMAGE ADs ===============")
		href := e.Attr("href")
		if len(href) > 0 {
			log.Println("	Title", e.Text)
			log.Println("	Link", e.Attr("href"))
		}
	})

	collector.OnHTML(selectors["topLinkAds"], func(e *colly.HTMLElement) {
		log.Println("TOP LINK ADs ===============")
		href := e.Attr("href")
		if len(href) > 0 {
			log.Println("	Title", e.Text)
			log.Println("	Link", e.Attr("href"))
		}
	})

	collector.OnHTML(selectors["sideImageAds"], func(e *colly.HTMLElement) {
		log.Println("SIDE IMAGE ADs ===============")
		href := e.Attr("href")
		if len(href) > 0 {
			log.Println("	Title", e.Text)
			log.Println("	Link", e.Attr("href"))
		}
	})

	collector.OnHTML(selectors["bottomAds"], func(e *colly.HTMLElement) {
		log.Println("BOTTOM ADs ===============")
		href := e.Attr("href")
		if len(href) > 0 {
			log.Println("	Title", e.Text)
			log.Println("	Link", e.Attr("href"))
		}
	})

	err := collector.Visit("http://www.google.com/search?q="+keyword)
	if err != nil {
		log.Fatal("Fail to start scraping ", err.Error())
	}
}
