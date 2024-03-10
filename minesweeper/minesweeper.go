// Package minesweeper contains solution for Minesweeper exercise on exercism.
package minesweeper

import "strings"

// directions to all of all sides in 2d matrix
type direction struct {
	row, col int
}

// directions represents all the directions that can be traversed from a point in a matrix.
var directions = [...]direction{
	{-1, -1}, // top left
	{-1, 0},  // top
	{-1, 1},  // top right
	{0, 1},   // right
	{1, 1},   // bottom right
	{1, 0},   // bottom
	{1, -1},  // bottom left
	{0, -1},  // left
}

// Annotate adds the mine counts to a completed Minesweeper board. it takes the board in form of a
// slice of string, where each string represents a row on the board.
func Annotate(board []string) []string {
	rows := len(board)
	if rows < 1 {
		return board // board can't be empty
	}
	col := len(board[0])
	if col < 1 {
		return board // a board needs to have columns
	}
	output := make([]string, rows)
	for i := 0; i < rows; i++ {
		var curRow strings.Builder
		curRow.Grow(col)
		for j := 0; j < col; j++ {
			if board[i][j] == '*' {
				curRow.WriteByte('*')
				continue
			}
			count := search(board, i, j)
			if count == 0 {
				curRow.WriteByte(' ')
				continue
			}
			curRow.WriteByte('0' + byte(count))
		}
		output[i] = curRow.String()
	}
	return output
}

// search takes the board and coordinates of the current position. it then
// searches in all possible directions from that position for a mine and returns the count.
func search(board []string, rowPos, colPos int) int {
	lastRowPos, LastColPos := len(board)-1, len(board[0])-1
	var count int
	for _, dir := range directions {
		nRowPos := rowPos + dir.row
		nColPos := colPos + dir.col
		if nRowPos > lastRowPos || nRowPos < 0 {
			continue
		}
		if nColPos > LastColPos || nColPos < 0 {
			continue
		}
		if board[nRowPos][nColPos] == '*' {
			count++
		}
	}
	return count
}
