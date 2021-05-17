package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Link struct {
	Base

	Result     *Result  `orm:"rel(fk)"`
	Link       string   `orm:"type(text)"`
}

func init() {
	orm.RegisterModel(new(Link))
}

// TableName set the custom table name to plural because the default table name is singular
func (l *Link) TableName() string {
	return "links"
}

// AddLink insert a new Link into database and returns last inserted Id on success.
func CreateLink(link *Link) (int64, error) {
	ormer := orm.NewOrm()
	id, err := ormer.Insert(link)

	return id, err
}

// GetLinkById retrieves Link by Id. Returns error if Id doesn't exist
func GetLinkById(id int64) (*Link, error) {
	querySeter := linkQuerySeter().Filter("Id", id).RelatedSel()
	link := &Link{}
	err := querySeter.One(link)
	if err != nil {
		return nil, err
	}

	return link, nil
}

// GetLinksByResultId retrieves all Links with Result Id. Returns empty list if no records exist
func GetLinksByResultId(resultId int64) ([]*Link, error) {
	querySeter := linkQuerySeter().Filter("result_id", resultId).RelatedSel()
	var links []*Link
	_, err := querySeter.All(&links)

	return links, err
}

func linkQuerySeter() orm.QuerySeter {
	ormer := orm.NewOrm()

	return ormer.QueryTable(Link{})
}
