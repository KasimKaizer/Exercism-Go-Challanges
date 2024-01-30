// Package wordy contains solution for the Wordy exercise on Exercism.
package wordy

import (
	"regexp"
	"strconv"
	"strings"
)

// re contain regex to match our requirements for a question, its evaluated during compile time.
var re = regexp.MustCompile(`^(What is) (-?\d+)( (plus|minus|multiplied by|divided by){1} (-?\d+))*( )?(\?$)`)

// r contains phrases we want to remove from our passed question.
var r = strings.NewReplacer("What is ", "", " by", "", "?", "")

// Answer takes a mathematical question and returns its answer.
// it returns false for invalid questions.
func Answer(question string) (int, bool) {
	if !re.MatchString(question) { // use regex to make sure the question fits our rules.
		return 0, false
	}
	// remove phrases we don't need from the question, this makes it easy to parse.
	// then split it by space, so we are left with just numbers and operations.
	sanitized := strings.Split(r.Replace(question), " ")
	// first item in the slice will always be a number for a valid question.
	cal, err := strconv.Atoi(sanitized[0])
	if err != nil {
		return 0, false
	}
	for i := 1; i < len(sanitized)-1; i++ { // iterate from second item to second last item.
		// see if the item after the current item is a number.
		val, err := strconv.Atoi(sanitized[i+1])
		if err != nil {
			continue // if not then move on to next iteration.
		}
		// if the item after current item is a number, then the current item has to be operation.
		// switch based on which operation it is and update the cal variable by performing this
		// operation on cal variable and the item after the current item in this loop.
		switch sanitized[i] {
		case "plus":
			cal += val
		case "minus":
			cal -= val
		case "multiplied":
			cal *= val
		case "divided":
			cal /= val
		}
	}
	return cal, true

}
