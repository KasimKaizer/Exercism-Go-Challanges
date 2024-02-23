// Package matrix contains solution for the Matrix Exercise on Exercism.
package matrix

import (
	"errors"
	"strconv"
	"strings"
)

var (
	errInvInput = errors.New("invalid / badly formatted input") // error for invalid input
)

// Matrix type is an array of arrays of int.
type Matrix [][]int

// New takes a string and converts it into a matrix of integers.
func New(s string) (Matrix, error) {
	if s == "" {
		return nil, errInvInput // return an error if input is empty
	}
	// split the string at new line, so we get each string representing each row of a matrix
	splitS, prevLen := strings.Split(s, "\n"), 0
	output := make(Matrix, len(splitS))

	for idx, row := range splitS {
		// split the row by whitespace, so now we have the array for the current row, but
		// its values are in form of string.
		curRow := strings.Split(strings.TrimSpace(row), " ")
		curLen := len(curRow)
		if idx > 0 && curLen != prevLen {
			// check if the number of items in current row are equal to number of items in the
			// previous row, if not then the input string was invalid.
			return nil, errInvInput
		}
		prevLen = curLen
		for _, numStr := range curRow {
			// go through each item in the current row, convert it into an integer.
			num, err := strconv.Atoi(numStr)
			if err != nil {
				// if its not able to be converted into an integer then input is invalid.
				return nil, err
			}
			// append the integer to the output matrix where row number is current idx.
			output[idx] = append(output[idx], num)
		}
	}
	return output, nil
}

// Cols method returns a transpose copy of current matrix.
func (m Matrix) Cols() [][]int {
	output := make([][]int, len(m[0])) // make a new matrix.
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			// append the items in the rows of the current matrix to the colum
			// of the output matrix.
			output[j] = append(output[j], m[i][j])

		}
	}
	return output
}

// Rows method returns a copy of current matrix in form of array of arrays of int.
func (m Matrix) Rows() [][]int {
	output := make([][]int, len(m)) // make a new matrix.
	rowLen := len(m[0])
	for i := range m {
		// for each row in current matrix, make a row of equal size in output matrix.
		output[i] = make([]int, rowLen)
		copy(output[i], m[i]) // copy all items from the current matrix to output matrix.
	}
	return output
}

// Set method takes a value, row number, and colum number and replaces the value at position
// mapping to the row and colum number with the input value.
func (m Matrix) Set(row, col, val int) bool {
	if row > len(m)-1 || row < 0 || col > len(m[0])-1 || col < 0 {
		return false // if row/ col number are out of bound then return false.
	}

	m[row][col] = val
	return true
}
