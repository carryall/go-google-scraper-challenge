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
		chunks[0] = results[0:10]
		chunks[1] = results[10:]

		return chunks
	}

	chunks := make([][]*models.Result, 1)
	chunks[0] = results

	return chunks
}
