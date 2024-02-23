// Package encode contains solution for Run-Length Encoding exercise on Exercism.
package encode

import (
	"log"
	"strconv"
	"strings"
)

// RunLengthEncode takes a string and encodes it according to Run-Length Encoding.
func RunLengthEncode(input string) string {
	inpByte := []byte(input)    // convert input into a slice of bytes.
	lastPos := len(inpByte) - 1 // get the position of last char in the byte slice.
	pos := 0                    // start position

	var output strings.Builder
	for pos <= lastPos {
		// if the current char is not equal to next char or we are at the at last position,
		// then add that character to output slice and continue with the loop.
		if pos == lastPos || inpByte[pos] != inpByte[pos+1] {
			output.WriteByte(inpByte[pos])
			pos++
			continue
		}
		j := pos
		// if the above condition doesn't trigger then we know we have
		// repeating characters, check for how long this character repeats.
		for j <= lastPos && inpByte[pos] == inpByte[j] {
			j++
		}
		output.WriteString(strconv.Itoa(j - pos)) // write the count to the output.
		// add the (count - 1) to the pos and continue with the loop
		pos = j - 1
	}

	return output.String()
}

// RunLengthDecode takes a string encoded using Run-Length Encoding and decodes it.
func RunLengthDecode(input string) string {
	inpByte := []byte(input)
	lastPos := len(inpByte) - 1
	pos := 0

	var output strings.Builder
	for pos <= lastPos {
		// if the current char is not a number then add it to the output and continue with the
		// loop.
		if inpByte[pos] < '0' || inpByte[pos] > '9' {
			output.WriteByte(inpByte[pos])
			pos++
			continue
		}
		// if the current char is a number then get the whole number.
		var numStr strings.Builder
		for inpByte[pos] >= '0' && inpByte[pos] <= '9' {
			numStr.WriteByte(inpByte[pos])
			pos++ // this loop will continue till pos is at a non numeric character.
		}
		num, err := strconv.Atoi(numStr.String())
		if err != nil {
			log.Fatal(err)
		}
		// write the char at current pos to output for 'num' amount of times.
		for i := 0; i < num; i++ {
			output.WriteByte(inpByte[pos])
		}
		pos++ // continue the loop.
	}
	return output.String()
}
