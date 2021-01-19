package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type User struct {
	Default

	Email             string `orm:"unique;size(128)"`
	EncryptedPassword string
}

func init() {
	orm.RegisterModel(new(User))

}

// AddUser insert a new User into database and returns last inserted Id on success.
func AddUser(m *User) (id int64, err error) {
	ormer := orm.NewOrm()
	id, err = ormer.Insert(m)
	return
}

// GetUserByEmail retrieves User by Email and returns error if Email doesn't exist.
func GetUserByEmail(email string) (user *User, err error) {
	ormer := orm.NewOrm()
	user = &User{Email: email}

	err = ormer.QueryTable(User{}).Filter("Email", email).RelatedSel().One(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
