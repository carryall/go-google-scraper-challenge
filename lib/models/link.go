package models

import (
	"go-google-scraper-challenge/database"
)

type Link struct {
	Base

	ResultID int64  `gorm:"not null;"`
	Link     string `gorm:"not null;"`

	Result *Result `gorm:"not null;"`
}

// AddLink insert a new Link into database and returns last inserted ID on success.
func CreateLink(link *Link) (int64, error) {
	result := database.GetDB().Create(link)

	return link.ID, result.Error
}

// GetLinkByID retrieves Link by ID. Returns error if ID doesn't exist
func GetLinkByID(id int64) (*Link, error) {
	link := &Link{}

	result := database.GetDB().First(&link, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return link, nil
}

// GetLinksByResultID retrieves all Links with Result ID. Returns empty list if no records exist
func GetLinksByResultID(resultID int64) ([]*Link, error) {
	links := []*Link{}

	result := database.GetDB().Model(&Link{}).Where("result_id = ?", resultID).Find(&links)

	return links, result.Error
}
