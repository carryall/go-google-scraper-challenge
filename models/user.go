package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type User struct {
	Base

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

// UserWithEmailExist retrieves user email and returns true if user with email already exist.
func UserWithEmailExist(email string) (userExist bool) {
	ormer := orm.NewOrm()

	return ormer.QueryTable(User{}).Filter("Email", email).RelatedSel().Exist()
}
