// Package cryptosquare contains the solution for the
package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

// Encode takes a string and encodes it to square code.
func Encode(pt string) string {

	normal := normalize(pt)
	if normal.Len() == 0 {
		return ""
	}
	columns, rows := calRectangleDimensions(normal.Len())

	// add whiteSpaces to our string till the length of it is equal to the area of the
	// sqr/rectangle we got from sideCal.
	normal.WriteString(strings.Repeat(" ", (columns*rows)-normal.Len()))

	sqrMatrix := splitByChunkLength(normal.String(), columns)

	var output strings.Builder
	for col := range sqrMatrix[0] {
		for row := range sqrMatrix {
			output.WriteByte(sqrMatrix[row][col])
		}
		output.WriteByte(' ')
	}
	return strings.TrimSuffix(output.String(), " ")
}

// sideCal takes an int and returns the sides of the closed square to it.
func calRectangleDimensions(x int) (int, int) {
	cal := math.Sqrt(float64(x))
	columns, rows := int(cal), int(cal)

	switch i := cal - float64(columns); {
	case i >= 0.5:
		columns++
		rows++
	case i < 0.5 && i > 0:
		columns++
	}
	return columns, rows
}

// splitByIndex takes a string and an interval and returns a [][]byte representing
// the string split at every occurrence of that interval and then converted into bytes.
func splitByChunkLength(str string, idx int) [][]byte {
	output := make([][]byte, len(str)/idx)
	j := 0
	for i := 0; i < len(str); i = i + idx {
		output[j] = []byte(str[i : i+idx])
		j++
	}
	return output
}

// normalize normalizes the input string by removing any character thats not a alphanumeric.
func normalize(input string) *strings.Builder {
	var output strings.Builder
	for _, char := range input {
		if unicode.IsNumber(char) || unicode.IsLetter(char) {
			if unicode.IsUpper(char) {
				char = unicode.ToLower(char)
			}
			output.WriteRune(char)
		}

	}
	return &output
}
