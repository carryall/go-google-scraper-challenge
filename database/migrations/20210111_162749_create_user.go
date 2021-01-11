package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CreateUser_20210111_162749 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateUser_20210111_162749{}
	m.Created = "20210111_162749"

	migration.Register("CreateUser_20210111_162749", m)
}

// Run the migrations
func (m *CreateUser_20210111_162749) Up() {
	m.SQL(`CREATE TABLE "user"
		(
			id SERIAL,
			email text UNIQUE,
			created_at timestamp NOT NULL,
			updated_at timestamp NOT NULL,
			CONSTRAINT user_pkey PRIMARY KEY (id)
		);`)
}

// Reverse the migrations
func (m *CreateUser_20210111_162749) Down() {
	m.SQL(`DROP TABLE "user";`)
}
