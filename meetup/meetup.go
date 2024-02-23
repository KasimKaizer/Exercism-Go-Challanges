// Package meetup contains solution for Meetup exercise on Exercism.
package meetup

import (
	"time"
)

// WeekSchedule defines our descriptors type.
type WeekSchedule int

// assign constants for our descriptors with there respective value.
const (
	First  WeekSchedule = 1
	Second WeekSchedule = 8
	Teenth WeekSchedule = 13
	Third  WeekSchedule = 15
	Fourth WeekSchedule = 22
	Last   WeekSchedule = -6
)

// Day takes our WeekSchedule with weekday, month and year and returns the date of it.
func Day(w WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	start := time.Date(year, month, int(w), 0, 0, 0, 0, time.UTC)
	if w < 0 {
		start = time.Date(year, month+1, int(w), 0, 0, 0, 0, time.UTC)
	}
	// iterate from start of the week till the end of the week
	for d := start; d.Before(start.AddDate(0, 0, 8)); d = d.AddDate(0, 0, 1) {
		if d.Weekday() == wDay {
			return d.Day()
		}
	}
	return 0
}
