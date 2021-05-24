package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
	"github.com/beego/beego/v2/core/logs"
)

// DO NOT MODIFY
type CreateLinks_20210513_185205 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateLinks_20210513_185205{}
	m.Created = "20210513_185205"

	err := migration.Register("CreateLinks_20210513_185205", m)
	if err != nil {
		logs.Error("Migration failed:", err)
	}
}

// Run the migrations
func (m *CreateLinks_20210513_185205) Up() {
	m.SQL(`CREATE TABLE "links"
		(
			id SERIAL,
			result_id integer REFERENCES "results" ON DELETE CASCADE,
			link text NOT NULL,
			created_at timestamp NOT NULL,
			updated_at timestamp NOT NULL,
			CONSTRAINT link_pkey PRIMARY KEY (id)
		);`)
}

// Reverse the migrations
func (m *CreateLinks_20210513_185205) Down() {
	m.SQL(`DROP TABLE "links";`)
}
