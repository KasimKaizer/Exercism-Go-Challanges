// Package house contains solution for House exercise on Exercism.
package house

import (
	"fmt"
	"strings"
)

const (
	firstVerse = 1
	lastVerse  = 12
)

// TODO: maybe group the nouns and there adjectives together.

// nouns in the song.
var noun = []string{
	"house that Jack built",
	"malt",
	"rat",
	"cat",
	"dog",
	"cow with the crumpled horn",
	"maiden all forlorn",
	"man all tattered and torn",
	"priest all shaven and shorn",
	"rooster that crowed in the morn",
	"farmer sowing his corn",
	"horse and the hound and the horn",
}

// adjectives to be followed by the nouns in the song.
var adjective = []string{
	"lay in",
	"ate",
	"killed",
	"worried",
	"tossed",
	"milked",
	"kissed",
	"married",
	"woke",
	"kept",
	"belonged to",
}

// Verse takes a number and returns the verse of 'This is the House that Jack Built' corresponding
// to that number.
func Verse(v int) string {
	if v > lastVerse || v < firstVerse {
		return ""
	}

	if v == firstVerse {
		return "This is the house that Jack built."
	}
	var output strings.Builder
	output.WriteString(fmt.Sprintf("This is the %s\n", noun[v-1]))

	for i := v - 2; i >= 0; i-- {
		output.WriteString(fmt.Sprintf(
			"that %s the %s",
			adjective[i],
			noun[i],
		))
		if i == 0 {
			output.WriteByte('.')
			continue
		}
		output.WriteByte('\n')
	}
	return output.String()
}

// Song returns 'This is the House that Jack Built' rhyme.
func Song() string {
	var output strings.Builder
	for i := firstVerse; i <= lastVerse; i++ {
		output.WriteString(Verse(i))
		if i < lastVerse {
			output.WriteString("\n\n")
		}
	}
	return output.String()
}
