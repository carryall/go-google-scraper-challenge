package scraper

import (
	"fmt"
	"net/url"

	"go-google-scraper-challenge/helpers/log"
	"go-google-scraper-challenge/lib/models"

	"github.com/gocolly/colly"
)

type Scraper struct {
	Result *models.Result
}

var selectors = map[string]string{
	"wholePage":    "html",
	"topImageAds":  "#tvcap .plantl a.pla-unit-title-link",
	"topLinkAds":   "#tvcap .d5oMvf > a",
	"sideImageAds": ".cu-container a.plantl.clickable-card",
	"nonAds":       "#search .yuRUbf > a",
	"bottomAds":    "#bottomads .uEierd > a",
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

	log.Info("Visiting ", request.URL)

	err := models.UpdateResultStatus(s.Result, models.ResultStatusProcessing)
	if err != nil {
		log.Error("Failed to process result:", err.Error())
	}
}

func (s *Scraper) responseHandler(response *colly.Response) {
	log.Info("Visited ", response.Request.URL)
}

func (s *Scraper) errorHandler(response *colly.Response, errResponse error) {
	result := s.Result
	err := models.UpdateResultStatus(s.Result, models.ResultStatusFailed)
	if err != nil {
		log.Error("Failed to fail result:", err.Error())
	}

	log.Error("Failed to scrap result ID:", result.ID, " URL:", response.Request.URL, " with response:", response, "\nError:", errResponse.Error())
}

func (s *Scraper) wholePageCollector(e *colly.HTMLElement) {
	result := s.Result
	result.PageCache = string(e.Response.Body)
	err := models.UpdateResult(result)
	if err != nil {
		log.Error("Failed to update result page cache:", err.Error())
	}
}

func (s *Scraper) addNonAdLinkToResult(element *colly.HTMLElement) {
	link := element.Attr("href")

	if len(link) > 0 {
		link := &models.Link{
			Result: s.Result,
			Link:   link,
		}
		_, err := models.CreateLink(link)
		if err != nil {
			log.Error("Failed to add link:", err.Error())
		}
	}
}

func (s *Scraper) addAdLinkToResult(linkType string, linkPosition string, element *colly.HTMLElement) {
	link := element.Attr("href")

	if len(link) > 0 {
		adLink := &models.AdLink{
			Result:   s.Result,
			Type:     linkType,
			Position: linkPosition,
			Link:     link,
		}
		_, err := models.CreateAdLink(adLink)
		if err != nil {
			log.Error("Failed to add adLink:", err.Error())
		}
	}
}

func (s *Scraper) finishScrapingHandler(_ *colly.Response) {
	result := s.Result
	err := models.UpdateResultStatus(s.Result, models.ResultStatusCompleted)
	if err != nil {
		log.Error("Failed to complete result:", err.Error())
	}
	log.Info("Finished scraping for keyword:", result.Keyword)
}
