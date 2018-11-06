package gigasecond

import "time"

const gigasecond = 1e9

// AddGigasecond API function.  It uses a type from the Go standard library.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Duration(gigasecond) * time.Second)
}
