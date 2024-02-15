package sunsama

import (
	"fmt"
	"time"

	"github.com/shurcooL/graphql"
)

type dayDurationCalculator interface {
	Duration() (time.Duration, error)
	IsSameDay(to time.Time) bool
}

type task struct {
	Text          graphql.String
	Completed     graphql.Boolean
	ScheduledTime []scheduledTime
	TimeEstimate  graphql.Int
	StreamIDs     []graphql.String `graphql:"streamIds"`
	Subtasks      []subtask
	ActualTime    []actualTime
}

// DurationForADay returns the duration of a task for it's given day
//
//	logic:
//	if actualTime is NOT empty
//	    return delta between actualTime entries startDate and endDate
//	if completed == true && actualTime is empty && scheduledTime is NOT empty
//	    return delta between scheduledTime entries startDate and endDate
//	If completed == true && actualTime is empty && scheduledTime is empty
//	    return timeEstimate
//
//	:returns: the total duration of the task for the day
func (t task) DurationForADay(day time.Time) (time.Duration, error) {
	subTaskDuration, err := t.calculateSubtask(day)
	if err != nil {
		return 0, err
	}

	taskDuration, err := t.calculateTaskDuration(day)
	if err != nil {
		return 0, nil
	}

	return subTaskDuration + taskDuration, nil

}

func (t task) TitleForDay(day time.Time) string {
	title := string(t.Text)

	for _, s := range t.Subtasks {
		if s.TimeLoggedOn(day) {
			title = fmt.Sprintf("%s\n\t- %s", title, s.String())
		}
	}

	return title
}

func (t task) calculateTaskDuration(day time.Time) (time.Duration, error) {
	// if ActualTime slice is not empty
	if len(t.ActualTime) != 0 {
		return calculateTotalDurationForDay(toDayCalculator(t.ActualTime), day)
	}

	// if Completed is true and ScheduledTime slice is not empty
	if t.Completed && len(t.ScheduledTime) != 0 {
		return calculateTotalDurationForDay(toDayCalculator(t.ScheduledTime), day)
	}

	// if Completed is true, and ActualTime slice is empty and ScheduledTime slice is empty
	if t.Completed {
		return time.Duration(t.TimeEstimate) * time.Second, nil
	}

	return 0, nil
}

func (t task) calculateSubtask(day time.Time) (time.Duration, error) {
	if len(t.Subtasks) > 0 {
		var cals []dayDurationCalculator
		for _, subtask := range t.Subtasks {
			cals = append(cals, toDayCalculator(subtask.ActualTime)...)
		}
		return calculateTotalDurationForDay(cals, day)
	}

	return 0, nil
}

func toDayCalculator[T actualTime | scheduledTime](items []T) []dayDurationCalculator {
	out := make([]dayDurationCalculator, len(items))
	for i, item := range items {
		out[i] = dayDurationCalculator(item)
	}

	return out
}

func calculateTotalDurationForDay(
	calculators []dayDurationCalculator,
	day time.Time,
) (
	time.Duration,
	error,
) {
	var delta time.Duration

	for _, cal := range calculators {
		if cal.IsSameDay(day) {
			d, err := cal.Duration()
			if err != nil {
				return 0, err
			}

			delta += d
		}
	}

	return delta, nil
}
