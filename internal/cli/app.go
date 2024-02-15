package cli

import (
	"context"
	"fmt"
	"strings"
	"time"

	"sunsamago/app"
	"sunsamago/internal/configurations"
	"sunsamago/internal/local"
	"sunsamago/internal/output"
	"sunsamago/internal/pipe"
	"sunsamago/internal/sunsama"

	"github.com/urfave/cli/v2"
)

func NewApp() *cli.App {
	return &cli.App{
		Name:     "sunsamago",
		Version:  "v0.0.1",
		Authors:  authors(),
		Commands: commands(),
	}
}

func authors() []*cli.Author {
	return []*cli.Author{
		{
			Name:  "Samuel Tissot",
			Email: "tissotjobin@gmail.com",
		},
	}
}

func commands() cli.Commands {
	return []*cli.Command{
		timesheetCmd(),
	}

}

func timesheetCmd() *cli.Command {
	validOuput := []string{
		"pretty", "csv",
	}

	return &cli.Command{
		Name:        "timesheet",
		Usage:       "timesheet",
		Description: "creates a timesheet from sunsama tasks between two dates",
		Flags: []cli.Flag{
			&cli.TimestampFlag{
				Name:        "start",
				Aliases:     []string{"s"},
				DefaultText: "start date, included",
				Required:    true,
				Layout:      time.DateOnly,
				Timezone:    time.Local,
			},
			&cli.TimestampFlag{
				Name:        "end",
				Aliases:     []string{"e"},
				DefaultText: "end date, included",
				Required:    true,
				Layout:      time.DateOnly,
				Timezone:    time.Local,
			},
			&cli.StringSliceFlag{
				Name:        "categories",
				DefaultText: "filters to just these categories",
				Usage:       "-c category1 -c category2 ...",
				Aliases: []string{
					"c",
				},
				KeepSpace: false,
			},
			&cli.Float64Flag{
				Name:        "factor",
				DefaultText: "multiply the duration by a given factor",
				Usage:       "-fc 1.12",
				Aliases: []string{
					"fc",
				},
			},
			&cli.IntFlag{
				Name:        "round",
				DefaultText: "round to minutes",
				Usage:       "-r 15",
				Aliases: []string{
					"r",
				},
			},
			&cli.StringFlag{
				Name:     "durationFormatter",
				Category: "",
				DefaultText: "format the duration, hour (1.25), clock (1:15), " +
					"default (1h15m0s)",
				FilePath: "",
				Usage:    "-dfmt clock",
				Aliases:  []string{"dfmt"},
			},
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				DefaultText: fmt.Sprintf(
					"output format. Valid option are %s, default: pretty",
					strings.Join(validOuput, ", "),
				),
				Usage: "-o csv",
				Action: func(c *cli.Context, s string) error {
					for _, v := range validOuput {
						if v == s {
							return nil
						}
					}

					return fmt.Errorf(
						"invalid output. accepted are: %s",
						strings.Join(validOuput, ", "),
					)
				},
			},
			&cli.PathFlag{
				Name: "file",
				DefaultText: "the file and path for the output. " +
					"default timesheet-start-stop.csv",
				Usage:       "-file ./path/name.csv",
				Destination: nil,
				Aliases:     []string{"f"},
			},
		},
		Action: timesheetAction,
	}
}

func timesheetAction(cCtx *cli.Context) error {
	start := cCtx.Timestamp("start")
	end := cCtx.Timestamp("end")
	factor := cCtx.Float64("factor")
	categories := cCtx.StringSlice("categories")

	ctx, cancel := context.WithCancel(cCtx.Context)
	defer cancel()

	specs, err := configurations.InitializeSpecification()
	if err != nil {
		return fmt.Errorf("failed to find specification, missing .env file?, %w", err)
	}

	client, err := sunsama.NewClient(specs.SunsamaSessionID, local.LocalTimezone())
	if err != nil {
		return fmt.Errorf("failed to create sunsama client, %w", err)
	}

	eventsStream := sunsama.EventStream(ctx, client, *start, *end)

	// filter out only the requested categories
	if len(categories) > 0 {
		eventsStream = pipe.FilterByCategories(ctx, eventsStream, categories...)
	}

	// apply the factor
	if factor != 0 {
		eventsStream = pipe.ModifyDuration(
			ctx, eventsStream,
			func(duration time.Duration) time.Duration {
				return time.Second * time.Duration(duration.Seconds()*factor)
			},
		)
	}

	if roundMin := cCtx.Int("round"); roundMin != 0 {
		eventsStream = pipe.ModifyDuration(
			ctx, eventsStream,
			func(duration time.Duration) time.Duration {
				return duration.Round(time.Minute * time.Duration(roundMin))
			},
		)
	}

	writer := getWriter(cCtx)
	return writer.Write(ctx, eventsStream)
}

type Writer interface {
	Write(ctx context.Context, stream <-chan app.Event) error
}

func getWriter(cCtx *cli.Context) Writer {
	switch cCtx.String("output") {
	case "csv":
		filepath := cCtx.Path("file")
		if filepath == "" {
			filepath = fmt.Sprintf(
				"timesheet_%s-%s.csv",
				cCtx.Timestamp("start").Format(time.DateOnly),
				cCtx.Timestamp("end").Format(time.DateOnly),
			)
		}
		return output.NewCsvWriter(filepath, getDurationFormatter(cCtx))

	default:
		return output.NewPrettyWriter(getDurationFormatter(cCtx))
	}
}

type DurationFormatter func(duration time.Duration) string

func getDurationFormatter(cCtx *cli.Context) output.DurationFormatter {
	switch cCtx.String("durationFormatter") {
	case "clock":
		return output.ClockFormat
	case "hour":
		return output.HourFormat
	default:
		return output.DefaultFormat
	}
}
