package models

import (
	"go-google-scraper-challenge/database"
)

type User struct {
	Base

	Results []*Result

	Email          string `gorm:"not null;unique;size:128"`
	HashedPassword string `gorm:"not null"`
}

// CreateUser insert a new User into database and returns last inserted ID on success.
func CreateUser(user *User) (int64, error) {
	result := database.GetDB().Create(user)

	return user.ID, result.Error
}

// GetUserByID get a user with given id, return error if user with ID does not exist
func GetUserByID(id int64) (*User, error) {
	user := &User{}

	result := database.GetDB().First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

// UserEmailAlreadyExisted retrieves user email and returns true if user with email already exist.
func UserEmailAlreadyExisted(email string) bool {
	_, err := GetUserByEmail(email)

	return err == nil
}

// GetUserByEmail retrieves User by Email and returns error if User with given Email doesn't exist.
func GetUserByEmail(email string) (*User, error) {
	user := &User{}

	result := database.GetDB().Where("LOWER(email) = LOWER(?)", email).First(&user)

	return user, result.Error
}

func DeleteUser(userID int64) error {
	result := database.GetDB().Delete(&User{}, userID)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
