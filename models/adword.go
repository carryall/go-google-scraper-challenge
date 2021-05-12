package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Adword struct {
	Base

	Result		    *Result  `orm:"rel(fk)"`
	Type            string   `orm:"type(text)"`
	Position        string   `orm:"type(text)"`
	Link            string   `orm:"type(text)"`
}

func init() {
	orm.RegisterModel(new(Adword))
}

// TableName set the custom table name to plural because the default table name is singular
func (a *Adword) TableName() string {
	return "adwords"
}

// CreateAdword insert a new Adword into database and returns last inserted Id on success.
func CreateAdword(adword *Adword) (id int64, err error) {
	ormer := orm.NewOrm()
	id, err = ormer.Insert(adword)
	return
}


// GetAdwordById retrieves Adword by Id. Returns error if Id doesn't exist
func GetAdwordById(id int64) (adword *Adword, err error) {
	querySeter := adwordQuerySeter().Filter("Id", id).RelatedSel()
	adword = &Adword{}
	err = querySeter.One(adword)
	if err != nil {
		return nil, err
	}

	return adword, nil
}

// GetAdwordsByResultId retrieves all Adwords with Result Id. Returns empty list if no records exist
func GetAdwordsByResultId(resultId int64) (adwords []*Adword, err error) {
	querySeter := adwordQuerySeter().Filter("result_id", resultId).RelatedSel()
	_, err = querySeter.All(&adwords)

	return adwords, err
}

func adwordQuerySeter() orm.QuerySeter {
	ormer := orm.NewOrm()
	return ormer.QueryTable(Adword{})
}
