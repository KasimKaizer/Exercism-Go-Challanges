// Package diamond contains solution for Diamond exercise on Exercism
package diamond

import (
	"errors"
	"strings"
)

// error for invalid input
var errInvInput = errors.New("input has to be an uppercase letter")

// Gen takes a capital letter and returns a diamond shape corelating to it.
func Gen(char byte) (string, error) {
	if char < 'A' || char > 'Z' {
		return "", errInvInput // diamond can only be constructed for capital letters.
	}

	length := (int(char-'A') * 2) // get the total number of rows in the diamond.

	output := make([]string, length+1)
	// we only need to iterate through the half of the total length as the diamond is symmetrical
	// from top and bottom.
	for i := 0; i <= length/2; i++ {
		// as number of rows is equal to the width of each row in the diamond, pass the width and
		// current row number to genLine function.
		line := genLine(length, i)

		// assign the line to its correlating rows.
		output[i], output[length-i] = line, line
	}
	// return a string which is just the array joined by a newline.
	return strings.Join(output, "\n"), nil
}

// genLine takes the width of the line and the position of a capital letters in the alphabet
// and returns a row of the diamond representing that letter.
func genLine(length, charPos int) string {
	output, toIter := make([]byte, length+1), length/2

	// we only need to iterate half of the row's length as two half's of row are symmetrical.
	for i := 0; i <= toIter; i++ {
		// check if the current position in the line is the position for the letter.
		if charPos == (toIter - i) {
			letter := byte('A' + charPos) // get the letter correlating to the position
			output[i], output[length-i] = letter, letter
			continue
		}
		output[i], output[length-i] = ' ', ' '
	}
	return string(output)
}
