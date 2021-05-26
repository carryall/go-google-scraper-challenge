package presenters

import (
	"go-google-scraper-challenge/helpers"
	"go-google-scraper-challenge/models"
)

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
