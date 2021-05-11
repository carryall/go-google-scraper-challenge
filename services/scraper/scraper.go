package scraper

import (
	"fmt"
	"log"
	"net/url"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
)

var selectors = map[string]string {
	"wholePage": "html",
	"topImageAds": "#tvcap .plantl a.pla-unit-title-link",
	"topLinkAds": "#tvcap .d5oMvf > a",
	"sideImageAds": ".cu-container a.plantl.clickable-card",
	"nonAds": "#search .yuRUbf > a",
	"bottomAds": "#bottomads .uEierd > a",
}

// TODO: Replace this struct with model
type Result struct {
	Keyword string
	TopAdLinks []string
	OtherAdLinks []string
	NonAdLinks []string
	PageCache string
}
var results = map[string]*Result{}

const GOOGLE_SEARCH_URL = "http://www.google.com/search?q=%s"

func Search(keywords []string) {
	collector := colly.NewCollector(colly.Async(true))

	searchQueue, err := queue.New(2, &queue.InMemoryQueueStorage{MaxSize: 10000})
	if err != nil {
		log.Fatal("Failed to create a queue", err.Error())
	}

	for _, k := range keywords {
		results[k] = &Result{Keyword: k}

		escapedKeyword := url.QueryEscape(k)
		err := searchQueue.AddURL(fmt.Sprintf(GOOGLE_SEARCH_URL, escapedKeyword))
		if err != nil {
			log.Println("Failed to add url to queue", err.Error())
		}
	}

	collector.OnRequest(RequestHandler)
	collector.OnResponse(ResponseHandler)
	collector.OnError(ErrorHandler)

	collector.OnHTML(selectors["wholePage"], func(e *colly.HTMLElement) {
		keyword := e.Request.Ctx.Get("keyword")
		results[keyword].PageCache = string(e.Response.Body)
	})

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

	collector.OnScraped(saveResult)

	err = searchQueue.Run(collector)
	if err != nil {
		log.Println("Failed to run the queue", err.Error())
	}
}

func RequestHandler (request *colly.Request) {
	usrAgent := RandomUserAgent()
	request.Headers.Set("User-Agent", usrAgent)

	log.Println("Visiting", request.URL)
	request.Ctx.Put("keyword", keywordFromUrl(request.URL.String()))
}

func ResponseHandler (response *colly.Response) {
	log.Println("Visited", response.Request.URL)
}

func ErrorHandler(response *colly.Response, err error) {
	log.Println("Failed to request URL:", response.Request.URL, "with response:", response, "\nError:", err)
}

func keywordFromUrl(urlStr string) (keyword string) {
	parsedUrl, err := url.Parse(urlStr)
	if err != nil {
		log.Println("Failed to parse url string", err.Error())
	}

	return parsedUrl.Query().Get("q")
}

func addResultLink(linkType string, element *colly.HTMLElement)  {
	keyword := element.Request.Ctx.Get("keyword")
	link := element.Attr("href")

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

func saveResult(response *colly.Response) {
	// TODO: add the result to database on another PR
	keyword := response.Request.Ctx.Get("keyword")
	result := results[keyword]

	log.Println("Finished scraping for keyword:", keyword, "==========")
	log.Println("Keyword:", result.Keyword)
	log.Println("	Top Ad:", len(result.TopAdLinks))
	log.Println("	Non Ad:", len(result.NonAdLinks))
	log.Println("	Other Ad:", len(result.OtherAdLinks))
}
