// Package gigasecond contains my solution for exercism's Go: gigasecond exercise
package gigasecond

import "time"

// AddGigasecond accepts a time.Time, adds 10^9 seconds, and returns the result
func AddGigasecond(t time.Time) time.Time {
	d, _ := time.ParseDuration("1000000000s")
	return t.Add(d)
}
