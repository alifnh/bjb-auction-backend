package dateutils

import "time"

func DateToTimestamp(date string) (time.Time, error) {
	const layout = "2006-01-02"
	result, err := time.Parse(layout, date)
	if err != nil {
		return time.Time{}, err // Return error if parsing fails
	}
	return result, nil
}
