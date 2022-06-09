package serializers

type UserResponse struct {
	ID    int64  `jsonapi:"primary,user"`
	Email string `jsonapi:"attr,email"`
}
