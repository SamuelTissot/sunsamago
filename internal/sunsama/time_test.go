package sunsama

import (
	"testing"
	"time"
)

func Test_time_format(t *testing.T) {
	tests := []struct {
		timeStr string
		want    time.Time
	}{
		{
			"2024-02-13T11:45:00.000Z",
			time.Date(2024, 02, 13, 11, 45, 0, 0, time.UTC),
		},
		{
			"2024-02-07T15:04:30.190Z",
			time.Date(2024, 02, 7, 15, 4, 30, 1.9e+8, time.UTC),
		},
		{
			"2024-02-07T15:04:32.444Z",
			time.Date(2024, 02, 07, 15, 4, 32, 4.44e+8, time.UTC),
		},
		{
			"2024-02-07T15:04:32.444444Z",
			time.Date(2024, 02, 07, 15, 4, 32, 444444000, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.timeStr, func(t *testing.T) {
				got, err := time.Parse(timeFormat, tt.timeStr)
				if err != nil {
					t.Error(err)
				}

				if !got.Equal(tt.want) {
					t.Errorf("conversion failed= %v, want %v", got, tt.want)
				}
			},
		)
	}
}
