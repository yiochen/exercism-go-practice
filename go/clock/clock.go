//Package clock Provides functionality of a clock
package clock

import "fmt"

const testVersion = 4

// Clock the Clock type
type Clock struct {
	hour   int
	minute int
}

// New Creates and returns a new Clock.
func New(hour, minute int) Clock {
	stepIn := 0

	if minute < 0 {
		borrowed := -minute/60 + 1
		minute = minute + borrowed*60
		stepIn -= borrowed
	}

	stepIn += minute / 60
	minute = minute % 60

	hour = hour + stepIn

	if hour < 0 {
		hour = (-hour/24+1)*24 + hour
	}

	hour = hour % 24
	return Clock{hour, minute}
}

// String return a string representation
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add return a new Clock with added minutes
func (c Clock) Add(minutes int) Clock {
	return New(c.hour, c.minute+minutes)
}
