package presenters

import (
	"go-google-scraper-challenge/models"
)

type AdLinkCollectionPresenter struct {
	Position        string
	Links           []string
	LinkCount       int
}

func GetAdLinkCollection(adlink []*models.AdLink) map[string][]string {
	adLinkCollection := map[string][]string{
		models.AdLinkPositionTop: []string{},
		models.AdLinkPositionBottom: []string{},
		models.AdLinkPositionSide: []string{},
	}

	for _, al := range adlink {
		adLinkCollection[al.Position] = append(adLinkCollection[al.Position], al.Link)
	}

	return adLinkCollection
}
