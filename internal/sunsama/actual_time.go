package sunsama

import (
	"time"

	"github.com/shurcooL/graphql"
)

type actualTime struct {
	StartDate graphql.String
	EndDate   graphql.String
	// DurationInSeconds is a field that I was told would be deprecated in the future
	// by sunsama
	DurationInSeconds graphql.Int `graphql:"duration"`
}

func (a actualTime) Duration() (time.Duration, error) {
	if a.StartDate != "" && a.EndDate != "" {
		return deltaBtwTimeStrs(a.StartDate, a.EndDate)
	}

	return time.Second * time.Duration(a.DurationInSeconds), nil
}

func (a actualTime) IsSameDay(to time.Time) bool {
	day, err := toTime(a.StartDate)
	if err != nil {
		return false
	}

	return isSameDay(day, to)
}
