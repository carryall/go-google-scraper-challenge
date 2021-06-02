package scraper

import (
	"fmt"
	"net/url"

	"go-google-scraper-challenge/models"

	"github.com/beego/beego/v2/core/logs"
	"github.com/gocolly/colly"
)

type Scraper struct {
	Result *models.Result
}

var selectors = map[string]string {
	"wholePage": "html",
	"topImageAds": "#tvcap .plantl a.pla-unit-title-link",
	"topLinkAds": "#tvcap .d5oMvf > a",
	"sideImageAds": ".cu-container a.plantl.clickable-card",
	"nonAds": "#search .yuRUbf > a",
	"bottomAds": "#bottomads .uEierd > a",
}

const GoogleSearchUrl = "http://www.google.com/search?q=%s"

func (s *Scraper) Run() error {
	escapedKeyword := url.QueryEscape(s.Result.Keyword)
	url := fmt.Sprintf(GoogleSearchUrl, escapedKeyword)

	return s.startScraping(url)
}

func (s *Scraper) startScraping(url string) error {
	collector := colly.NewCollector()

	collector.OnRequest(s.requestHandler)
	collector.OnResponse(s.responseHandler)
	collector.OnError(s.errorHandler)

	collector.OnHTML(selectors["wholePage"], s.wholePageCollector)

	collector.OnHTML(selectors["nonAds"], func(e *colly.HTMLElement) {
		s.addNonAdLinkToResult(e)
	})

	collector.OnHTML(selectors["topImageAds"], func(e *colly.HTMLElement) {
		s.addAdLinkToResult(models.AdLinkTypeImage, models.AdLinkPositionTop, e)
	})

	collector.OnHTML(selectors["topLinkAds"], func(e *colly.HTMLElement) {
		s.addAdLinkToResult(models.AdLinkTypeLink, models.AdLinkPositionTop, e)
	})

	collector.OnHTML(selectors["sideImageAds"], func(e *colly.HTMLElement) {
		s.addAdLinkToResult(models.AdLinkTypeImage, models.AdLinkPositionSide, e)
	})

	collector.OnHTML(selectors["bottomAds"], func(e *colly.HTMLElement) {
		s.addAdLinkToResult(models.AdLinkTypeLink, models.AdLinkPositionBottom, e)
	})

	collector.OnScraped(s.finishScrapingHandler)

	err := collector.Visit(url)
	if err != nil {
		return err
	}

	return nil
}

func (s *Scraper) requestHandler(request *colly.Request) {
	userAgent := RandomUserAgent()
	request.Headers.Set("User-Agent", userAgent)

	logs.Info("Visiting ", request.URL)

	err := s.Result.Process()
	if err != nil {
		logs.Error("Failed to process result:", err.Error())
	}
}

func (s *Scraper) responseHandler(response *colly.Response) {
	logs.Info("Visited ", response.Request.URL)
}

func (s *Scraper) errorHandler(response *colly.Response, errResponse error) {
	result := s.Result
	err := result.Fail()
	if err != nil {
		logs.Error("Failed to fail result:", err.Error())
	}

	logs.Error("Failed to scrap result ID:", result.Id, " URL:", response.Request.URL, " with response:", response, "\nError:", errResponse.Error())
}

func (s *Scraper) wholePageCollector(e *colly.HTMLElement) {
	result := s.Result
	result.PageCache = string(e.Response.Body)
	err := models.UpdateResultById(result)
	if err != nil {
		logs.Error("Failed to update result page cache:", err.Error())
	}
}

func (s *Scraper) addNonAdLinkToResult(element *colly.HTMLElement) {
	link := element.Attr("href")

	if len(link) > 0 {
		link := &models.Link{
			Result: s.Result,
			Link: link,
		}
		_, err := models.CreateLink(link)
		if err != nil {
			logs.Error("Failed to add link:", err.Error())
		}
	}
}

func (s *Scraper) addAdLinkToResult(linkType string, linkPosition string, element *colly.HTMLElement) {
	link := element.Attr("href")

	if len(link) > 0 {
		adLink := &models.AdLink{
			Result: s.Result,
			Type: linkType,
			Position: linkPosition,
			Link: link,
		}
		_, err := models.CreateAdLink(adLink)
		if err != nil {
			logs.Error("Failed to add adLink:", err.Error())
		}
	}
}

func (s *Scraper) finishScrapingHandler(response *colly.Response) {
	result := s.Result
	err := result.Complete()
	if err != nil {
		logs.Error("Failed to complete result:", err.Error())
	}
	logs.Info("Finished scraping for keyword:", result.Keyword)
}
