// Package piglatin contains solution for the Pig Latin exercise on Exercism
package piglatin

import "strings"

// results from three consecutive benchmarks:
// BenchmarkSentence-4       375846              3242 ns/op             608 B/op         47 allocs/op
// BenchmarkSentence-4       266100              3852 ns/op             608 B/op         47 allocs/op
// BenchmarkSentence-4       311935              3546 ns/op             608 B/op         47 allocs/op

// Sentence translates an english sentence to pig latin.
func Sentence(sentence string) string {
	wordArray := strings.Fields(sentence) // split the sentence at every occurrence of whitespace.
	for idx, word := range wordArray {
		// translate the word into pig latin, then replace the original word with translated word.
		wordArray[idx] = wordToPigLatin(word)
	}
	// join all words in the array with a whitespace and return it as string.
	return strings.Join(wordArray, " ")
}

// wordToPigLatin translates an english to pig latin.
func wordToPigLatin(word string) string {
	wrdSlice, iter := []byte(word), 0 // convert the word into a slice of bytes.
loop:
	for !isVowel(wrdSlice[0]) { // loop for as long as the first word in slice is not a vowel
		switch string(wrdSlice[:2]) {
		case "xr", "yt": // check if first two characters make a vowel sound, if so break loop.
			break loop
		case "qu":
			// move 'qu' to the end of slice.
			wrdSlice = append(wrdSlice[2:], wrdSlice[:2]...)
			// this is potential not needed, we can break the loop here as 'qu' is always
			// followed by a  vowel, but I have kept this here as I couldn't find any official
			// confirmation for my theory.
			iter += 2
			continue
		}

		// handle y after a consonant cluster, i.e. more then one iteration of non vowel
		if iter > 1 && wrdSlice[0] == 'y' {
			break
		}
		// move the consonant letter to the end of the slice.
		wrdSlice = append(wrdSlice[1:], wrdSlice[0])
		iter++
	}

	wrdSlice = append(wrdSlice, []byte{'a', 'y'}...) // add 'ay' to the end of the byte slice.
	return string(wrdSlice)                          // convert the byte slice to string and return.
}

// isVowel takes a byte and returns true or false based on if the provided byte is a vowel.
func isVowel(char byte) bool {
	switch char {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	}
	return false
}
