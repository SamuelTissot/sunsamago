package local

import (
	"testing"
	"time"
)

func TestLocalTimezone(t *testing.T) {
	systemLocal := time.Local

	tests := []struct {
		local string
		want  string
	}{
		{
			"EST",
			"EST",
		},
		{
			"Atlantic/Jan_Mayen",
			"CET",
		},
		{
			"America/Inuvik",
			"MST",
		},
		{
			"Europe/Paris",
			"CET",
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.local, func(t *testing.T) {
				loc, err := time.LoadLocation(tt.local)
				if err != nil {
					t.Error(err)
					return
				}

				time.Local = loc

				if got := LocalTimezone(); got != tt.want {
					t.Errorf("LocalTimezone() = %v, want %v", got, tt.want)
				}
			},
		)
	}

	time.Local = systemLocal
}
