// Package matrix contains tools to implement and work with matrix data structure
package matrix

import (
	"errors"
	"strconv"
	"strings"
)

var (
	errInvInput = errors.New("invalid or badly formatted input")
)

// Matrix type is an array of arrays of int.
type Matrix [][]int

// New takes a string and converts it into a matrix of integers.
func New(s string) (Matrix, error) {
	if s == "" {
		return nil, errInvInput
	}
	splitS := strings.Split(s, "\n")
	prevLen := 0
	output := make(Matrix, len(splitS))
	for idx, row := range splitS {
		curRow := strings.Split(strings.TrimSpace(row), " ")
		curLen := len(curRow)
		if idx > 0 && curLen != prevLen {
			// check if the number of items in current row are equal to number of items in the
			// previous row, if not then the input string was invalid.
			return nil, errInvInput
		}
		prevLen = curLen
		for _, numStr := range curRow {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return nil, err
			}
			output[idx] = append(output[idx], num)
		}
	}
	return output, nil
}

// Cols method returns a transpose copy of current matrix.
func (m Matrix) Cols() [][]int {
	output := make([][]int, len(m[0]))
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			output[j] = append(output[j], m[i][j])

		}
	}
	return output
}

// Rows method returns a copy of current matrix in form of array of arrays of int.
func (m Matrix) Rows() [][]int {
	output := make([][]int, len(m))
	rowLen := len(m[0])
	for i := range m {
		output[i] = make([]int, rowLen)
		copy(output[i], m[i])
	}
	return output
}

// Set method takes a value, row number, and colum number and replaces the value at position
// mapping to the row and colum number with the input value.
func (m Matrix) Set(row, col, val int) bool {
	if row > len(m)-1 || row < 0 || col > len(m[0])-1 || col < 0 {
		return false // if row or col number are out of bound then return false.
	}
	m[row][col] = val
	return true
}
