package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type User struct {
	Default

	Email             string `orm:"unique;size(128)"`
	EncryptedPassword string
}

var ormer orm.Ormer

func init() {
	orm.RegisterModel(new(User))
	ormer = orm.NewOrm()
}

// AddUser insert a new User into database and returns last inserted Id on success.
func AddUser(m *User) (id int64, err error) {
	id, err = ormer.Insert(m)
	return
}

// GetUserByEmail retrieves User by Email and returns error if Email doesn't exist.
func GetUserByEmail(email string) (user *User, err error) {
	user = &User{Email: email}

	err = ormer.QueryTable(User{}).Filter("Email", email).RelatedSel().One(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
