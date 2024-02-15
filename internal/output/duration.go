package output

import (
	"fmt"
	"time"
)

type DurationFormatter func(duration time.Duration) string

func ClockFormat(d time.Duration) string {
	d = d.Round(time.Minute)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	return fmt.Sprintf("%d:%02d", h, m)
}

func HourFormat(d time.Duration) string {
	return fmt.Sprintf("%.02f", d.Hours())
}

func DefaultFormat(d time.Duration) string {
	return d.String()
}
