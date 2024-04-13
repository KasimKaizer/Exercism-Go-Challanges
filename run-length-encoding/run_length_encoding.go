// Package encode contains tools to implement run length encoding
package encode

import (
	"log"
	"strconv"
	"strings"
)

// RunLengthEncode takes a string and encodes it according to Run-Length Encoding.
func RunLengthEncode(input string) string {
	lastPos := len(input) - 1
	pos := 0 // start position

	var output strings.Builder
	for pos <= lastPos {
		if pos == lastPos || input[pos] != input[pos+1] {
			output.WriteByte(input[pos])
			pos++
			continue
		}
		j := pos
		// if the above condition doesn't trigger then we know we have
		// repeating characters, check for how long this character repeats.
		for j <= lastPos && input[pos] == input[j] {
			j++
		}
		output.WriteString(strconv.Itoa(j - pos))
		pos = j - 1
	}

	return output.String()
}

// RunLengthDecode takes a string encoded using Run-Length Encoding and decodes it.
func RunLengthDecode(input string) string {
	lastPos := len(input) - 1
	pos := 0

	var output strings.Builder
	for pos <= lastPos {
		if input[pos] < '0' || input[pos] > '9' {
			output.WriteByte(input[pos])
			pos++
			continue
		}
		// if the current char is a number then get the whole number.
		var numStr strings.Builder
		for input[pos] >= '0' && input[pos] <= '9' {
			numStr.WriteByte(input[pos])
			pos++ // this loop will continue till pos is at a non numeric character.
		}
		num, err := strconv.Atoi(numStr.String())
		if err != nil {
			log.Fatal(err)
		}
		// write the char at current pos to output for 'num' amount of times.
		for i := 0; i < num; i++ {
			output.WriteByte(input[pos])
		}
		pos++
	}
	return output.String()
}
