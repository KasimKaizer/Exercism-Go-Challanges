// Package railfence contains various tools to encode/decode text in rail fence cipher
package railfence

import (
	"strings"
)

// grid represents a grid used to encode/ decode text.
type grid [][]rune

// railfence creates a new gird of the provided size.
func newGrid(rows, cols int) grid {
	output := make(grid, rows)
	for idx := range output {
		output[idx] = make([]rune, cols)
	}
	return output
}

// Encode encodes a message using rail fence cipher and returns the encoded message.
func Encode(message string, rails int) string {
	msgSlice := []rune(message)

	grid := newGrid(rails, len(msgSlice))

	grid.populate(msgSlice, len(msgSlice))

	var output strings.Builder
	for _, row := range grid {
		for _, char := range row {
			if char == '\x00' {
				continue
			}
			output.WriteRune(char)
		}
	}
	return output.String()
}

// Decode decodes a message using rail fence cipher and returns the result.
func Decode(message string, rails int) string {
	msgSlice := []rune(message)
	grid := newGrid(rails, len(msgSlice))
	grid.populate([]rune{'?'}, len(msgSlice))
	idx := 0
	for i := 0; i < len(grid) && idx < len(msgSlice); i++ {
		for j := 0; j < len(grid[i]) && idx < len(msgSlice); j++ {
			if grid[i][j] != '?' {
				continue
			}
			grid[i][j] = msgSlice[idx]
			idx++
		}
	}
	var output strings.Builder
	i := 0
	for i < len(msgSlice) {
		for j := 0; j < len(grid) && i < len(msgSlice); j++ {
			output.WriteRune(grid[j][i])
			i++
		}
		for k := len(grid) - 2; k > 0 && i < len(msgSlice); k-- {
			output.WriteRune(grid[k][i])
			i++
		}
	}
	return output.String()
}

// populate method populates the contents of the message into the grid.
func (g grid) populate(message []rune, msgLen int) {
	i := 0
	for i < msgLen {
		for j := 0; j < len(g) && i < msgLen; j++ {
			g[j][i] = message[i%len(message)]
			i++
		}
		for k := len(g) - 2; k > 0 && i < msgLen; k-- {
			g[k][i] = message[i%len(message)]
			i++
		}
	}
}
