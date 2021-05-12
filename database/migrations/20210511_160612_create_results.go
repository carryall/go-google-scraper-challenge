package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
	"log"
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
		log.Fatal("Migration failed:", err)
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
			non_ad_links json,
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
