// Package contains tools to work with time.
package clock

import (
	"fmt"
)

const (
	minsInDay  = 1440
	minsInHour = 60
)

// limit takes the un-sanitized value of minutes and limits them minutes in a day.
func limit(m int) Clock {

	m %= minsInDay

	if m < 0 {
		m = minsInDay + m
	}
	return Clock{minutes: m}
}

// Clock struct contains (minutes) integers which represent total minutes of a day.
type Clock struct {
	minutes int
}

// New takes the desired hours and minutes and returns a clock.
func New(h, m int) Clock {
	return limit((h * minsInHour) + m)
}

// Add takes the number of minutes to add to clock and returns the updated clock.
func (c Clock) Add(m int) Clock {
	return limit(c.minutes + m)
}

// Subtract takes the number of minutes to subtract from the clock and returns the updated clock.
func (c Clock) Subtract(m int) Clock {
	return limit(c.minutes - m)
}

// String returns the clock with proper format as a string.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minutes/minsInHour, c.minutes%minsInHour)
}
