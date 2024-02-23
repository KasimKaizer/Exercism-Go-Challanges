// Package foodchain contains solution for Food Chain exercise on Exercism.
package foodchain

import (
	"fmt"
	"strings"
)

const (
	lastVerse   = 8
	secondVerse = 2
	firstVerse  = 1
)

// animal defines a animal's name and their complementary lines.
type animal struct {
	name       string
	complement string
}

// list of all animals and there complementary line in the 'I Know an Old Lady Who Swallowed a Fly' song.
var animalList = []animal{
	{"fly", "\nI don't know why she swallowed the fly. Perhaps she'll die."},
	{"spider", " that wriggled and jiggled and tickled inside her."},
	{"bird", "\nHow absurd to swallow a bird!"},
	{"cat", "\nImagine that, to swallow a cat!"},
	{"dog", "\nWhat a hog, to swallow a dog!"},
	{"goat", "\nJust opened her throat and swallowed a goat!"},
	{"cow", "\nI don't know how she swallowed a cow!"},
	{"horse", "\nShe's dead, of course!"},
}

// Verse takes a number and returns its corresponding verse.
func Verse(v int) string {
	if v > lastVerse {
		return ""
	}

	var output strings.Builder
	output.WriteString(fmt.Sprintf("I know an old lady who swallowed a %s.", animalList[v-1].name))

	complement := animalList[v-1].complement
	if v == secondVerse {
		complement = fmt.Sprint("\nIt wriggled and jiggled and tickled inside her.")
	}
	output.WriteString(complement)

	if v == lastVerse {
		return output.String()
	}

	for i := v - 1; i > 0; i-- {
		output.WriteString(fmt.Sprintf("\nShe swallowed the %s to catch the %s",
			animalList[i].name,
			animalList[i-1].name,
		))
		if i != secondVerse {
			output.WriteByte('.')
		}
		if i <= secondVerse {
			output.WriteString(animalList[i-1].complement)
		}
	}
	return output.String()
}

// Verses returns the song from verse to corresponding 'start' till verse corresponding to 'end'.
func Verses(start, end int) string {
	if start > end || end > lastVerse {
		return ""
	}
	var output strings.Builder
	for i := start; i <= end; i++ {
		output.WriteString(Verse(i))
		output.WriteString("\n\n")
	}
	return strings.TrimSuffix(output.String(), "\n\n")
}

// Song returns the 'I Know an Old Lady Who Swallowed a Fly' song.
func Song() string {
	return Verses(firstVerse, lastVerse)
}
