package scraper

import (
	"fmt"
	"log"
	"net/url"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
)

var selectors = map[string]string {
	"topImageAds": "#tvcap .plantl a.pla-unit-title-link",
	"topLinkAds": "#tvcap .d5oMvf > a",
	"sideImageAds": ".cu-container a.plantl.clickable-card",
	"nonAds": "#search .yuRUbf > a",
	"bottomAds": "#bottomads .uEierd > a",
}

// TODO: Replace this struct with model
type SearchResult struct {
	Keyword string
	TopAdLinks []string
	OtherAdLinks []string
	NonAdLinks []string
	PageCache string
}
var results = map[string]*SearchResult{}

const GOOGLE_SEARCH_URL = "http://www.google.com/search?q=%s"

func Search(keywords []string) {
	collector := colly.NewCollector(colly.Async(true))

	q, err := queue.New(2, &queue.InMemoryQueueStorage{MaxSize: 10000})
	if err != nil {
		log.Fatal("Failed to create a queue", err.Error())
	}

	for _, keyword := range keywords {
		results[keyword] = &SearchResult{Keyword: keyword}

		escapedKeyword := url.QueryEscape(keyword)
		err := q.AddURL(fmt.Sprintf(GOOGLE_SEARCH_URL, escapedKeyword))
		if err != nil {
			log.Println("Failed to add url to queue", err.Error())
		}
	}

	collector.OnRequest(RequestHandler)
	collector.OnResponse(ResponseHandler)
	collector.OnError(ErrorHandler)

	collector.OnHTML(selectors["nonAds"], func(e *colly.HTMLElement) {
		addResultLink("nonAd", e)
	})

	collector.OnHTML(selectors["topImageAds"], func(e *colly.HTMLElement) {
		addResultLink("topAd", e)
	})

	collector.OnHTML(selectors["topLinkAds"], func(e *colly.HTMLElement) {
		addResultLink("topAd", e)
	})

	collector.OnHTML(selectors["sideImageAds"], func(e *colly.HTMLElement) {
		addResultLink("otherAd", e)
	})

	collector.OnHTML(selectors["bottomAds"], func(e *colly.HTMLElement) {
		addResultLink("otherAd", e)
	})

	collector.OnScraped(func(r *colly.Response) {
		// TODO: add the result to database on another PR
		if q.IsEmpty() {
			log.Println("Finished scraping ==========")
			for _, result := range results {
				log.Println("Keyword:", result.Keyword)
				log.Println("	Top Ad:", len(result.TopAdLinks))
				log.Println("	Non Ad:", len(result.NonAdLinks))
				log.Println("	Other Ad:", len(result.OtherAdLinks))
			}
		}
	})

	err = q.Run(collector)
	if err != nil {
		log.Println("Failed to run the queue", err.Error())
	}
}

func RequestHandler (r *colly.Request) {
	ua := RandomUserAgent()
	r.Headers.Set("User-Agent", ua)

	log.Println("Visiting", r.URL)
	r.Ctx.Put("keyword", keywordFromUrl(r.URL.String()))
}

func ResponseHandler (r *colly.Response) {
	log.Println("Visited", r.Request.URL)
}

func ErrorHandler(r *colly.Response, err error) {
	log.Println("Failed to request URL:", r.Request.URL, "with response:", r, "\nError:", err)
}

func keywordFromUrl(urlStr string) (keyword string) {
	u, err := url.Parse(urlStr)
	if err != nil {
		log.Println("Failed to parse url string", err.Error())
	}

	return u.Query().Get("q")
}

func addResultLink(linkType string, e *colly.HTMLElement)  {
	keyword := e.Request.Ctx.Get("keyword")
	link := e.Attr("href")

	result := results[keyword]
	if len(link) > 0 {
		switch linkType {
		case "nonAd":
			result.NonAdLinks = append(result.NonAdLinks, link)
		case "topAd":
			result.TopAdLinks = append(result.TopAdLinks, link)
		case "otherAd":
			result.OtherAdLinks = append(result.OtherAdLinks, link)
		}
	}
}
