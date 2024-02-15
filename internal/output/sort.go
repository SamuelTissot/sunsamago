package output

import (
	"sort"

	"sunsamago/app"
)

func sortByDay(events []app.Event) {
	sort.Slice(
		events, func(i, j int) bool {
			return events[i].Day().Before(events[j].Day())
		},
	)
}
