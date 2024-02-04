// Package bottlesong contains solution for the Bottle Song exercise on Exercism.
package bottlesong

import (
	"fmt"
	"strings"
)

// numbers is a slice of string that contains english numbers.
var numbers = []string{"No", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}

// pluralB takes a number and based on that returns singular or plural of bottle.
func pluralB(i int) string {
	if i == 1 {
		return "bottle"
	}
	return "bottles"
}

// Recite takes the number of starting bottles and how many bottles we need to take down, and based on that create a
// children's poem.
func Recite(startBottles, takeDown int) []string {
	var output []string

	for i := startBottles; i > (startBottles - takeDown); i-- {

		if i != startBottles {
			output = append(output, "") // insert an empty string inside the output array after each verse.
		}

		for j := 0; j < 2; j++ { // as first two lines in a verse are same, we repeat it.
			output = append(output, fmt.Sprintf(
				"%s green %s hanging on the wall,",
				numbers[i],
				pluralB(i), // we check if it should be bottle or bottles.
			))
		}

		output = append(output, "And if one green bottle should accidentally fall,")

		output = append(output, fmt.Sprintf(
			"There'll be %s green %s hanging on the wall.",
			strings.ToLower(numbers[i-1]),
			pluralB(i-1),
		))
	}

	return output
}
