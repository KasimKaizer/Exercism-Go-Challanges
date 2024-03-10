// Package matrix contains solution for Saddle Points exercise on Exercism.
package matrix

import (
	"math"
	"strconv"
	"strings"
)

// Matrix represents a matrix of int.
type Matrix struct {
	values [][]int
}

// Pair represents the position on a matrix.
type Pair struct {
	row int
	col int
}

// New creates a new matrix from the given string.
func New(s string) (*Matrix, error) {
	if len(s) < 1 {
		return new(Matrix), nil
	}
	rows := strings.Split(s, "\n")
	output := make([][]int, len(rows))

	for i, row := range rows {
		splitRow := strings.Fields(row)
		output[i] = make([]int, len(splitRow))
		for j, numChar := range splitRow {
			num, err := strconv.Atoi(numChar)
			if err != nil {
				return nil, err
			}
			output[i][j] = num
		}
	}
	return &Matrix{values: output}, nil
}

// Saddle returns the position of trees that are ideal for a tree house location.
func (m *Matrix) Saddle() []Pair {
	output := make([]Pair, 0)
	if len(m.values) < 1 {
		return output
	}
	cache := make([]int, len(m.values[0])) // save the minimum values of columns in a cache.
	for row := range m.values {
		for _, colNum := range m.rowMaxPos(row) {
			if cache[colNum] == 0 {
				cache[colNum] = m.colMin(colNum)
			}
			if m.values[row][colNum] != cache[colNum] {
				continue
			}
			output = append(output, Pair{row + 1, colNum + 1})
		}
	}
	return output
}

// rowMaxPos method takes a number number and returns the positions of the largest
// numbers in that row.
func (m *Matrix) rowMaxPos(row int) []int {
	max := 0
	output := make([]int, 0)
	for idx, num := range m.values[row] {
		if num < max {
			continue
		}
		if num == max {
			output = append(output, idx)
			continue
		}
		max = num
		output = []int{idx}
	}
	return output
}

// colMin takes a column number and returns the smaller number in that column.
func (m *Matrix) colMin(col int) int {
	min := math.MaxInt64
	for idx := range m.values {
		if m.values[idx][col] > min {
			continue
		}
		min = m.values[idx][col]
	}
	return min
}
