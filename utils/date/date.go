package date

import "time"

const (
	apiDateLayout = "2019-01-01T15:04:05Z"
	apiDBLayout   = "2019-01-01 15:04:05"
)

// GetNow returns the actual date
func GetNow() time.Time {
	return time.Now().UTC()
}

// GetNowString returns the actual date as a string
func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

// GetNowDBFormat returns the actual date as a string
func GetNowDBFormat() string {
	return GetNow().Format(apiDBLayout)
}
