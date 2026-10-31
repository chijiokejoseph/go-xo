package logic

import "time"

// Sleep stops time execution for a number of milliseconds.
// The value for this number is from its argument milliseconds.
func Sleep(milliseconds int) {
	time.Sleep(time.Duration(milliseconds) * time.Millisecond)
}
