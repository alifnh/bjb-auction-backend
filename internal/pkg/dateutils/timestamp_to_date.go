package dateutils

import "time"

func TimestampToDate(timestamp time.Time) (string, error) {
	return timestamp.Format("2006-01-02"), nil
}

func TimestampToDateTime(timestamp time.Time) (string, error) {
	return timestamp.Format("2006-01-02 15:04:05"), nil
}
