package dateutils

import (
	"strings"
	"time"
)

func DateToTimestamp(date string) (time.Time, error) {
	const layout = "2006-01-02"
	date = strings.Trim(date, "\"")
	result, err := time.Parse(layout, date)
	if err != nil {
		return time.Time{}, err // Return error if parsing fails
	}
	return result, nil
}

func TimestampToDate(timestamp time.Time) string {
	return timestamp.Format("2006-01-02")
}
