package main

import (
	"log"

	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type RenameUser_20210204_165028 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &RenameUser_20210204_165028{}
	m.Created = "20210204_165028"

	err := migration.Register("RenameUser_20210204_165028", m)
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
}

// Run the migrations
func (m *RenameUser_20210204_165028) Up() {
	m.SQL(`ALTER TABLE "user"
		RENAME TO "users";
	`)
}

// Reverse the migrations
func (m *RenameUser_20210204_165028) Down() {
	m.SQL(`ALTER TABLE "users"
		RENAME TO "user";
	`)
}
