package serializers

type ResultsResponse struct {
	Results ResultResponse
}

type ResultResponse struct {
	ID      int64  `jsonapi:"primary,result"`
	Keyword string `jsonapi:"attr,keyword"`
	UserID  int64  `jsonapi:"attr,user_id"`
}

type ResultsJSONResponse struct {
	Data []struct {
		ID         string `json:"id"`
		Attributes struct {
			Keyword string `json:"keyword"`
			UserID  int64  `json:"user_id"`
		} `json:"attributes"`
	} `json:"data"`
}
