// Handle all functions related to new time measure: "gigasecond"
package gigasecond

import "time"

const GIGASECOND_SECONDS = 1_000_000_000

// AddGigasecond add 1_000_000_000 seconds (called gigasecond) to a given date
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * GIGASECOND_SECONDS)
}
