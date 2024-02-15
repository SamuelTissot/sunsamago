package sunsama

import (
	"time"

	"github.com/shurcooL/graphql"
)

type subtask struct {
	Title      graphql.String
	ActualTime []actualTime
}

func (s subtask) String() string {
	return string(s.Title)
}

func (s subtask) TimeLoggedOn(day time.Time) bool {
	for _, at := range s.ActualTime {
		if at.IsSameDay(day) {
			return true
		}
	}

	return false
}
