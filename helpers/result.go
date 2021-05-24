package helpers

import (
	"go-google-scraper-challenge/models"
)

func DivideByHalf(results []*models.Result) [][]*models.Result {
	if len(results) == 0 {
		return nil
	}
	divided := make([][]*models.Result, 2)
	splitIndex := (len(results)+1)/2
	divided[0] = results[0:splitIndex]
	divided[1] = results[splitIndex:len(results)]

	return divided
}
