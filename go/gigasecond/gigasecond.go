// Package gigasecond
package gigasecond

import "time"

// Constant declaration.
const testVersion = 4

// AddGigasecond adds 10^9 seconds to t, returns the new time.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Duration(1000000000 * time.Second))
}
