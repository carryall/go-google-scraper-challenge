package models

import (
	database "go-google-scraper-challenge/bootstrap"
)

type Link struct {
	Base

	ResultId int64   `gorm:"not null;"`
	Result   *Result `gorm:"not null;foreignKey:ResultId"`
	Link     string  `gorm:"not null;"`
}

// TableName set the custom table name to plural because the default table name is singular
func (Link) TableName() string {
	return "links"
}

// AddLink insert a new Link into database and returns last inserted Id on success.
func CreateLink(link *Link) (int64, error) {
	result := database.GetDB().Create(link)

	return link.Id, result.Error
}

// GetLinkById retrieves Link by Id. Returns error if Id doesn't exist
func GetLinkById(id int64) (*Link, error) {
	link := &Link{}

	result := database.GetDB().First(&link, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return link, nil
}

// GetLinksByResultId retrieves all Links with Result Id. Returns empty list if no records exist
func GetLinksByResultId(resultId int64) ([]*Link, error) {
	links := []*Link{}

	result := database.GetDB().Model(&Link{}).Where("result_id = ?", resultId).Find(&links)

	return links, result.Error
}
