package cryptosquare

import (
	"math"
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`[^a-z0-9]+`)

// TODO: use strings.lower to convert the whole string into lowercase.
// TODO: use func trim to remove everything thats not a unicode letter or character.
// TODO: find the perfect length to split the string from, ideally it should be as close
// some multiple, we are looking for a square shape, this would require some thinking.
// TODO: square formula is s*s = a, here the len(input) is the area, so to get the
// to get the rows and colums, we would just underroot(len(input)). we would have to convert
// it back to int value.
// pad the string with spaces to make sure the string fits the multiple completely
// split the string by that multiple and turn it into slice of string, then we can range over
// on of the slices, adn write the letter from all slices in that slice at that idx in a new
// string builder, also add a empty space after the multiple.
func Encode(pt string) string {

	if pt == "" {
		return ""
	}

	normalized := re.ReplaceAllString(strings.ToLower(pt), "")
	normLen := len(normalized) // rough area of the rect / square

	sqrLen, sqBth := sideCal(normLen) // get the square sides.

	for i := 0; i < ((sqrLen * sqBth) - normLen); i++ {
		normalized += " "
	}

	sqrMatrix := matrix(normalized, sqrLen)
	var output strings.Builder

	for col := range sqrMatrix[0] {
		for row := range sqrMatrix {
			output.WriteByte(sqrMatrix[row][col])
		}
	}
	return addSpace(output.String(), sqBth)
}

func sideCal(x int) (int, int) {
	cal := math.Sqrt(float64(x))
	len, bth := int(cal), int(cal)

	switch i := cal - float64(len); {
	case i >= 0.5:
		len++
		bth++
	case i < 0.5 && i > 0:
		len++
	}
	return len, bth
}

func matrix(str string, idx int) [][]byte {
	output := make([][]byte, 0)
	for i := 0; i < len(str); i = i + idx {
		output = append(output, []byte(str[i:i+idx]))
	}
	return output
}

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
