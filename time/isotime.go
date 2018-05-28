package time

import (
	"time"
	"fmt"
)

// ISO8601Time represents a time that can be marshalled and unmarshalled to ISO 8601 convention
type ISO8601Time time.Time

// Now returns the current time, in the ISO 8601 format
func Now() ISO8601Time {
	return ISO8601Time(time.Now())
}

// MarshalJSON creates a JSON representation for a time in ISO 8601 convention
func (t ISO8601Time) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02T15:04:05.000-07:00"))
	return []byte(stamp), nil
}

// UnmarshalJSON parses a JSON representation in the ISO 8601 convention to a ISO8601Time
func (t *ISO8601Time) UnmarshalJSON(b []byte) error {
	time, err := time.Parse("\"2006-01-02T15:04:05.000-07:00\"", string(b))
	if err == nil {
		*t = ISO8601Time(time)
	}

	return err
}
