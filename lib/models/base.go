package models

import (
	"time"
)

type Base struct {
	Id        int64 `gorm:"unique;primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
