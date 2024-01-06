package main

import "time"

type DateFormat string

const (
	YYYY_MM_DD DateFormat = "YYYY-MM-DD"
)

func daysSinceEpoch(date time.Time) int {
	return int(date.Unix() / 86400)
}
