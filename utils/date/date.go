package date

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
	apiDBLayout   = "2006-01-02 15:04:05"
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
