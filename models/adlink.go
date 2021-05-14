package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type AdLink struct {
	Base

	Result          *Result  `orm:"rel(fk)"`
	Type            string   `orm:"type(text)"`
	Position        string   `orm:"type(text)"`
	Link            string   `orm:"type(text)"`
}

func init() {
	orm.RegisterModel(new(AdLink))
}

// TableName set the custom table name to plural because the default table name is singular
func (a *AdLink) TableName() string {
	return "ad_links"
}

// CreateAdLink insert a new AdLink into database and returns last inserted Id on success.
func CreateAdLink(adLink *AdLink) (int64, error) {
	ormer := orm.NewOrm()
	id, err := ormer.Insert(adLink)

	return id, err
}

// GetAdLinkById retrieves AdLink by Id. Returns error if Id doesn't exist
func GetAdLinkById(id int64) (*AdLink, error) {
	querySeter := adLinkQuerySeter().Filter("Id", id).RelatedSel()
	adLink := &AdLink{}
	err := querySeter.One(adLink)
	if err != nil {
		return nil, err
	}

	return adLink, nil
}

// GetAdLinksByResultId retrieves all AdLinks with Result Id. Returns empty list if no records exist
func GetAdLinksByResultId(resultId int64) ([]*AdLink, error) {
	querySeter := adLinkQuerySeter().Filter("result_id", resultId).RelatedSel()
	adLinks := []*AdLink{}
	_, err := querySeter.All(&adLinks)

	return adLinks, err
}

func adLinkQuerySeter() orm.QuerySeter {
	ormer := orm.NewOrm()

	return ormer.QueryTable(AdLink{})
}
