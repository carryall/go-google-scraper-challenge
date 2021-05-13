package models

import (
	"log"

	"go-google-scraper-challenge/models/results"

	"github.com/beego/beego/v2/client/orm"
)

type Result struct {
	Base

	User    *User     `orm:"rel(fk)"`
	AdLinks []*AdLink `orm:"reverse(many)"`
	Links   []*Link   `orm:"reverse(many)"`

	Keyword       	string    `orm:"type(text)"`
	Status			string    `orm:"type(text);default(pending)"`
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

// UpdateResult updates Result by Id and returns error if the record to be updated doesn't exist
func UpdateResultById(result *Result) (err error) {
	ormer := orm.NewOrm()
	_, err = GetResultById(result.Id)
	if err != nil {
		return err
	}

	num, err := ormer.Update(result)
	if err != nil {
		return err
	}

	log.Println("Updated ", num, " results in database")
	return
}

func resultQuerySeter() orm.QuerySeter {
	ormer := orm.NewOrm()
	return ormer.QueryTable(Result{})
}
