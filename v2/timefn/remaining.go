package timefn

import "time"

func Remaining(oldesTime, newestTime time.Time) time.Duration {
	return newestTime.Sub(oldesTime)
}
