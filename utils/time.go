package utils

import "time"

func Now() time.Time {
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		loc = time.UTC
	}
	return time.Now().In(loc)
}
