package models

import (
	database "go-google-scraper-challenge/bootstrap"
)

type AdLink struct {
	Base

	ResultId int     `gorm:"not null;"`
	Result   *Result `gorm:"not null;foreignKey:ResultId"`
	Type     string  `gorm:"not null;"`
	Position string  `gorm:"not null;"`
	Link     string  `gorm:"not null;"`
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

// TableName set the custom table name to plural because the default table name is singular
func (al *AdLink) TableName() string {
	return "ad_links"
}

// CreateAdLink insert a new AdLink into database and returns last inserted Id on success.
func CreateAdLink(adLink *AdLink) (int64, error) {
	result := database.GetDB().Create(adLink)

	return adLink.Id, result.Error
}

// GetAdLinkById retrieves AdLink by Id. Returns error if Id doesn't exist
func GetAdLinkById(id int64) (*AdLink, error) {
	adLink := &AdLink{}

	result := database.GetDB().First(&adLink, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return adLink, nil
}

// GetAdLinksByResultId retrieves all AdLinks with Result Id. Returns empty list if no records exist
func GetAdLinksByResultId(resultId int64) ([]*AdLink, error) {
	adLinks := []*AdLink{}

	result := database.GetDB().Model(&AdLink{}).Where("result_id = ?", resultId).Find(&adLinks)

	return adLinks, result.Error
}
