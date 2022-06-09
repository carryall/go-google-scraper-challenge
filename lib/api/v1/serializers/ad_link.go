package serializers

type AdLinkResponse struct {
	ID       int64  `jsonapi:"primary,ad_link"`
	ResultID int64  `jsonapi:"attr,result_id"`
	Type     string `jsonapi:"attr,type"`
	Position string `jsonapi:"attr,position"`
	Link     string `jsonapi:"attr,link"`
}
