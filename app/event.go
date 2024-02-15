package app

import "time"

type Event interface {
	Day() time.Time
	Name() string
	Duration() (time.Duration, error)
	Context() string
	Channel() string
	BelongsTo(names ...string) bool
	IsPrivate() bool
	Err() error
}
