package serializers

type AuthenticationResponse struct {
	ID           int64  `jsonapi:"primary,authentications"`
	AccessToken  string `jsonapi:"attr,access_token"`
	RefreshToken string `jsonapi:"attr,refresh_token"`
	ExpiresIn    int64  `jsonapi:"attr,expires_in"`
	TokenType    string `jsonapi:"attr,token_type"`
}

type AuthenticationJSONResponse struct {
	Data struct {
		ID         string `json:"id"`
		Attributes struct {
			UserID       int64  `json:"user_id"`
			AccessToken  string `json:"access_token"`
			RefreshToken string `json:"refresh_token"`
			ExpiresIn    int64  `json:"expires_in"`
			TokenType    string `json:"token_type"`
		} `json:"attributes"`
	} `json:"data"`
}

type AuthenticationToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
	TokenType    string `json:"token_type"`
}
