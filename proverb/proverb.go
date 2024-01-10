// proverb package contains fun tools to make a sentences.
package proverb

import "fmt"

// Proverb takes a slice of words and returns a proverb made from them.
func Proverb(rhyme []string) []string {

	var output []string
	ln := len(rhyme)

	for i, j := 0, 1; i < ln; i, j = i+1, j+1 {

		// Condition to check if we are on the last word from the
		// slice.
		if j == ln {
			return append(output, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
		}

		output = append(output, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[j]))
	}
	return output
}
