package serializers

type OAuthClientResponse struct {
	ID           string `jsonapi:"primary,oauth_clients"`
	ClientID     string `jsonapi:"attr,client_id"`
	ClientSecret string `jsonapi:"attr,client_secret"`
}

type OAuthClientJSONResponse struct {
	Data struct {
		ID         string `json:"id"`
		Attributes struct {
			ClientID     string `json:"client_id"`
			ClientSecret string `json:"client_secret"`
		} `json:"attributes"`
	} `json:"data"`
}
