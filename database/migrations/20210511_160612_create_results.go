package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
	"github.com/beego/beego/v2/core/logs"
)

// DO NOT MODIFY
type CreateResults_20210511_160612 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateResults_20210511_160612{}
	m.Created = "20210511_160612"

	err := migration.Register("CreateResults_20210511_160612", m)
	if err != nil {
		logs.Error("Migration failed:", err)
	}
}

// Run the migrations
func (m *CreateResults_20210511_160612) Up() {
	m.SQL(`CREATE TABLE "results"
		(
			id SERIAL,
			user_id integer REFERENCES "users" ON DELETE CASCADE,
			keyword text NOT NULL,
			status text NOT NULL,
			page_cache text,
			created_at timestamp NOT NULL,
			updated_at timestamp NOT NULL,
			CONSTRAINT result_pkey PRIMARY KEY (id)
		);`)
}

// Reverse the migrations
func (m *CreateResults_20210511_160612) Down() {
	m.SQL(`DROP TABLE "results";`)
}
