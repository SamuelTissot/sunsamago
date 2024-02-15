package sunsama

import (
	"time"

	"sunsamago/app"

	"github.com/shurcooL/graphql"
)

type event struct {
	task task
	day  time.Time
	err  error
	eventContexts
}

func (e event) Day() time.Time {
	return e.day
}

func (e event) Name() string {
	return e.task.TitleForDay(e.day)
}

func (e event) Duration() (time.Duration, error) {
	return e.task.DurationForADay(e.day)
}

func (e event) Err() error {
	return e.err
}

func EventStream(
	stopper app.Stopper,
	client *Client,
	startTime time.Time,
	endTime time.Time,
) <-chan app.Event {
	out := make(chan app.Event)

	go func() {
		defer close(out)

		sts, err := client.streamsByGroupID()
		if err != nil {
			out <- event{err: err}
			return
		}

		for d := startTime; d.After(endTime) == false; d = d.AddDate(0, 0, 1) {
			day := d

			select {
			case <-stopper.Done():
				return
			default:
				tasks, err := client.TaskByDay(day)
				if err != nil {
					return
				}

				for _, task := range tasks {
					out <- event{
						task:          task,
						eventContexts: sts.contextsForIds(task.StreamIDs),
						day:           day,
					}
				}
			}
		}
	}()

	return out
}

func getChannelName(channels []stream, channelIDs []graphql.String) []string {
	var channelNames []string
	for _, id := range channelIDs {
		for _, c := range channels {
			if c.ID == id {
				channelNames = append(channelNames, string(c.Description))
			}
		}
	}

	return channelNames
}
