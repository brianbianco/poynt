package main

import (
	"time"
)

const TimeFormat = "20060102150405.000"
const TimeFormatMaxLength = len(TimeFormat)

func StringToTime(s string) (time.Time, error) {
	t, err := time.Parse(TimeFormat, s)
	return t, err
}

func TimeToString(t time.Time) string {
	s := t.Format(TimeFormat)
	return s
}
