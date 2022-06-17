package scraper

import (
	"fmt"
	"net/url"

	"go-google-scraper-challenge/database"
	"go-google-scraper-challenge/helpers/log"
	"go-google-scraper-challenge/lib/models"

	"github.com/gocolly/colly"
	"gorm.io/gorm"
)

type Scraper struct {
	Result     *models.Result
	NonAdLinks []string
	AdLinks    []AdLink
	PageCache  string
}

type AdLink struct {
	Type      string
	Postition string
	Link      string
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

	collector.OnHTML(selectors["wholePage"], func(e *colly.HTMLElement) {
		s.PageCache = string(e.Response.Body)
	})

	collector.OnHTML(selectors["nonAds"], func(e *colly.HTMLElement) {
		url := e.Attr("href")
		s.NonAdLinks = append(s.NonAdLinks, url)
	})

	collector.OnHTML(selectors["topImageAds"], func(e *colly.HTMLElement) {
		link := e.Attr("href")
		newAdLink := AdLink{models.AdLinkTypeImage, models.AdLinkPositionTop, link}
		s.AdLinks = append(s.AdLinks, newAdLink)
	})

	collector.OnHTML(selectors["topLinkAds"], func(e *colly.HTMLElement) {
		link := e.Attr("href")
		newAdLink := AdLink{models.AdLinkTypeLink, models.AdLinkPositionTop, link}
		s.AdLinks = append(s.AdLinks, newAdLink)
	})

	collector.OnHTML(selectors["sideImageAds"], func(e *colly.HTMLElement) {
		link := e.Attr("href")
		newAdLink := AdLink{models.AdLinkTypeImage, models.AdLinkPositionSide, link}
		s.AdLinks = append(s.AdLinks, newAdLink)
	})

	collector.OnHTML(selectors["bottomAds"], func(e *colly.HTMLElement) {
		link := e.Attr("href")
		newAdLink := AdLink{models.AdLinkTypeLink, models.AdLinkPositionBottom, link}
		s.AdLinks = append(s.AdLinks, newAdLink)
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

func (s *Scraper) savePageCache() error {
	result := s.Result
	result.PageCache = string(s.PageCache)
	err := models.UpdateResult(result)
	if err != nil {
		log.Error("Failed to update result page cache:", err.Error())

		return err
	}

	return nil
}

func (s *Scraper) addNonAdLinksToResult() error {
	for _, link := range s.NonAdLinks {
		if len(link) > 0 {
			link := &models.Link{
				Result: s.Result,
				Link:   link,
			}
			_, err := models.CreateLink(link)
			if err != nil {
				log.Error("Failed to add link:", err.Error())

				return err
			}
		}
	}

	return nil
}

func (s *Scraper) addAdLinksToResult() error {
	for _, adLink := range s.AdLinks {
		if len(adLink.Link) > 0 {
			adLink := &models.AdLink{
				Result:   s.Result,
				Type:     adLink.Type,
				Position: adLink.Postition,
				Link:     adLink.Link,
			}
			_, err := models.CreateAdLink(adLink)
			if err != nil {
				log.Error("Failed to add adLink:", err.Error())

				return err
			}
		}
	}

	return nil
}

func (s *Scraper) finishScrapingHandler(_ *colly.Response) {
	db := database.GetDB()
	db.Transaction(func(tx *gorm.DB) error {
		err := s.savePageCache()
		if err != nil {
			return err
		}

		err = s.addNonAdLinksToResult()
		if err != nil {
			return err
		}

		err = s.addAdLinksToResult()
		if err != nil {
			return err
		}

		result := s.Result
		err = models.UpdateResultStatus(result, models.ResultStatusCompleted)
		if err != nil {
			log.Error("Failed to complete result:", err.Error())
			return err
		}
		log.Info("Finished scraping for keyword:", result.Keyword)

		return nil
	})
}
