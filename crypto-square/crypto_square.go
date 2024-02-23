// Package cryptosquare contains the solution for the
package cryptosquare

import (
	"math"
	"regexp"
	"strings"
)

// re matches all non alphanumeric characters.
var re = regexp.MustCompile(`[^a-z0-9]+`)

// Encode takes a string and encodes it to square code.
func Encode(pt string) string {

	var normalized strings.Builder
	normalized.WriteString(re.ReplaceAllString(strings.ToLower(pt), ""))
	if normalized.String() == "" {
		return ""
	}
	normLen := normalized.Len()
	columns, rows := sideCal(normLen)

	// add whiteSpaces to our string till the length of it is equal to the area of the
	// sqr/rectangle we got from sideCal.
	for i := 0; i < ((columns * rows) - normLen); i++ {
		normalized.WriteByte(' ')
	}

	sqrMatrix := splitByIndex(normalized.String(), columns)

	var output strings.Builder
	for col := range sqrMatrix[0] {
		for row := range sqrMatrix {
			output.WriteByte(sqrMatrix[row][col])
		}
	}
	return addSpace(output.String(), rows)
}

// sideCal takes an int and returns the sides of the closed square to it.
func sideCal(x int) (int, int) {
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
func splitByIndex(str string, idx int) [][]byte {
	output := make([][]byte, 0)
	for i := 0; i < len(str); i = i + idx {
		output = append(output, []byte(str[i:i+idx]))
	}
	return output
}

// addSpace takes a string and a interval and returns a string where a whiteSpace is added
// at every occurrence of the provided interval.
func addSpace(str string, idx int) string {
	var output strings.Builder
	pos := idx - 1
	lastPos := len(str) - 1
	for i, char := range []byte(str) {
		output.WriteByte(char)
		if i%idx == pos && i != lastPos {
			output.WriteByte(' ')
		}
	}
	return output.String()
}
