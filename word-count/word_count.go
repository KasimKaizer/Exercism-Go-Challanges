// wordcount contains tools to count frequency of words.
package wordcount

import (
	"regexp"
	"strings"
)

// Frequency is custom type which stores words as string and their frequency as int.
type Frequency map[string]int

// regex to find words inside the phrase.
var re = regexp.MustCompile(`(\w+('\b)?\w*)`)

// WordCount takes a phrase as string and returns a map which contains words from the given phrase
// and there frequency.
func WordCount(phrase string) Frequency {

	lowPhrase := strings.ToLower(phrase)

	wordMap := make(Frequency)

	// here we iterate through all the matches (which are words) in the phrase.
	for _, word := range re.FindAllString(lowPhrase, -1) {

		wordMap[word]++
	}

	return wordMap
}
