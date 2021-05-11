package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CreateAdwords_20210511_162955 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateAdwords_20210511_162955{}
	m.Created = "20210511_162955"

	migration.Register("CreateAdwords_20210511_162955", m)
}

// Run the migrations
func (m *CreateAdwords_20210511_162955) Up() {
	m.SQL(`CREATE TYPE adword_position AS ENUM ('top', 'bottom', 'side');
		CREATE TYPE adword_type AS ENUM ('image', 'link');
		CREATE TABLE "ad_words"
		(
			id SERIAL,
			result_id integer REFERENCES "results" ON DELETE CASCADE,
			type adword_type NOT NULL,
			position adword_position NOT NULL,
			created_at timestamp NOT NULL,
			updated_at timestamp NOT NULL,
			CONSTRAINT adword_pkey PRIMARY KEY (id)
		);`)
}

// Reverse the migrations
func (m *CreateAdwords_20210511_162955) Down() {
	m.SQL(`DROP TABLE "ad_words";`)
}
