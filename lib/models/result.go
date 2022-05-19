package models

import (
	"errors"
	database "go-google-scraper-challenge/bootstrap"
	"go-google-scraper-challenge/helpers"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Result struct {
	Base

	UserId  int64 `gorm:"not null;"`
	User    *User `gorm:"not null;foreignKey:UserId"`
	AdLinks []*AdLink
	Links   []*Link

	Keyword   string `gorm:"not null;"`
	Status    string `gorm:"not null;default:pending"`
	PageCache string
}

const (
	//Result statuses
	ResultStatusPending    = "pending"
	ResultStatusProcessing = "processing"
	ResultStatusCompleted  = "completed"
	ResultStatusFailed     = "failed"
)

var ResultStatuses = []string{ResultStatusPending, ResultStatusProcessing, ResultStatusCompleted, ResultStatusFailed}

// CreateResult insert a new Result into database and returns last inserted Id on success.
func CreateResult(result *Result) (int64, error) {
	queryResult := database.GetDB().Create(result)

	return result.Id, queryResult.Error
}

// GetResultById retrieves Result by Id. Returns error if Id doesn't exist
func GetResultById(id int64) (*Result, error) {
	result := &Result{}

	queryResult := database.GetDB().First(&result, id)
	if queryResult.Error != nil {
		return nil, queryResult.Error
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
	}

	result := &Result{}
	queryResult := database.GetDB().Where(query).Order("created_at").First(&result)

	return result, queryResult.Error
}

func ContainKeyword(keyword string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("keyword ilike ?", "%"+keyword+"%")
	}
}

func query(condition map[string]interface{}, orderBy string, offset int, limit int) (*gorm.DB, []*Result) {
	db := database.GetDB()

	if len(orderBy) > 0 {
		orderColumn := orderBy
		orderDesc := false

		if strings.HasPrefix(orderColumn, "-") {
			orderColumn = strings.SplitAfter(orderColumn, "-")[1]
			orderDesc = true
		}

		orderClause := clause.OrderByColumn{
			Column: clause.Column{Name: orderColumn},
			Desc:   orderDesc,
		}

		db = db.Order(orderClause)
	}

	limitClause := limit
	if limit < 0 {
		limitClause = helpers.GetPaginationPerPage()
	}
	db = db.Limit(limitClause)

	if offset > 0 {
		db = db.Offset(offset)
	}

	if len(condition) == 0 {
		condition = nil
	}

	if condition["keyword"] != nil {
		keyword := condition["keyword"].(string)
		delete(condition, "keyword")

		db = db.Scopes(ContainKeyword(keyword))
	}

	var results []*Result

	queryResult := db.Find(&results, condition)

	return queryResult, results
}

// GetResultsBy retrieves Results with given query. Returns empty list if no records exist
// possible query params are order, limit, offset and result property filter
func GetResultsBy(condition map[string]interface{}, orderBy string, offset int, limit int) ([]*Result, error) {
	queryResult, results := query(condition, orderBy, offset, limit)

	if queryResult.Error != nil {
		return nil, queryResult.Error
	}

	return results, nil
}

// CountResultsBy count all Results with given query. Returns 0 if no records exist
func CountResultsBy(condition map[string]interface{}, orderBy string, offset int, limit int) (int64, error) {
	count := int64(0)
	db, _ := query(condition, orderBy, offset, limit)
	countResult := db.Count(&count)

	if countResult.Error != nil {
		return count, countResult.Error
	}

	return count, nil
}

// UpdateResult updates Result and returns error if the record to be updated doesn't exist
func UpdateResult(result *Result) error {
	_, err := GetResultById(result.Id)
	if err != nil {
		return err
	}

	updateResult := database.GetDB().Save(result)

	return updateResult.Error
}

// UpdateResultStatus updates Result status returns any error from updating
func UpdateResultStatus(result *Result, status string) error {
	if !validResultStatus(status) {
		return errors.New("Invalid result status")
	}
	result.Status = status

	updateResult := database.GetDB().Save(result)

	return updateResult.Error
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
