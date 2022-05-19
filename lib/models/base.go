package models

import (
	"time"
)

type Base struct {
	ID        int64 `gorm:"unique;primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
