package models

import (
	database "go-google-scraper-challenge/bootstrap"
)

type AdLink struct {
	Base

	ResultID int64  `gorm:"not null;"`
	Type     string `gorm:"not null;"`
	Position string `gorm:"not null;"`
	Link     string `gorm:"not null;"`

	Result *Result `gorm:"not null;"`
}

const (
	// AdLink positions
	AdLinkPositionTop    = "top"
	AdLinkPositionBottom = "bottom"
	AdLinkPositionSide   = "side"
	// AdLink types
	AdLinkTypeImage = "image"
	AdLinkTypeLink  = "link"
)

// CreateAdLink insert a new AdLink into database and returns last inserted ID on success.
func CreateAdLink(adLink *AdLink) (int64, error) {
	result := database.GetDB().Create(adLink)

	return adLink.ID, result.Error
}

// GetAdLinkByID retrieves AdLink by ID. Returns error if ID doesn't exist
func GetAdLinkByID(id int64) (*AdLink, error) {
	adLink := &AdLink{}

	result := database.GetDB().First(&adLink, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return adLink, nil
}

// GetAdLinksByResultID retrieves all AdLinks with Result ID. Returns empty list if no records exist
func GetAdLinksByResultID(resultID int64) ([]*AdLink, error) {
	adLinks := []*AdLink{}

	result := database.GetDB().Model(&AdLink{}).Where("result_id = ?", resultID).Find(&adLinks)

	return adLinks, result.Error
}
