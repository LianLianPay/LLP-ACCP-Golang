package utils

import (
	"time"
)

const defaultPattern = "20060102150405"

func GetTimestamp() string {
	return Format(time.Now(), defaultPattern)
}

func Format(date time.Time, pattern string) string {
	return date.Format(pattern)
}
