package serializers

import (
	"go-google-scraper-challenge/lib/models"
)

type AdLinkResponse struct {
	ID       int64  `jsonapi:"primary,ad_link"`
	ResultID int64  `jsonapi:"attr,result_id"`
	Type     string `jsonapi:"attr,type"`
	Position string `jsonapi:"attr,position"`
	Link     string `jsonapi:"attr,link"`
}

type AdLinkSerializer struct {
	AdLink *models.AdLink
}

func (s AdLinkSerializer) Response() (response *AdLinkResponse) {
	return &AdLinkResponse{
		ID:       s.AdLink.ID,
		ResultID: s.AdLink.ResultID,
		Link:     s.AdLink.Link,
		Position: s.AdLink.Position,
		Type:     s.AdLink.Type,
	}
}
