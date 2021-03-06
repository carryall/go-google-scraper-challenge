package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
	"github.com/beego/beego/v2/core/logs"
)

// DO NOT MODIFY
type CreateUser_20210111_162749 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateUser_20210111_162749{}
	m.Created = "20210111_162749"

	err := migration.Register("CreateUser_20210111_162749", m)
	if err != nil {
		logs.Error("Migration failed:", err)
	}
}

// Run the migrations
func (m *CreateUser_20210111_162749) Up() {
	m.SQL(`CREATE EXTENSION citext;
		CREATE TABLE "user"
		(
			id SERIAL,
			email citext UNIQUE,
			hashed_password text,
			created_at timestamp NOT NULL,
			updated_at timestamp NOT NULL,
			CONSTRAINT user_pkey PRIMARY KEY (id)
		);`)
}

// Reverse the migrations
func (m *CreateUser_20210111_162749) Down() {
	m.SQL(`DROP TABLE "user";`)
}
