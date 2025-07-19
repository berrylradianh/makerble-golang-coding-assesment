package helper

import (
	"time"
)

func ConvertStringToTime(dateStr string) (time.Time, error) {
	const layout = "2 January 2006"
	return time.Parse(layout, dateStr)
}
