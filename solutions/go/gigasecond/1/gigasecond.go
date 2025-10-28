// Handle all functions related to new time measure: "gigasecond"
package gigasecond

import "time"

// AddGigasecond add 1_000_000_000 seconds (called gigasecond) to a given date
func AddGigasecond(t time.Time) time.Time {
    gigasecond, _ := time.ParseDuration("1000000000s")
	return t.Add(gigasecond)
}
