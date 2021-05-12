package models

import (
	"go-google-scraper-challenge/models/results"

	"github.com/beego/beego/v2/client/orm"
)

type Result struct {
	Base

	User		    *User     `orm:"rel(fk)"`
	AdWords			[]*Adword `orm:"reverse(many)"`

	Keyword       	string    `orm:"type(text)"`
	Status			string    `orm:"type(text);default(pending)"`
	NonAdLinks    	string    `orm:"type(json);null"`
	PageCache	  	string    `orm:"type(text);null"`
}

func init() {
	orm.RegisterModel(new(Result))
}

// TableName set the custom table name to plural because the default table name is singular
func (r *Result) TableName() string {
	return "results"
}

// CreateResult insert a new Result into database and returns last inserted Id on success.
func CreateResult(result *Result) (id int64, err error) {
	ormer := orm.NewOrm()
	result.Status = results.Pending
	id, err = ormer.Insert(result)
	return
}

// GetResultById retrieves Result by Id. Returns error if Id doesn't exist
func GetResultById(id int64) (result *Result, err error) {
	querySeter := resultQuerySeter().Filter("Id", id).RelatedSel()
	result = &Result{}
	err = querySeter.One(result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetResultsByUserId retrieves all Results with User Id. Returns empty list if no records exist
func GetResultsByUserId(userId int64) (results []*Result, err error) {
	querySeter := resultQuerySeter().Filter("user_id", userId).RelatedSel()
	_, err = querySeter.All(&results)

	return results, err
}

func resultQuerySeter() orm.QuerySeter {
	ormer := orm.NewOrm()
	return ormer.QueryTable(Result{})
}
