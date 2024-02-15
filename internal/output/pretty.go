package output

import (
	"context"
	"os"
	"strconv"
	"time"

	"sunsamago/app"

	"github.com/jedib0t/go-pretty/v6/table"
)

type prettyWriter struct {
	durationFormatter DurationFormatter
}

func NewPrettyWriter(durationFormatter DurationFormatter) *prettyWriter {
	return &prettyWriter{durationFormatter: durationFormatter}
}

func (c prettyWriter) Write(ctx context.Context, stream <-chan app.Event) error {
	events := gather(stream)
	sortByDay(events)

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(
		table.Row{
			"#", "Day", "Name", "Duration", "Channel", "Private", "Error",
		},
	)

	totalDuration := time.Duration(0)
	for i, event := range events {
		d, err := event.Duration()
		if err != nil {
			return err
		}

		totalDuration += d

		t.AppendRow(
			table.Row{
				strconv.Itoa(i),
				event.Day().Format(time.DateOnly),
				event.Name(),
				c.durationFormatter(d),
				event.Channel(),
				event.IsPrivate(),
			},
		)
	}

	t.AppendSeparator()
	t.AppendFooter(table.Row{"", "", "Total", totalDuration.String()})
	t.Render()

	return nil
}
