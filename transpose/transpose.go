// Package transpose contains solution for Transpose exercise on Exercism.
package transpose

import (
	"strings"
)

// Transpose return the transpose of the given matrix.
func Transpose(input []string) []string {
	max := length(input)

	// we use a buffer here instead of a string as we don't want to constantly concatenate a string.
	buffer := make([]strings.Builder, max)

	for size, row := range input {
		for col, char := range row {
			for i := buffer[col].Len(); i < size; i++ {
				buffer[col].WriteByte(' ')
			}
			buffer[col].WriteRune(char)
		}
	}

	output := make([]string, max)

	for i := 0; i < max; i++ {
		output[i] = buffer[i].String()
	}

	return output
}

// length finds the longest string in a the given slice and returns its length.
func length(input []string) int {
	max := 0
	for _, str := range input {
		strLen := len(str)
		if strLen > max {
			max = strLen
		}
	}
	return max
}

// old solution, This solution is 30% faster but it doesn't account for UTF-8 characters.
// func transpose(input []string, max, min int) []string {
// 	output := make([]string, max)
// 	colMax := len(input)
// 	for i := 0; i < max; i++ {
// 		var curRow strings.Builder
// 		replace := false
// 		for j := 0; j < colMax; j++ {
// 			if len(input[j]) <= i {
// 				curRow.WriteByte(' ')
// 				continue
// 			}
// 			if input[j][i] == ' ' {
// 				curRow.WriteByte('\x00')
// 				replace = true
// 				continue
// 			}
// 			curRow.WriteByte(input[j][i])
// 		}
// 		row := curRow.String()
// 		if i >= min {
// 			row = strings.TrimRight(row, " ")
// 		}
// 		if replace {
// 			row = strings.ReplaceAll(row, "\x00", " ")
// 		}
// 		output[i] = row
// 	}
// 	return output
// }
