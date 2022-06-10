package serializers

import "go-google-scraper-challenge/lib/models"

type UserResponse struct {
	ID    int64  `jsonapi:"primary,user"`
	Email string `jsonapi:"attr,email"`
}

type UserSerializer struct {
	User *models.User
}

func (s UserSerializer) Response() (response *UserResponse) {
	return &UserResponse{
		ID:    s.User.ID,
		Email: s.User.Email,
	}
}
