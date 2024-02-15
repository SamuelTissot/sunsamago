package output

import "sunsamago/app"

func gather(stream <-chan app.Event) []app.Event {
	var events []app.Event

	for event := range stream {
		events = append(events, event)
	}

	return events
}
