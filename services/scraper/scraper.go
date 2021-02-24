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

const GOOGLE_SEARCH_URL = "http://www.google.com/search?q=%s"
//const currentBrowser = "Chrome/88.0.4324.182"
//const currentOs = "Macintosh; Intel Mac OS X 10_15_5"

func Search(keywords []string) {
	//runMode, err := web.AppConfig.String("runmode")
	//if err != nil {
	//	log.Fatal("Run mode not found: ", err)
	//}

	results := map[string] *SearchResult{}
	collector := colly.NewCollector(colly.Async(true))
	//extensions.RandomUserAgent(collector)

	//if runMode == "dev" {
	//	collector.SetDebugger(&debug.LogDebugger{})
	//}

	q, err := queue.New(2, &queue.InMemoryQueueStorage{MaxSize: 10000})
	if err != nil {
		log.Fatal("Failed to create a queue", err.Error())
	}

	//collector.Limit(&colly.LimitRule{
	//	DomainGlob:  "*httpbin.*",
	//	Parallelism: 2,
	//})

	for _, keyword := range keywords {
		results[keyword] = &SearchResult{Keyword: keyword}

		escapedKeyword := url.QueryEscape(keyword)
		err := q.AddURL(fmt.Sprintf(GOOGLE_SEARCH_URL, escapedKeyword))
		if err != nil {
			log.Println("Failed to add url to queue", err.Error())
		}
	}
	q.IsEmpty()

	collector.OnRequest(RequestHandler)
	collector.OnResponse(ResponseHandler)
	collector.OnError(ErrorHandler)

	collector.OnHTML(selectors["nonAds"], func(e *colly.HTMLElement) {
		keyword := e.Request.Ctx.Get("keyword")
		link := e.Attr("href")
		if len(link) > 0 {
			results[keyword].NonAdLinks = append(results[keyword].NonAdLinks, link)
		}
	})

	collector.OnHTML(selectors["topImageAds"], func(e *colly.HTMLElement) {
		keyword := e.Request.Ctx.Get("keyword")
		link := e.Attr("href")
		if len(link) > 0 {
			results[keyword].TopAdLinks = append(results[keyword].TopAdLinks, link)
		}
	})

	collector.OnHTML(selectors["topLinkAds"], func(e *colly.HTMLElement) {
		keyword := e.Request.Ctx.Get("keyword")
		link := e.Attr("href")
		if len(link) > 0 {
			results[keyword].TopAdLinks = append(results[keyword].TopAdLinks, link)
		}
	})

	collector.OnHTML(selectors["sideImageAds"], func(e *colly.HTMLElement) {
		keyword := e.Request.Ctx.Get("keyword")
		link := e.Attr("href")
		if len(link) > 0 {
			results[keyword].OtherAdLinks = append(results[keyword].OtherAdLinks, link)
		}
	})

	collector.OnHTML(selectors["bottomAds"], func(e *colly.HTMLElement) {
		keyword := e.Request.Ctx.Get("keyword")
		link := e.Attr("href")
		if len(link) > 0 {
			results[keyword].OtherAdLinks = append(results[keyword].OtherAdLinks, link)
		}
	})

	collector.OnScraped(func(r *colly.Response) {
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
	//userAgent := fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) %s Safari/537.36", currentOs, currentBrowser)
	//r.Headers.Set("User-Agent", userAgent)

	ua := RandomUserAgent()
	log.Println("========USER AGENT========", ua)
	r.Headers.Set("User-Agent", ua)
	//ua := r.Headers.Get("User-Agent")
	//log.Println("========USER AGENT========", ua)

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
