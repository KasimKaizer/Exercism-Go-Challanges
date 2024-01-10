// wordcount contains tools to count frequency of words.
package wordcount

import (
	"strings"
)

// Frequency is custom type which stores words as string and their frequency as int.
type Frequency map[string]int

// parse takes a rune and returns true or false based on if its a illegal character.
func parse(c rune) bool {
	return strings.ContainsAny(string(c), "\":!&@$%^&,\n\t. ")
}

// WordCount takes a phrase as string and returns a map which contains words from the given phrase
// and there frequency.
func WordCount(phrase string) Frequency {

	lowPhrase := strings.ToLower(phrase)

	// we parse our phrase here while also splitting it at the point when illegal character is
	// encountered
	wordSlice := strings.FieldsFunc(lowPhrase, parse)

	wordMap := make(Frequency, len(wordSlice))

	for _, word := range wordSlice {

		// remove leading and trailing ' from our words, we can't do this in prase func because it
		// would change words like don't to dont.
		word = strings.TrimLeft(word, "'")
		word = strings.TrimRight(word, "'")

		// Skip any word which is just an empty string.
		if word == "" {
			continue
		}

		wordMap[word]++
	}

	return wordMap
}
