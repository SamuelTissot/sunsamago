package pipe

import (
	"time"

	"sunsamago/app"
)

type eventMultiplier struct {
	app.Event
	fn func(duration time.Duration) time.Duration
}

func (e eventMultiplier) Duration() (time.Duration, error) {
	d, err := e.Event.Duration()
	if err != nil {
		return 0, err
	}

	return e.fn(d), nil
}

func ModifyDuration(
	stopper app.Stopper,
	eventStream <-chan app.Event,
	fn func(duration time.Duration) time.Duration,
) <-chan app.Event {
	out := make(chan app.Event)

	go func() {
		defer close(out)
		for event := range eventStream {
			select {
			case <-stopper.Done():
				return
			default:
				out <- eventMultiplier{
					Event: event,
					fn:    fn,
				}
			}
		}
	}()

	return out
}
