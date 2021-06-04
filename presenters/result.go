package presenters

import (
	"go-google-scraper-challenge/helpers"
	"go-google-scraper-challenge/models"
)

type ResultPresenter struct {
	AdLinks  map[string][]string
	Links    []*models.Link
	Id          int64
	Keyword     string
	Status      string
	PageCache   string
	TotalLinkCount int
	AdLinkCount int
	NonAdLinkCount int
}

func GetResultPresenter(result *models.Result) ResultPresenter {
	resultPresenter := ResultPresenter{
		Id: result.Id,
		Keyword: result.Keyword,
		Status: result.Status,
		AdLinks: GetAdLinkCollection(result.AdLinks),
		Links: result.Links,
		PageCache: result.PageCache,
		TotalLinkCount: len(result.AdLinks) + len(result.Links),
		AdLinkCount: len(result.AdLinks),
		NonAdLinkCount: len(result.Links),
	}

	return resultPresenter
}

func PrepareResultSet(results []*models.Result) [][]*models.Result {
	if len(results) == 0 {
		return nil
	}

	defaultPerPage := helpers.GetPaginationPerPage()
	splitIndex := defaultPerPage / 2
	if len(results) > splitIndex {
		chunks := make([][]*models.Result, 2)
		chunks[0] = results[0:splitIndex]
		chunks[1] = results[splitIndex:]

		return chunks
	}

	chunks := make([][]*models.Result, 1)
	chunks[0] = results

	return chunks
}
