// Package meetup contains solution for Meetup exercise on Exercism.
package meetup

import (
	"time"
)

// WeekSchedule defines our descriptors type.
type WeekSchedule int

// assign constants for our descriptors with there respective value.
// the value here indicates the starting of that specific week in the month.
// first is assigned 1 as first week of the month starts at 1st of the month.
// second is assigned 8 as second week of the month starts at 8th of the month.
// and so on..
const (
	First  WeekSchedule = 1
	Second WeekSchedule = 8
	Teenth WeekSchedule = 13
	Third  WeekSchedule = 15
	Fourth WeekSchedule = 22
	Last   WeekSchedule = -1
)

// Day takes our WeekSchedule with weekday, month and year and returns the date of it.
func Day(w WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	start := time.Date(year, month, int(w), 0, 0, 0, 0, time.UTC) // get the start of the week.
	if w < 0 {
		// if w is negative then get the last week of the month.
		start = time.Date(year, month+1, int(w), 0, 0, 0, 0, time.UTC).AddDate(0, 0, -5)
	}
	// iterate from start of the week till the end of the week
	for d := start; d.Before(start.AddDate(0, 0, 8)); d = d.AddDate(0, 0, 1) {
		if d.Weekday() == wDay {
			return d.Day() // return meetup date as soon as we find the weekday which matches.
		}
	}
	return 0 // return zero if there was no match, this indicates an error.
}
