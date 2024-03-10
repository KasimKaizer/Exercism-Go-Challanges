package wordsearch

import "errors"

// directions contains various directions in 2d matrix.
var directions = [][]int{
	{-1, 0},  // top
	{1, 0},   // bottom
	{0, 1},   // right
	{0, -1},  // left
	{-1, 1},  // top right
	{1, 1},   // bottom right
	{1, -1},  // bottom left
	{-1, -1}, // top left
}

var (
	errNotFound = errors.New("word not present in puzzle")
	zeroArray   = [2]int{}
	zeroMatrix  = [2][2]int{}
)

// Solve takes a word puzzle and a list of words to find in them, and returns each word's position
// in the word puzzle.
func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	wordMap := make(map[string][2][2]int)
	for _, word := range words {
		pos, ok := solveWord(word, puzzle)
		if !ok {
			return nil, errNotFound
		}
		wordMap[word] = pos
	}
	return wordMap, nil
}

// solveWord takes a word and a word puzzle, it returns the position of the word in the
// word puzzle.
func solveWord(word string, puzzle []string) ([2][2]int, bool) {
	rowLen, colLen := len(puzzle), len(puzzle[0])
	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			end, ok := find(word, puzzle, i, j)
			if !ok {
				continue
			}
			return [2][2]int{{j, i}, end}, true
		}
	}
	return zeroMatrix, false
}

// find takes a word puzzle, a word, and a position in word puzzle as input. it searches
// for matches for that word in all directions from that position. if a match is found
// then it returns the last position of the word in word puzzle.
func find(word string, puzzle []string, row, col int) ([2]int, bool) {
	if word[0] != puzzle[row][col] {
		return zeroArray, false
	}
	rowLen, colLen, wordLen := len(puzzle), len(puzzle[row]), len(word)-1

	// range though all possible directions we can can proceed
	for _, dir := range directions {
		rowDir, colDir := dir[0], dir[1]
		nRow, nCol := row+rowDir, col+colDir // get the new row and col position

		for i := 1; i <= wordLen; i++ {
			if nRow >= rowLen || nRow < 0 {
				break
			}
			if nCol >= colLen || nCol < 0 {
				break
			}
			if word[i] != puzzle[nRow][nCol] {
				break
			}
			if i == wordLen {
				return [2]int{nCol, nRow}, true
			}
			nRow += rowDir
			nCol += colDir
		}

	}

	return zeroArray, false
}
