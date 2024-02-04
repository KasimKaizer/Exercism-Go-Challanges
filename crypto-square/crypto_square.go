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
	// normalized the input string by removing all non alphanumeric characters and
	// turning the string into lowercase.
	normalized := re.ReplaceAllString(strings.ToLower(pt), "")
	if normalized == "" {
		return "" // return if normalized string is empty.
	}
	normLen := len(normalized)

	// get the length and breadth of our sqr/rectangle. here sqrLen = rows and
	// sqrBth = columns.
	sqrLen, sqrBth := sideCal(normLen)

	// add whiteSpaces to our string till the length of it is equal to the area of the
	// sqr/rectangle we got from sideCal.
	for i := 0; i < ((sqrLen * sqrBth) - normLen); i++ {
		normalized += " "
	}
	// turn our normalized string into a matrix, basically splitting it by number of rows we want.
	sqrMatrix := matrix(normalized, sqrLen)
	var output strings.Builder

	// Example of matrix iteration:
	// our normalized string: abcdefghij
	// our matrix:
	//   0 1 2 <- columns(c)
	// 0 a b c
	// 1 e f g
	// 2 h i j
	// ^
	// |
	// rows(r)
	// we iterate though like this (where colum is c and row is r):
	// c:0 r:0, then c:0 r:1, then c:0 r:2, c:1 r:0, c:1 r:1, c:1 r:2, c:2 r:1... etc
	// we achieve this by iterating through each colum and in each iteration thet we
	// iterating through each row.
	// so we get string: aehbficgj
	for col := range sqrMatrix[0] { // iterate through each colum of our matrix
		for row := range sqrMatrix { // iterate through each row of our matrix
			// write the char at that particular idx of the matrix to our output
			output.WriteByte(sqrMatrix[row][col])
		}
	}
	return addSpace(output.String(), sqrBth)
}

// sideCal takes an int and returns the sides of the closed rectangle/square to it.
func sideCal(x int) (int, int) {
	// we get the side of our theoretical square.
	cal := math.Sqrt(float64(x))
	// we assume our input x was perfect square and assign initial values cal
	// which was a float value as int. (i.e we drop the decimal point.)
	// as length and breadth.
	len, bth := int(cal), int(cal)

	switch i := cal - float64(len); { // we check out assumption.
	// if our assumption was wrong:
	// if the the difference between int val and float val
	// is more then or equal to 0.5 then increment both len and bth by 1
	case i >= 0.5:
		len++
		bth++
	// if difference was less then 0.5 then just increment len by one.
	case i < 0.5 && i > 0:
		len++
	}
	return len, bth
}

// matrix takes string and an interval and returns a matrix of bytes representing
// the string split at every occurrence of that interval and then converted into bytes.
func matrix(str string, idx int) [][]byte {
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
