package serializers

type ResultsResponse struct {
	Results ResultResponse
}

type ResultResponse struct {
	ID      int64  `jsonapi:"primary,result"`
	Keyword string `jsonapi:"attr,keyword"`
}

type ResultJSONResponse struct {
	Data struct {
		ID         string `json:"id"`
		Attributes struct {
			Ketword int64 `json:"keyword"`
		} `json:"attributes"`
	} `json:"data"`
}
