package serializers

type AuthenticationResponse struct {
	ID           string `jsonapi:"primary,authentication"`
	AccessToken  string `jsonapi:"attr,access_token"`
	RefreshToken string `jsonapi:"attr,refresh_token"`
}
