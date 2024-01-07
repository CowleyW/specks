package main

import "time"

type DateFormat string

const (
	YYYY_MM_DD DateFormat = "YYYY-MM-DD"
	HH_MM_SS              = "hh:mm:ss"
)

func daysSinceEpoch(date time.Time) int {
	return int(date.Unix() / 86400)
}
