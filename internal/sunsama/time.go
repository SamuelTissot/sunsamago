package sunsama

import (
	"time"

	"github.com/shurcooL/graphql"
)

// timeFormat is the time format of the sunsama api
//
//	"2024-02-13T11:45:00.000Z"
const timeFormat = "2006-01-02T15:04:05.999Z"

func toTime(dateStr graphql.String) (time.Time, error) {
	return time.Parse(timeFormat, string(dateStr))
}

func isSameDay(day1, day2 time.Time) bool {
	y1, m1, d1 := day1.Date()
	y2, m2, d2 := day2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

func deltaBtwTimeStrs(startStr, endStr graphql.String) (time.Duration, error) {
	start, err := toTime(startStr)
	if err != nil {
		return 0, err
	}

	end, err := toTime(endStr)
	if err != nil {
		return 0, err
	}

	return end.Sub(start), nil
}
