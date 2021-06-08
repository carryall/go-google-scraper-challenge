package models

import (
	"errors"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type Result struct {
	Base

	User     *User      `orm:"rel(fk)"`
	AdLinks  []*AdLink  `orm:"reverse(many)"`
	Links    []*Link    `orm:"reverse(many)"`

	Keyword     string  `orm:"type(text)"`
	Status      string  `orm:"type(text);default(pending)"`
	PageCache   string  `orm:"type(text);null"`
}

const (
	//Result statuses
	ResultStatusPending = "pending"
	ResultStatusProcessing = "processing"
	ResultStatusCompleted = "completed"
	ResultStatusFailed = "failed"
)

var ResultStatuses = []string{ ResultStatusPending, ResultStatusProcessing, ResultStatusCompleted, ResultStatusFailed }

func init() {
	orm.RegisterModel(new(Result))
}

// TableName set the custom table name to plural because the default table name is singular
func (r *Result) TableName() string {
	return "results"
}

// CreateResult insert a new Result into database and returns last inserted Id on success.
func CreateResult(result *Result) (int64, error) {
	ormer := orm.NewOrm()
	id, err := ormer.Insert(result)

	return id, err
}

// GetResultById retrieves Result by Id. Returns error if Id doesn't exist
func GetResultById(id int64) (*Result, error) {
	query := map[string]interface{}{
		"id": id,
	}

	querySeter := resultQuerySeter(query).RelatedSel()
	result := &Result{}
	err := querySeter.One(result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetResultByIdWithRelations retrieves Result by Id with assigned relations. Returns error if Id doesn't exist
func GetResultByIdWithRelations(id int64) (*Result, error) {
	result, err := GetResultById(id)
	if err != nil {
		return nil, err
	}

	result.Links, err = GetLinksByResultId(result.Id)
	if err != nil {
		return result, err
	}

	result.AdLinks, err = GetAdLinksByResultId(result.Id)
	if err != nil {
		return result, err
	}

	return result, nil
}

// GetOldestPendingResult retrieves Result with pending status. Return err if no pending result
func GetOldestPendingResult() (*Result, error) {
	query := map[string]interface{}{
		"status": ResultStatusPending,
		"order": "created_at",
	}
	querySeter := resultQuerySeter(query).RelatedSel()
	result := &Result{}
	err := querySeter.One(result)
	if err != nil {
		return nil, err
	}

	return result, err
}

// GetResultsBy retrieves Results with given query. Returns empty list if no records exist
// possible query params are order, limit, offset and result property filter
func GetResultsBy(query map[string]interface{}) ([]*Result, error) {
	querySeter := resultQuerySeter(query).RelatedSel()
	var results []*Result
	_, err := querySeter.All(&results)

	return results, err
}

// CountResultsBy count all Results with given query. Returns 0 if no records exist
func CountResultsBy(query map[string]interface{}) (int64, error) {
	querySeter := resultQuerySeter(query)
	count, err := querySeter.Count()

	return count, err
}

// UpdateResultById updates Result by Id and returns error if the record to be updated doesn't exist
func UpdateResultById(result *Result) error {
	ormer := orm.NewOrm()
	_, err := GetResultById(result.Id)
	if err != nil {
		return err
	}

	num, err := ormer.Update(result)
	if err != nil {
		return err
	}

	logs.Info("Updated ", num, " results in database")
	return nil
}

// UpdateResultStatus updates Result status returns any error from updating
func UpdateResultStatus(result *Result, status string) error {
	if !validResultStatus(status) {
		return errors.New("Invalid result status")
	}
	result.Status = status

	return UpdateResultById(result)
}

func resultQuerySeter(query map[string]interface{}) orm.QuerySeter {
	ormer := orm.NewOrm()
	querySeter := ormer.QueryTable(Result{})

	for k, v := range query {
		switch k {
		case "limit":
			querySeter = querySeter.Limit(v)
		case "offset":
			querySeter = querySeter.Offset(v)
		case "order":
			querySeter = querySeter.OrderBy(v.(string))
		default:
			querySeter = querySeter.Filter(k, v)
		}
	}

	return querySeter
}

func validResultStatus(status string) bool {
	valid := false
	for _, resultStatus := range ResultStatuses {
		if status == resultStatus {
			valid = true
			break
		}
	}

	return valid
}
