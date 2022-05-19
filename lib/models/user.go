package models

import (
	database "go-google-scraper-challenge/bootstrap"
)

type User struct {
	Base

	Results []*Result

	Email          string `gorm:"not null;unique;size:128"`
	HashedPassword string `gorm:"not null"`
}

// CreateUser insert a new User into database and returns last inserted Id on success.
func CreateUser(user *User) (int64, error) {
	result := database.GetDB().Create(user)

	return user.Id, result.Error
}

// GetUserById get a user with given id, return error if user with id does not exist
func GetUserById(id int64) (*User, error) {
	user := &User{}

	result := database.GetDB().First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

// UserEmailAlreadyExist retrieves user email and returns true if user with email already exist.
func UserEmailAlreadyExist(email string) bool {
	_, err := GetUserByEmail(email)

	return err == nil
}

// GetUserByEmail retrieves User by Email and returns error if User with given Email doesn't exist.
func GetUserByEmail(email string) (*User, error) {
	user := &User{}

	result := database.GetDB().Where("email = ?", email).First(&user)

	return user, result.Error
}
