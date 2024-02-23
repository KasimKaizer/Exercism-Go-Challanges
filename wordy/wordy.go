// Package wordy contains solution for the Wordy exercise on Exercism.
package wordy

import (
	"strconv"
	"strings"
)

// r contains phrases we want to remove from our passed question.
var r = strings.NewReplacer("What is", "", "by", "", "?", "")

// Answer takes a mathematical question and returns its answer.
// it returns false for invalid questions.
func Answer(question string) (int, bool) {

	// remove phrases we don't need from the question and split by space.
	sanitized := strings.Fields(r.Replace(question))
	sanLen := len(sanitized)
	if sanLen%2 == 0 { // if there are even number of items then the question is invalid.
		return 0, false
	}
	// first item in the slice will always be a number for a valid question.
	cal, err := strconv.Atoi(sanitized[0])
	if err != nil {
		return 0, false
	}

	for i := 1; i < sanLen-1; i = i + 2 { // only loop through operators.
		// operators must be followed by numbers, if not then its a invalid question.
		val, err := strconv.Atoi(sanitized[i+1])
		if err != nil {
			return 0, false
		}
		switch sanitized[i] {
		case "plus":
			cal += val
		case "minus":
			cal -= val
		case "multiplied":
			cal *= val
		case "divided":
			cal /= val
		default: // if its not a supported operator then we have invalid question.
			return 0, false
		}
	}
	return cal, true

}
