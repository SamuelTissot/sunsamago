package pipe

import "sunsamago/app"

func FilterByCategories(
	stopper app.Stopper,
	eventStream <-chan app.Event,
	categories ...string,
) <-chan app.Event {
	out := make(chan app.Event)

	go func() {
		defer close(out)
		for event := range eventStream {
			select {
			case <-stopper.Done():
				return
			default:
				if event.BelongsTo(categories...) {
					out <- event
				}
			}
		}
	}()

	return out
}
