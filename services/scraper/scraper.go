package scraper

import (
	"fmt"
	"log"
	"net/url"

	"go-google-scraper-challenge/models"
	"go-google-scraper-challenge/models/adlinks"
	"go-google-scraper-challenge/models/results"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
	"github.com/siddontang/go/num"
)

var selectors = map[string]string {
	"wholePage": "html",
	"topImageAds": "#tvcap .plantl a.pla-unit-title-link",
	"topLinkAds": "#tvcap .d5oMvf > a",
	"sideImageAds": ".cu-container a.plantl.clickable-card",
	"nonAds": "#search .yuRUbf > a",
	"bottomAds": "#bottomads .uEierd > a",
}

var scrapingResults = map[string]int64{}

const GOOGLE_SEARCH_URL = "http://www.google.com/search?q=%s"

func Search(keywords []string, user *models.User) {
	collector := colly.NewCollector(colly.Async(true))

	searchQueue, err := queue.New(2, &queue.InMemoryQueueStorage{MaxSize: 10000})
	if err != nil {
		log.Fatal("Failed to create a queue", err.Error())
	}

	for _, k := range keywords {
		result := &models.Result{
			User: user,
			Keyword: k,
		}
		resultID, err := models.CreateResult(result)
		if err != nil {
			log.Fatal("Failed to create result", err.Error())
		}
		scrapingResults[k] = resultID

		escapedKeyword := url.QueryEscape(k)
		err = searchQueue.AddURL(fmt.Sprintf(GOOGLE_SEARCH_URL, escapedKeyword))
		if err != nil {
			log.Println("Failed to add url to queue", err.Error())
		}
	}

	collector.OnRequest(RequestHandler)
	collector.OnResponse(ResponseHandler)
	collector.OnError(ErrorHandler)

	collector.OnHTML(selectors["wholePage"], func(e *colly.HTMLElement) {
		result := ResultFromContext(e.Request.Ctx)
		result.Status = results.Processing
		result.PageCache = string(e.Response.Body)
		err = models.UpdateResultById(result)
		if err != nil {
			log.Fatal("Failed to update result page cache", err.Error())
		}
	})

	collector.OnHTML(selectors["nonAds"], func(e *colly.HTMLElement) {
		addNonAdLinkToResult(e)
	})

	collector.OnHTML(selectors["topImageAds"], func(e *colly.HTMLElement) {
		addAdLinkToResult(adlinks.Image, adlinks.Top, e)
	})

	collector.OnHTML(selectors["topLinkAds"], func(e *colly.HTMLElement) {
		addAdLinkToResult(adlinks.Link, adlinks.Top, e)
	})

	collector.OnHTML(selectors["sideImageAds"], func(e *colly.HTMLElement) {
		addAdLinkToResult(adlinks.Side, adlinks.Image, e)
	})

	collector.OnHTML(selectors["bottomAds"], func(e *colly.HTMLElement) {
		addAdLinkToResult(adlinks.Bottom, adlinks.Link, e)
	})

	collector.OnScraped(finishScrapingResult)

	err = searchQueue.Run(collector)
	if err != nil {
		log.Println("Failed to run the queue", err.Error())
	}
}

func RequestHandler (request *colly.Request) {
	userAgent := RandomUserAgent()
	request.Headers.Set("User-Agent", userAgent)

	log.Println("Visiting", request.URL)
	keyword := keywordFromUrl(request.URL.String())
	request.Ctx.Put("resultID", fmt.Sprint(resultIDFromKeyword(keyword)))
}

func ResponseHandler (response *colly.Response) {
	log.Println("Visited", response.Request.URL)
}

func ErrorHandler(response *colly.Response, err error) {
	log.Println("Failed to request URL:", response.Request.URL, "with response:", response, "\nError:", err)
}

func ResultFromContext(context *colly.Context) (result *models.Result) {
	resultID := ResultIDFromContext(context)

	result, err := models.GetResultById(resultID)
	if err != nil {
		log.Fatal("Failed to get result by ID", err.Error())
	}

	return result
}

func ResultIDFromContext(context *colly.Context) (resultID int64) {
	rID := context.Get("resultID")
	resultID, err := num.ParseInt64(rID)
	if err != nil {
		log.Fatal("Failed to parse result ID", err.Error())
	}

	return resultID
}

func resultIDFromKeyword(keyword string) (resultID int64) {
	return scrapingResults[keyword]
}

func keywordFromUrl(urlStr string) (keyword string) {
	parsedUrl, err := url.Parse(urlStr)
	if err != nil {
		log.Println("Failed to parse url string", err.Error())
	}

	return parsedUrl.Query().Get("q")
}

func addNonAdLinkToResult(element *colly.HTMLElement) {
	link := element.Attr("href")
	result := ResultFromContext(element.Request.Ctx)

	if len(link) > 0 {
		link := &models.Link{
			Result: result,
			Link: link,
		}
		_, err := models.CreateLink(link)
		if err != nil {
			log.Fatal("Failed to creat link", err.Error())
		}
	}
}

func addAdLinkToResult(linkType string, linkPosition string, element *colly.HTMLElement) {
	link := element.Attr("href")
	result := ResultFromContext(element.Request.Ctx)

	if len(link) > 0 {
		adLink := &models.AdLink{
			Result: result,
			Type: linkType,
			Position: linkPosition,
			Link: link,
		}
		_, err := models.CreateAdLink(adLink)
		if err != nil {
			log.Fatal("Failed to creat adLink", err.Error())
		}
	}
}

func finishScrapingResult(response *colly.Response) {
	result := ResultFromContext(response.Ctx)
	result.Status = results.Completed
	err := models.UpdateResultById(result)
	if err != nil {
		log.Fatal("Failed to complete result", err.Error())
	}
	log.Println("Finished scraping for keyword:", result.Keyword)
}
