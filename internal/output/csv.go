package output

import (
	"context"
	"encoding/csv"
	"os"
	"time"

	"sunsamago/app"
)

type csvWriter struct {
	filePath          string
	durationFormatter DurationFormatter
}

func NewCsvWriter(filePath string, durationformatter DurationFormatter) *csvWriter {
	return &csvWriter{
		filePath:          filePath,
		durationFormatter: durationformatter,
	}
}

func (c csvWriter) Write(ctx context.Context, stream <-chan app.Event) error {

	file, err := os.Create(c.filePath)
	if err != nil {
		return err
	}

	// get the data
	events := gather(stream)
	sortByDay(events)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(
		[]string{
			// channel, day, duration, name
			"Task", "Date", "Hours", "Notes",
		},
	)
	if err != nil {
		return err
	}

	for _, event := range events {
		d, err := event.Duration()
		if err != nil {
			return err
		}

		err = writer.Write(
			[]string{
				event.Channel(),
				event.Day().Format(time.DateOnly),
				c.durationFormatter(d),
				event.Name(),
			},
		)
		if err != nil {
			return err
		}
	}

	return nil
}
