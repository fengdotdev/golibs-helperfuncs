package timefn_test

import (
	"testing"
	"time"

	"github.com/fengdotdev/golibs-helperfuncs/timefn"
	"github.com/fengdotdev/golibs-testing/assert"
)

func TestRemaining_1(t *testing.T) {
	tests := []struct {
		name     string
		time0    time.Time
		time1    time.Time
		expected time.Duration
	}{
		{
			name:     "one  year  difference",
			time0:    time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
			time1:    time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: time.Hour * 24 * 365,
		},
		{
			name:     "one leap year  difference",
			time0:    time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			time1:    time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: time.Hour * 24 * 366, // 2020 was leap year
		},
		{
			name:     "one day difference", 
			time0:    time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			time1:    time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC),
			expected: time.Hour * 24,
		},
		{
			name:     "one hour difference",
			time0:    time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			time1:    time.Date(2021, 1, 1, 1, 0, 0, 0, time.UTC), 
			expected: time.Hour,
		},
		{
			name:     "negative difference",
			time0:    time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC),
			time1:    time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: -24 * time.Hour,
		},

	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := timefn.Remaining(tt.time0, tt.time1)
			assert.Equal(t, r, tt.expected)
		})
	}

}
