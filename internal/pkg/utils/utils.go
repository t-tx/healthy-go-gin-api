package utils

import "time"

func GetTimeNowString() string {
	return time.Now().UTC().Format(time.RFC3339)
}
func GetTimeString(t time.Time) string {
	return t.UTC().Format(time.RFC3339)
}
func ParseTime(timeString string) (time.Time, error) {
	return time.Parse(time.RFC3339, timeString)
}
