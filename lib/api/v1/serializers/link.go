package serializers

type LinkResponse struct {
	ID       int64  `jsonapi:"primary,link"`
	ResultID int64  `jsonapi:"attr,result_id"`
	Link     string `jsonapi:"attr,link"`
}
