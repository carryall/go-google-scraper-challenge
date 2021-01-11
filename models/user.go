package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type User struct {
	Default

	Id    int64  `orm:"auto"`
	Email string `orm:"unique;size(128)"`
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
