
// Package gigasecond should have a package comment that summarizes what it's about.
//add a billion seconds
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond should have a comment documenting it.
func AddGigasecond(t time.Time) time.Time {
	var d time.Duration = time.Second * 1000_000_000
	n := t.Add(d)
	return n
}
