package main

import (
	"log"

	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CreateAdLinks_20210511_162955 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateAdLinks_20210511_162955{}
	m.Created = "20210511_162955"

	err := migration.Register("CreateAdLinks_20210511_162955", m)
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
}

// Run the migrations
func (m *CreateAdLinks_20210511_162955) Up() {
	m.SQL(`CREATE TABLE "ad_links"
		(
			id SERIAL,
			result_id integer REFERENCES "results" ON DELETE CASCADE,
			type text NOT NULL,
			position text NOT NULL,
			link text NOT NULL,
			created_at timestamp NOT NULL,
			updated_at timestamp NOT NULL,
			CONSTRAINT ad_link_pkey PRIMARY KEY (id)
		);`)
}

// Reverse the migrations
func (m *CreateAdLinks_20210511_162955) Down() {
	m.SQL(`DROP TABLE "ad_links";`)
}
