package serializers

type RegistrationResponse struct {
	ID           string `jsonapi:"primary,registration"`
	UserID       int64  `jsonapi:"attr,user_id"`
	AccessToken  string `jsonapi:"attr,access_token"`
	RefreshToken string `jsonapi:"attr,refresh_token"`
}
