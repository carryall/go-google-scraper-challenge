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

// AddUser insert a new User into database and returns
// last inserted Id on success.
func AddUser(m *User) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}
