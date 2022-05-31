package serializers

type RegistrationResponse struct {
	ID           string `jsonapi:"primary,registrations"`
	UserID       int64  `jsonapi:"attr,user_id"`
	AccessToken  string `jsonapi:"attr,access_token"`
	RefreshToken string `jsonapi:"attr,refresh_token"`
}

type RegistrationJSONResponse struct {
	Data struct {
		ID         string `json:"id"`
		Attributes struct {
			UserID       int64  `json:"user_id"`
			AccessToken  string `json:"access_token"`
			RefreshToken string `json:"refresh_token"`
		} `json:"attributes"`
	} `json:"data"`
}
