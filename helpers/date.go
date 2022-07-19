package helpers

import (
	"time"
)

const (
	DateTimeLayout = "02/01/2006 15:04"
)

func FormatDateTime(dateTime time.Time) string {
	return dateTime.Local().Format(DateTimeLayout)
}
