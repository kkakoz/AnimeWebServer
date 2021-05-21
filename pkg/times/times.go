package times

import (
	"time"
)

func GetNow() time.Time {
	return time.Now()
}

func ParseYMD(value string) (time.Time, error) {
	return time.Parse("2006-01-02", value)
}

func ParseAll(value string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", value)
}

func FormatYMD(t time.Time) string {
	return t.Format("2006-01-02")
}

func FormatAll(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}