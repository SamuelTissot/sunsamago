package sunsama

import (
	"time"

	"github.com/shurcooL/graphql"
)

type scheduledTime struct {
	StartDate graphql.String
	EndDate   graphql.String
	IsAllDay  graphql.Boolean
}

func (s scheduledTime) Duration() (time.Duration, error) {
	if s.IsAllDay {
		return time.Hour * 24, nil
	}

	return deltaBtwTimeStrs(s.StartDate, s.EndDate)
}

func (s scheduledTime) IsSameDay(to time.Time) bool {
	day, err := toTime(s.StartDate)
	if err != nil {
		return false
	}

	return isSameDay(day, to)
}
