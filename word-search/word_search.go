package wordsearch

import "errors"

var directions = [][]int{
	// {row, column}
	{-1, 0},  // north
	{1, 0},   // south
	{0, 1},   // east
	{0, -1},  // west
	{-1, 1},  // north east
	{1, 1},   // south east
	{1, -1},  // south west
	{-1, -1}, // north west
}

var (
	errNotFound = errors.New("word not present in puzzle")
	defArray    = [2]int{}
	defMatrix   = [2][2]int{}
)

func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	wordMap := make(map[string][2][2]int)
	for _, word := range words {
		pos, ok := wordSolve(word, puzzle)
		if !ok {
			return nil, errNotFound
		}
		wordMap[word] = pos
	}
	return wordMap, nil
}

// function to search a single word
func wordSolve(word string, puzzle []string) ([2][2]int, bool) {
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
	return defMatrix, false
}

func find(word string, puzzle []string, row, col int) ([2]int, bool) {
	if word[0] != puzzle[row][col] {
		return defArray, false
	}
	// we  found a possible starting point
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

	return defArray, false
}
