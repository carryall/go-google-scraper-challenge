package main

import (
	"log"

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

	if err := migration.Register("CreateUser_20210111_162749", m); err != nil {
		log.Fatal("Migration failed:", err)
	}
}

// Run the migrations
func (m *CreateUser_20210111_162749) Up() {
	m.SQL(`CREATE TABLE "user"
		(
			id SERIAL,
			email text UNIQUE,
			encrypted_password text,
			created_at timestamp NOT NULL,
			updated_at timestamp NOT NULL,
			CONSTRAINT user_pkey PRIMARY KEY (id)
		);`)
}

// Reverse the migrations
func (m *CreateUser_20210111_162749) Down() {
	m.SQL(`DROP TABLE "user";`)
}
