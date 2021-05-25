package helpers

import (
	"go-google-scraper-challenge/models"
)

func PrepareResultSet(results []*models.Result) [][]*models.Result {
	if len(results) == 0 {
		return nil
	}

	if len(results) > 10 {
		chunks := make([][]*models.Result, 2)
		splitIndex := (len(results)+1)/2
		chunks[0] = results[0:splitIndex]
		chunks[1] = results[splitIndex:]

		return chunks
	}

	chunks := make([][]*models.Result, 1)
	chunks[0] = results

	return chunks
}
