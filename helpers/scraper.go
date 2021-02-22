package helpers

import (
	"fmt"
	"log"
	"net/url"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	"github.com/gocolly/colly/extensions"
	"github.com/gocolly/colly/queue"
)

var selectors = map[string]string {
	"topImageAds": "#tvcap .plantl a.pla-unit-title-link",
	"topLinkAds": "#tvcap .d5oMvf > a",
	"sideImageAds": ".cu-container a.plantl.clickable-card",
	"nonAds": "#search .yuRUbf > a",
	"bottomAds": "#bottomads .uEierd > a",
}

const GOOGLE_SEARCH_URL = "http://www.google.com/search?q=%s"

func Scrape(keywords []string) {
	collector := colly.NewCollector(
			colly.Debugger(&debug.LogDebugger{}),
			colly.CacheDir("./cache/"),
			colly.Async(true),
		)

	q, _ := queue.New(
		2, // Number of consumer threads
		&queue.InMemoryQueueStorage{MaxSize: 10000}, // Use default queue storage
	)

	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*httpbin.*",
		Parallelism: 2,
		//Delay:      5 * time.Second,
	})

	for _, keyword := range keywords {
		excapedKeyword := url.QueryEscape(keyword)
		q.AddURL(fmt.Sprintf(GOOGLE_SEARCH_URL, excapedKeyword))
	}

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

	collector.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	q.Run(collector)
	//err := collector.Visit(GOOGLE_SEARCH_URL + keyword)
	//if err != nil {
	//	log.Fatal("Fail to start scraping ", err.Error())
	//}
}
