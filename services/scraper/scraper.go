package scraper

import (
	"fmt"
	"log"
	"net/url"

	"go-google-scraper-challenge/models"

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
	queue := setupQueue(user, keywords)

	startScraping(queue)
}

func setupQueue(user *models.User, keywords []string) *queue.Queue {
	searchQueue, err := queue.New(2, &queue.InMemoryQueueStorage{MaxSize: 10000})
	if err != nil {
		log.Fatal("Failed to create a queue", err.Error())
	}

	for _, k := range keywords {
		scrapingResults[k] = createResult(user, k)

		escapedKeyword := url.QueryEscape(k)
		err = searchQueue.AddURL(fmt.Sprintf(GOOGLE_SEARCH_URL, escapedKeyword))
		if err != nil {
			log.Println("Failed to add url to queue", err.Error())
		}
	}

	return searchQueue
}

func createResult(user *models.User, keyword string) int64  {
	result := &models.Result{
		User: user,
		Keyword: keyword,
	}
	resultID, err := models.CreateResult(result)
	if err != nil {
		log.Fatal("Failed to create result", err.Error())
	}

	return resultID
}

func startScraping(queue *queue.Queue)  {
	collector := colly.NewCollector(colly.Async(true))

	collector.OnRequest(requestHandler)
	collector.OnResponse(responseHandler)
	collector.OnError(errorHandler)

	collector.OnHTML(selectors["wholePage"], wholePageCollector)

	collector.OnHTML(selectors["nonAds"], func(e *colly.HTMLElement) {
		addNonAdLinkToResult(e)
	})

	collector.OnHTML(selectors["topImageAds"], func(e *colly.HTMLElement) {
		addAdLinkToResult(models.AdLinkTypeImage, models.AdLinkPositionTop, e)
	})

	collector.OnHTML(selectors["topLinkAds"], func(e *colly.HTMLElement) {
		addAdLinkToResult(models.AdLinkTypeLink, models.AdLinkPositionTop, e)
	})

	collector.OnHTML(selectors["sideImageAds"], func(e *colly.HTMLElement) {
		addAdLinkToResult(models.AdLinkTypeImage, models.AdLinkPositionSide, e)
	})

	collector.OnHTML(selectors["bottomAds"], func(e *colly.HTMLElement) {
		addAdLinkToResult(models.AdLinkTypeLink, models.AdLinkPositionBottom, e)
	})

	collector.OnScraped(finishScrapingHandler)

	err := queue.Run(collector)
	if err != nil {
		log.Println("Failed to run the queue", err.Error())
	}
}

func requestHandler(request *colly.Request) {
	userAgent := RandomUserAgent()
	request.Headers.Set("User-Agent", userAgent)

	log.Println("Visiting", request.URL)
	keyword := keywordFromUrl(request.URL.String())
	request.Ctx.Put("resultID", fmt.Sprint(resultIDFromKeyword(keyword)))
}

func responseHandler(response *colly.Response) {
	log.Println("Visited", response.Request.URL)
}

func errorHandler(response *colly.Response, err error) {
	log.Println("Failed to request URL:", response.Request.URL, "with response:", response, "\nError:", err)
}

func wholePageCollector(e *colly.HTMLElement) {
	result := getResultFromContext(e.Request.Ctx)
	result.PageCache = string(e.Response.Body)
	err := models.UpdateResultById(result)
	if err != nil {
		log.Fatal("Failed to update result page cache", err.Error())
	}
}

func addNonAdLinkToResult(element *colly.HTMLElement) {
	link := element.Attr("href")
	result := getResultFromContext(element.Request.Ctx)

	if len(link) > 0 {
		link := &models.Link{
			Result: result,
			Link: link,
		}
		_, err := models.CreateLink(link)
		if err != nil {
			log.Fatal("Failed to creat link:", err.Error())
		}
	}
}

func addAdLinkToResult(linkType string, linkPosition string, element *colly.HTMLElement) {
	link := element.Attr("href")
	result := getResultFromContext(element.Request.Ctx)

	if len(link) > 0 {
		adLink := &models.AdLink{
			Result: result,
			Type: linkType,
			Position: linkPosition,
			Link: link,
		}
		_, err := models.CreateAdLink(adLink)
		if err != nil {
			log.Fatal("Failed to creat adLink:", err.Error())
		}
	}
}

func finishScrapingHandler(response *colly.Response) {
	result := getResultFromContext(response.Ctx)
	result.Status = models.ResultStatusCompleted
	err := models.UpdateResultById(result)
	if err != nil {
		log.Fatal("Failed to complete result", err.Error())
	}
	log.Println("Finished scraping for keyword:", result.Keyword)
}

func getResultFromContext(context *colly.Context) *models.Result {
	resultID := getResultIDFromContext(context)

	result, err := models.GetResultById(resultID)
	if err != nil {
		log.Fatal("Failed to get result by ID", err.Error())
	}

	return result
}

func getResultIDFromContext(context *colly.Context) int64 {
	rID := context.Get("resultID")
	resultID, err := num.ParseInt64(rID)
	if err != nil {
		log.Fatal("Failed to parse result ID", err.Error())
	}

	return resultID
}

func resultIDFromKeyword(keyword string) int64 {
	return scrapingResults[keyword]
}

func keywordFromUrl(urlStr string) string {
	parsedUrl, err := url.Parse(urlStr)
	if err != nil {
		log.Println("Failed to parse url string", err.Error())
	}

	return parsedUrl.Query().Get("q")
}
