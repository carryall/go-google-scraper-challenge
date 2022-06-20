package serializers

import (
	"go-google-scraper-challenge/lib/models"
)

type LinkResponse struct {
	ID       int64  `jsonapi:"primary,link"`
	ResultID int64  `jsonapi:"attr,result_id"`
	Link     string `jsonapi:"attr,link"`
}

type LinkSerializer struct {
	Link *models.Link
}

func (s LinkSerializer) Response() (response *LinkResponse) {
	return &LinkResponse{
		ID:       s.Link.ID,
		ResultID: s.Link.ResultID,
		Link:     s.Link.Link,
	}
}
