package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type User struct {
	Base

	Email          string `orm:"unique;size(128)"`
	HashedPassword string
}

func init() {
	orm.RegisterModel(new(User))
}

// CreateUser insert a new User into database and returns last inserted Id on success.
func CreateUser(m *User) (id int64, err error) {
	ormer := orm.NewOrm()
	return ormer.Insert(m)
}

// GetUserById get a user with given id, return error if user with id does not exist
func GetUserById(id int64) (user *User, err error) {
	ormer := orm.NewOrm()
	user = &User{}

	err = ormer.QueryTable(User{}).Filter("Id", id).RelatedSel().One(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// UserEmailAlreadyExist retrieves user email and returns true if user with email already exist.
func UserEmailAlreadyExist(email string) (userExist bool) {
	ormer := orm.NewOrm()

	return ormer.QueryTable(User{}).Filter("Email", email).RelatedSel().Exist()
}

// FindUserByEmail retrieves User by Email and returns error if User with given Email doesn't exist.
func FindUserByEmail(email string) (user *User, err error) {
	ormer := orm.NewOrm()
	user = &User{Email: email}

	err = ormer.QueryTable(User{}).Filter("Email", email).RelatedSel().One(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
