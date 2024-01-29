// Package twelve contains the solution for Twelve Days exercise on Exercism.
package twelve

import (
	"fmt"
	"strings"
)

// totalVerse is the total number of verses in twelve day song.
const totalVerse = 12

// numbers contains all english numbers from first to twelfth.
var numbers = []string{"first", "second", "third", "fourth", "fifth", "sixth",
	"seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}

// presents contains all the presents from twelve day song.
var presents = []string{" twelve Drummers Drumming", " eleven Pipers Piping",
	" ten Lords-a-Leaping", " nine Ladies Dancing", " eight Maids-a-Milking",
	" seven Swans-a-Swimming", " six Geese-a-Laying", " five Gold Rings",
	" four Calling Birds", " three French Hens", " two Turtle Doves",
	" and a Partridge"}

// Verse takes a number and returns the verse of twelve day song that correlates with it.
func Verse(i int) string {
	return fmt.Sprintf("On the %s day of Christmas my true love gave to me:%s in a Pear Tree.", numbers[i-1], giftVerse(i))
}

// Song returns the twelve day song.
func Song() string {
	var output strings.Builder
	for i := 1; i <= totalVerse; i++ { // cycle through all the verses.
		output.WriteString(Verse(i)) // write the current verse.
		output.WriteString("\n")     // add a line break.
	}
	return strings.TrimSuffix(output.String(), "\n") // remove '\n' from the end of the song.
}

// giftVerse constructs the gift part of the verse for the given day.
func giftVerse(x int) string {
	if x == 1 {
		return " a Partridge"
	}
	// join all the gifts from the given point till the beginning with a ','
	return strings.Join(presents[totalVerse-x:], ",")
}
